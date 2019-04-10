package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

// NATSSettings represents NATS streaming client connection options
type NATSSettings struct {
	ClientID           string        `json:"client_id" mapstructure:"client_id"`
	ClusterID          string        `json:"cluster_id" mapstructure:"cluster_id"`
	NATSURL            string        `json:"nats_url" mapstructure:"nats_url"`                           // NATSURL is the default URL the client connects to
	DiscoverPrefix     string        `json:"discover_prefix" mapstructure:"discover_prefix"`             // DiscoverPrefix is the prefix subject used to connect to the NATS Streaming server
	DefaultACKPrefix   string        `json:"default_ack_prefix" mapstructure:"default_ack_prefix"`       // ACKPrefix is the prefix subject used to send ACKs to the NATS Streaming server
	ConnectTimeout     time.Duration `json:"connect_timeout" mapstructure:"connect_timeout"`             // ConnectWait is the default timeout used for the connect operation
	AckTimeout         time.Duration `json:"ack_timeout" mapstructure:"ack_timeout"`                     // AckWait indicates how long the server should wait for an ACK before resending a message
	MaxPubAcksInflight int           `json:"max_pub_ack_in_flight" mapstructure:"max_pub_ack_in_flight"` // MaxInflight indicates how many messages with outstanding ACKs the server can send
	PingIterval        int           `json:"ping_interval" mapstructure:"ping_interval"`                 // PingInterval is the default interval (in seconds) at which a connection sends a PING to the server
	PingMaxOut         int           `json:"ping_max_out" mapstructure:"ping_max_out"`                   // PingMaxOut is the number of PINGs without a response before the connection is considered lost.
}

type ClientConfig struct {
	NATSSettings NATSSettings `json:"nats_settings" mapstructure:"nats_settings"`
}

// Get parses ClientConfig from env vars
func Get(filePath string, additionalPaths ...string) (ClientConfig, error) {
	var cfg ClientConfig
	var err error

	if filePath == "" {
		viper.SetConfigName("client")
		for _, path := range additionalPaths {
			viper.AddConfigPath(path)
		}
	} else {
		viper.SetConfigFile(filePath)
	}

	variables := []string{
		"nats_connect_addr",
	}

	for _, v := range variables {
		_ = viper.BindEnv(v) // will return error if only len(v) == 0
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	err = viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}
	err = viper.Unmarshal(&cfg)
	return cfg, err
}
