package config

import (
	"io/fs"

	"github.com/spf13/viper"
)

type Config struct {
	ListenAddr string `mapstructure:"LISTEN_ADDR"`
}

func Load() (*Config, error) {
	viper.AutomaticEnv()
	if err := viper.BindEnv("LISTEN_ADDR"); err != nil {
		return nil, err
	}


	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*fs.PathError); !ok {
			return nil, err
		}
	}


	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
