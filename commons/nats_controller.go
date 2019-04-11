package commons

import (
	"time"

	nats "github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/sirupsen/logrus"
)

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

//ConnectToStan ...
func ConnectToStan(cfg NATSSettings) (stan.Conn, error) {
	var natsClient stan.Conn
	natsConn, err := nats.Connect(
		cfg.NATSURL,
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

	natsClient, err = stan.Connect(
		cfg.ClusterID,
		cfg.ClientID,
		stan.NatsURL(cfg.NATSURL),
		stan.ConnectWait(cfg.ConnectTimeout),
		stan.PubAckWait(cfg.AckTimeout),
		stan.MaxPubAcksInflight(cfg.MaxPubAcksInflight),
		stan.NatsConn(natsConn),
	)

	if err != nil {
		logrus.Error("unable to connect to nats streaming server", err)
		return natsClient, err
	}

	logrus.WithFields(logrus.Fields{"URL": cfg.NATSURL}).Info("Connected to NATS")
	return natsClient, nil
}
