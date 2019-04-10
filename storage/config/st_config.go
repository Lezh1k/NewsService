package stconfig

import (
	"github.com/spf13/viper"
)

// DBSettings represents db options
type DBSettings struct {
	ConnectionString string `json:"connection_str" mapstructure:"connection_str"`
}

type StorageConfig struct {
	DBSettings DBSettings `json:"db_settings" mapstructure:"db_settings"`
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

	err = viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}
	err = viper.Unmarshal(&cfg)
	return cfg, err
}
