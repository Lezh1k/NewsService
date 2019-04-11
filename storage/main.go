package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/golang/protobuf/ptypes"

	"github.com/Lezh1k/NewsService/commons"
	msg "github.com/Lezh1k/NewsService/commons/pb"
	"github.com/Lezh1k/NewsService/storage/config"
	"github.com/gogo/protobuf/proto"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/sirupsen/logrus"
)

func newRegisterSubscribe(natsClient stan.Conn) error {
	sub, err := natsClient.Subscribe("register_req", func(m *stan.Msg) {
		nrrq := new(msg.NewsRegisterRequestContainer)
		err := proto.Unmarshal(m.Data, nrrq)
		if err != nil {
			logrus.Errorf("Couldn't unmarshal message. Err : %v", err)
			return
		}
		logrus.Infof("Received a message. Guid : %s", nrrq.Guid)

		nrrs := new(msg.NewsRegisterResponseContainer)
		nrrs.Guid = nrrq.Guid
		nrrs.Id = "ok" // todo get from db

		nrrsData, err := proto.Marshal(nrrs)
		if err != nil {
			logrus.Errorf("Couldn't marshal response. Err : %v", err)
			return
		}
		natsClient.Publish("register_resp", nrrsData)
	})
	if err != nil {
		return err
	}
	defer func() {
		err := sub.Unsubscribe()
		logrus.Infof("Unsubscribe err : %v", err)
	}()
	return nil
}

func newInfoSubscribe(natsClient stan.Conn) error {
	sub, err := natsClient.Subscribe("info_req", func(m *stan.Msg) {
		nirq := new(msg.NewsInfoRequestContainer)
		err := proto.Unmarshal(m.Data, nirq)
		if err != nil {
			logrus.Errorf("Couldn't unmarshal message. Err : %v", err)
			return
		}
		logrus.Infof("Received a message. Guid : %s", nirq.Guid)

		nirs := new(msg.NewsInfoResponseContainer)
		nirs.Guid = nirq.Guid
		//todo get item from bd
		nirs.Item = &msg.NewsItem{Header: "stub", Date: ptypes.TimestampNow()}

		nirsData, err := proto.Marshal(nirs)
		if err != nil {
			logrus.Errorf("Couldn't marshal response. Err : %v", err)
			return
		}
		natsClient.Publish("info_resp", nirsData)
	})
	if err != nil {
		return err
	}
	defer func() {
		err := sub.Unsubscribe()
		logrus.Infof("Unsubscribe err : %v", err)
	}()
	return nil
}

func startServer(cfg stconfig.StorageConfig) error {
	natsClient, err := commons.ConnectToStan(cfg.NATSSettings)
	if err != nil {
		return err
	}
	defer natsClient.Close()
	err = newRegisterSubscribe(natsClient)
	if err != nil {
		return err
	}

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
		err = startServer(conf)
		echoServerStarted = err == nil
		fmt.Printf("Starting server. Attempt : %d, res : %v", i, err)
	}

	if !echoServerStarted {
		fmt.Printf("Couldn't start server after %d attempts. Check settings", startEchoMaxAttemptsCount)
	}
}
