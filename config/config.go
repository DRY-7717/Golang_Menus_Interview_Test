package config

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`

}

type PsqlDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	MaxIdle  int    `json:"max_idle"`
	MaxOpen  int    `json:"max_open"`
}

type Config struct {
	App  App
	Psql PsqlDB
}

func NewConfig() *Config {

	return &Config{
		App: App{
			AppEnv:       viper.GetString("APP_ENV"),
			AppPort:      viper.GetString("APP_PORT"),
		},

		Psql: PsqlDB{
			Host:     viper.GetString("DATABASE_HOST"),
			Port:     viper.GetString("DATABASE_PORT"),
			Name:     viper.GetString("DATABASE_NAME"),
			User:     viper.GetString("DATABASE_USER"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			MaxIdle:  viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
			MaxOpen:  viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
		},
	}
}
