package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/Lezh1k/NewsService/commons"
	msg "github.com/Lezh1k/NewsService/commons/pb"
	"github.com/Lezh1k/NewsService/storage/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogo/protobuf/proto"
	"github.com/jmoiron/sqlx"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/sirupsen/logrus"
)

func newRegisterSubscribe(natsClient stan.Conn, conn *sqlx.DB) (stan.Subscription, error) {
	return natsClient.Subscribe("register_req", func(m *stan.Msg) {
		nrrq := new(msg.NewsRegisterRequestContainer)
		err := proto.Unmarshal(m.Data, nrrq)
		if err != nil {
			logrus.Errorf("Couldn't unmarshal message. Err : %v", err)
			return
		}
		logrus.Infof("Received a message. Guid : %s", nrrq.Guid)

		nrrs := new(msg.NewsRegisterResponseContainer)
		nrrs.Guid = nrrq.Guid

		t, err := ptypes.Timestamp(nrrq.Item.Date)
		if err != nil {
			logrus.Errorf("Failed to unmarshal nrrq.Item.Date. err : %v", err)
			return
		}

		r, err := conn.Exec("insert into news (registered, header) values (?, ?)", t, nrrq.Item.Header)
		liid := int64(-1)
		if err != nil {
			logrus.Errorf("Failed to exec sql request. Err : %v", err)
		} else {
			liid, err = r.LastInsertId()
		}
		nrrs.Id = fmt.Sprintf("%d", liid)
		nrrsData, err := proto.Marshal(nrrs)
		if err != nil {
			logrus.Errorf("Couldn't marshal response. Err : %v", err)
			return
		}
		natsClient.Publish("register_resp", nrrsData)
	})
}

//NewsDbItem represents db table News.
type NewsDbItem struct {
	ID     int       `db:"id"`
	Date   time.Time `db:"registered"`
	Header string    `db:"header"`
}

func newInfoSubscribe(natsClient stan.Conn, conn *sqlx.DB) (stan.Subscription, error) {
	return natsClient.Subscribe("info_request", func(m *stan.Msg) {
		nirq := new(msg.NewsInfoRequestContainer)
		err := proto.Unmarshal(m.Data, nirq)
		if err != nil {
			logrus.Errorf("Couldn't unmarshal message. Err : %v", err)
			return
		}
		logrus.Infof("Received a message. Guid : %s", nirq.Guid)

		nirs := new(msg.NewsInfoResponseContainer)
		nirs.Guid = nirq.Guid

		var ni NewsDbItem
		err = conn.Get(&ni, "select * from news where id=?", nirq.Id)
		if err != nil {
			logrus.Errorf("Couldn't get item with id=%v. Err : %v", nirq.Id, err)
			return
		}

		nirs.Item = &msg.NewsItem{Header: ni.Header, Date: ptypes.TimestampNow()}
		nirsData, err := proto.Marshal(nirs)
		if err != nil {
			logrus.Errorf("Couldn't marshal response. Err : %v", err)
			return
		}
		natsClient.Publish("info_response", nirsData)
	})

}

func startServer(cfg stconfig.StorageConfig) error {
	conn, err := sqlx.Open("mysql", cfg.DBSettings.ConnectionString)
	if err != nil {
		return err
	}

	natsClient, err := commons.ConnectToStan(cfg.NATSSettings)
	if err != nil {
		return err
	}
	defer natsClient.Close()

	regSub, err := newRegisterSubscribe(natsClient, conn)
	if err != nil {
		return err
	}
	defer func() {
		err := regSub.Unsubscribe()
		logrus.Infof("Register unsubscribe err : %v", err)
	}()

	infSub, err := newInfoSubscribe(natsClient, conn)
	if err != nil {
		return err
	}
	defer func() {
		err := infSub.Unsubscribe()
		logrus.Infof("Info unsubscribe err : %v", err)
	}()

	//todo remove this and add normal service %)
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			// Do not unsubscribe a durable on exit, except if asked to.
			cleanupDone <- true
		}
	}()
	<-cleanupDone
	return nil
}

func main() {
	const startEchoMaxAttemptsCount = 5
	vcpPtr := flag.String("c", "", "set path to config file")
	flag.Parse()
	cfgPaths := []string{".", "..", *vcpPtr}
	conf, err := stconfig.Get("", cfgPaths...)

	if err != nil {
		fmt.Printf("Couldn't get config for client. Err : %v", err)
		return
	}

	echoServerStarted := false
	for i := 0; i < startEchoMaxAttemptsCount && !echoServerStarted; i++ {
		time.Sleep(time.Second * 1)
		err = startServer(conf)
		echoServerStarted = err == nil
		fmt.Printf("Starting server. Attempt : %d, res : %v", i, err)
	}

	if !echoServerStarted {
		fmt.Printf("Couldn't start server after %d attempts. Check settings", startEchoMaxAttemptsCount)
	}
}
