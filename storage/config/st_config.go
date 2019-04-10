package stconfig

import (
	"strings"

	natssettings "github.com/Lezh1k/NewsService/commons/structures"
	"github.com/spf13/viper"
)

// DBSettings represents db options
type DBSettings struct {
	ConnectionString string `json:"connection_str" mapstructure:"connection_str"`
}

type StorageConfig struct {
	NATSSettings natssettings.NATSSettings `json:"nats_settings" mapstructure:"nats_settings"`
	DBSettings   DBSettings                `json:"db_settings" mapstructure:"db_settings"`
}

// Get parses ClientConfig from env vars
func Get(filePath string, additionalPaths ...string) (StorageConfig, error) {
	var cfg StorageConfig
	var err error

	if filePath == "" {
		viper.SetConfigName("storage")
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
