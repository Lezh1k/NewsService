package clconfig

import (
	"strings"

	natssettings "github.com/Lezh1k/NewsService/commons/structures"
	"github.com/spf13/viper"
)

type ServerSettings struct {
	EchoAddress string `json:"echo_addr" mapstructure:"echo_addr"`
}

type ClientConfig struct {
	ServerSettings ServerSettings            `json:"server_settings" mapstructure:"server_settings"`
	NATSSettings   natssettings.NATSSettings `json:"nats_settings" mapstructure:"nats_settings"`
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
