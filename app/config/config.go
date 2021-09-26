package config

import (
	"zj/pkg/config"
	"zj/pkg/db"
	"zj/pkg/web"
)

type Config struct {
	Mysql *db.Mysql
	Http  *web.Http
}

func New() *Config {
	config.Init("config.toml")
	return &Config{
		Mysql: &db.Mysql{
			Host:   config.GetString("mysql.host"),
			Port:   config.GetString("mysql.port"),
			User:   config.GetString("mysql.username"),
			Pass:   config.GetString("mysql.password"),
			DbName: config.GetString("mysql.dbname"),
		},
		Http: &web.Http{
			Port:  config.GetString("http.port"),
			Debug: config.GetBool("http.debug"),
		},
	}
}
