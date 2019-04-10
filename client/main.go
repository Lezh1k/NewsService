package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Lezh1k/NewsService/client/config"

	nats "github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
)

type registerRequest struct {
	Created string `json:"created"`
	Header  string `json:"header"`
}

func disconnectHandler(conn *nats.Conn) {
	logrus.WithFields(
		logrus.Fields{
			"isReconnecting": conn.IsReconnecting(),
			"isConnected":    conn.IsConnected(),
			"lastError":      conn.LastError(),
		}).Warn("disconnected from nats server")
}

func reconnectHandler(conn *nats.Conn) {
	logrus.WithFields(
		logrus.Fields{
			"isReconnecting": conn.IsReconnecting(),
			"isConnected":    conn.IsConnected(),
			"lastError":      conn.LastError(),
		}).Warn("reconnected to nats server")
}

func closedHandler(conn *nats.Conn) {
	logrus.WithFields(
		logrus.Fields{
			"isReconnecting": conn.IsReconnecting(),
			"isConnected":    conn.IsConnected(),
			"lastError":      conn.LastError(),
		}).Warn("connection closed to nats server")
}

// errorHandler is used to process asynchronous errors encountered
// while processing inbound messages.
func errorHandler(conn *nats.Conn, sub *nats.Subscription, err error) {
	logrus.WithFields(
		logrus.Fields{
			"isConnected": conn.IsConnected(),
			"subject":     sub.Subject,
			"queue":       sub.Queue,
			"error":       err.Error(),
		}).Error("encountered error")
}

func connectToStan(cfg clconfig.ClientConfig) (stan.Conn, error) {
	var natsClient stan.Conn
	natsConn, err := nats.Connect(
		cfg.NATSSettings.NATSURL,
		nats.MaxReconnects(20),
		nats.Timeout(time.Second*2),
	)

	if err != nil {
		return natsClient, err
	}
	natsConn.SetReconnectHandler(reconnectHandler)
	natsConn.SetDisconnectHandler(disconnectHandler)
	natsConn.SetClosedHandler(closedHandler)
	natsConn.SetErrorHandler(errorHandler)

	//store nats connection
	natsClient, err = stan.Connect(
		cfg.NATSSettings.ClusterID,
		cfg.NATSSettings.ClientID,
		stan.NatsURL(cfg.NATSSettings.NATSURL),
		stan.ConnectWait(cfg.NATSSettings.ConnectTimeout),
		stan.PubAckWait(cfg.NATSSettings.AckTimeout),
		stan.MaxPubAcksInflight(cfg.NATSSettings.MaxPubAcksInflight),
		stan.NatsConn(natsConn),
	)

	if err != nil {
		logrus.Error("unable to connect to nats streaming server", err)
		return natsClient, err
	}

	logrus.WithFields(logrus.Fields{"URL": cfg.NATSSettings.NATSURL}).Info("Connected to NATS")
	return natsClient, nil
}

func startServer(cfg clconfig.ClientConfig) error {

	natsClient, err := connectToStan(cfg)
	if err != nil {
		return err
	}

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/request_news_item", func(c echo.Context) error {
		newItemID := c.QueryParam("id")
		if strings.TrimSpace(newItemID) == "" {
			return c.String(http.StatusBadRequest, "Empty ID in URL is not allowed")
		}

		err := natsClient.Publish("request", []byte(newItemID))
		if err != nil {
			logrus.Errorf("Publish error : %v", err)
			return c.String(http.StatusInternalServerError, "Couldn't publish message to nats")
		}

		var wg sync.WaitGroup
		wg.Add(1)
		sub, _ := natsClient.Subscribe("request", func(m *stan.Msg) {
			logrus.Infof("Received a message: %s\n", string(m.Data))
			wg.Done()
		})
		wg.Wait()
		// Unsubscribe
		sub.Unsubscribe()

		return c.String(http.StatusOK, newItemID)
	})

	e.POST("/register_news_item", func(c echo.Context) error {
		regReq := new(registerRequest)
		if err := c.Bind(regReq); err != nil {
			fmt.Printf("Invalid request JSON received. Err : %v", err)
			return c.String(http.StatusBadRequest, "Invalid JSON")
		}
		return c.String(http.StatusOK, fmt.Sprintf("%s\n%s\n", regReq.Created, regReq.Header))
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
		err = startServer(conf)
		echoServerStarted = err == nil
		fmt.Printf("Starting server. Attempt : %d, res : %v", i, err)
	}

	if !echoServerStarted {
		fmt.Printf("Couldn't start server after %d attempts. Check settings", startEchoMaxAttemptsCount)
	}
}
