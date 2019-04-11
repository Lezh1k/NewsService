package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Lezh1k/NewsService/client/config"
	"github.com/Lezh1k/NewsService/commons"
	msg "github.com/Lezh1k/NewsService/commons/pb"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"

	stan "github.com/nats-io/go-nats-streaming"
)

type registerRequest struct {
	Created string `json:"created"`
	Header  string `json:"header"`
}

func handleNewsInfoRequest(natsClient stan.Conn, c echo.Context) error {
	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return c.String(http.StatusBadRequest, "Empty ID in URL is not allowed")
	}

	nirqc := new(msg.NewsInfoRequestContainer)
	nirqc.Guid = uuid.New().String()
	nirqc.Id = id
	nirqcData, err := proto.Marshal(nirqc)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Marshaling error")
	}

	var wg sync.WaitGroup
	wg.Add(1)

	hdr := ""

	sub, err := natsClient.Subscribe("info_response", func(m *stan.Msg) {
		nirsc := new(msg.NewsInfoResponseContainer)
		err := proto.Unmarshal(m.Data, nirsc)
		if err != nil {
			logrus.Errorf("Unmarshalling error : %v", err)
			return
		}
		if nirsc.Guid != nirqc.Guid {
			return //just ignore it
		}
		hdr = nirsc.Item.Header
		wg.Done()
	})

	if err != nil {
		logrus.Errorf("Couldn't subscribe request. Err : %v", err)
		return c.String(http.StatusInternalServerError, "Subscribtion failed")
	}
	defer func() {
		err := sub.Unsubscribe()
		logrus.Infof("Unsubscribe err : %v", err)
	}()

	err = natsClient.Publish("info_request", nirqcData)
	if err != nil {
		logrus.Errorf("Publish error : %v", err)
		return c.String(http.StatusInternalServerError, "Couldn't publish message to nats")
	}

	wg.Wait()

	//todo check that hdr != ""
	return c.String(http.StatusOK, hdr)
}

func handleNewsRegisterRequest(natsClient stan.Conn, c echo.Context) error {
	regReq := new(registerRequest)
	if err := c.Bind(regReq); err != nil {
		fmt.Printf("Invalid request JSON received. Err : %v", err)
		return c.String(http.StatusBadRequest, "Invalid JSON")
	}

	var registerDate *timestamp.Timestamp
	if strings.TrimSpace(regReq.Created) == "" {
		registerDate = ptypes.TimestampNow()
	} else {
		const layout = "2006-01-02T15:04:05.000Z"
		t, err := time.Parse(layout, regReq.Created)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid time int request. Follow this layout : "+layout)
		}
		registerDate, err = ptypes.TimestampProto(t)
		if err != nil {
			logrus.Errorf("Couldn't convert time to timestamp. Err : %v", err)
			return c.String(http.StatusInternalServerError, "Something wrong with timestamp")
		}
	}

	ni := new(msg.NewsItem)
	ni.Header = regReq.Header
	ni.Date = registerDate

	nrrq := new(msg.NewsRegisterRequestContainer)
	nrrq.Guid = uuid.New().String()
	nrrq.Item = ni

	publishData, err := proto.Marshal(nrrq)

	if err != nil {
		logrus.Errorf("Couldn't marshal NewsItem. Err : %v", err)
		return c.String(http.StatusInternalServerError, "Couldn't marshal NewsItem")
	}
	id := ""

	var wg sync.WaitGroup
	wg.Add(1)
	sub, err := natsClient.Subscribe("register_resp", func(m *stan.Msg) {
		nrrs := new(msg.NewsRegisterResponseContainer)
		err := proto.Unmarshal(m.Data, nrrs)
		if err != nil {
			logrus.Errorf("Couldn't unmarshal register response. Err : %v", err)
			return
		}
		if nrrs.Guid != nrrq.Guid { //just ignore this message.
			return
		}
		id = nrrs.Id
		wg.Done()
	})

	if err != nil {
		logrus.Errorf("Couldn't subscribe. Err : %v", err)
		return c.String(http.StatusInternalServerError, "Subscribtion failed")
	}
	defer sub.Unsubscribe()

	err = natsClient.Publish("register_req", []byte(publishData))
	if err != nil {
		logrus.Errorf("Publish error : %v", err)
		return c.String(http.StatusInternalServerError, "Couldn't publish message to nats")
	}
	wg.Wait()
	//todo check that id != ""
	return c.String(http.StatusOK, id)
}

func startServer(cfg clconfig.ClientConfig) error {

	natsClient, err := commons.ConnectToStan(cfg.NATSSettings)
	if err != nil {
		return err
	}
	defer natsClient.Close()

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/news/request", func(c echo.Context) error {
		return handleNewsInfoRequest(natsClient, c)
	})

	e.POST("/news/register", func(c echo.Context) error {
		return handleNewsRegisterRequest(natsClient, c)
	})

	// Start server
	return e.Start(cfg.ServerSettings.EchoAddress)
}

func main() {
	const startEchoMaxAttemptsCount = 5
	vcpPtr := flag.String("c", "", "set path to config file")
	flag.Parse()
	cfgPaths := []string{".", "..", *vcpPtr}
	conf, err := clconfig.Get("", cfgPaths...)

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
