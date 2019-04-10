package clconfig

import (
	"strings"

	"github.com/spf13/viper"
)

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
