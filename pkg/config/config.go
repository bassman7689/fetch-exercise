package config

import "github.com/spf13/viper"

type Config struct {
	ListenAddr string `mapstructure:"LISTEN_ADDR"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func Load() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	err := viper.Unmarshal(config)
	return config, err
}
