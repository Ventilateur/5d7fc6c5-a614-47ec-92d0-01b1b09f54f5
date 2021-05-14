package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mongo struct {
		ConnectionString string `mapstructure:"connection_string"`
		DatabaseName     string `mapstructure:"database_name"`
	} `mapstructure:"mongo"`
	JWT struct {
		SigningKey  string `mapstructure:"signing_key"`
		DurationMin int64  `mapstructure:"duration_min"`
	}
	UserDataDir string `mapstructure:"user_data_dir"`
	ListenPort  int64  `mapstructure:"listen_port"`
}

func GetConfig() (Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.BindEnv("jwt.signing_key", "JWT_SIGNING_KEY"); err != nil {
		return Config{}, err
	}

	if err := viper.BindEnv("mongo.connection_string", "MONGO_CONNECTION_STRING"); err != nil {
		return Config{}, err
	}

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
