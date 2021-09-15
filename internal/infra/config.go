package infra

import "github.com/spf13/viper"

type AppConfig struct {
	DSN         string
	HttpAddr    string
	HostName    string
	Development bool
	ShowSql     bool
}

func NewConfig() (*AppConfig, error) {
	viper.SetDefault("DSN", "host=localhost user=postgres password=postgres dbname=flow port=5432 sslmode=disable TimeZone=Europe/Moscow")
	viper.SetDefault("HostName", "localhost:8080")
	viper.SetDefault("HttpAddr", ":8080")
	viper.SetDefault("Development", true)
	viper.SetDefault("ShowSql", true)

	_ = viper.BindEnv("DSN")
	_ = viper.BindEnv("HostName")
	_ = viper.BindEnv("HttpAddr")
	_ = viper.BindEnv("Development")
	_ = viper.BindEnv("ShowSql")

	config := &AppConfig{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
