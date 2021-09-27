package config

import (
	"github.com/duanchuansong/stk/config"
	"github.com/duanchuansong/stk/db"
	"github.com/duanchuansong/stk/web"
	"github.com/duanchuansong/stk/xlog"
)

type Config struct {
	Mysql  *db.Mysql
	Http   *web.Http
	Static *web.Static
	Log    *xlog.Config
}

func New() *Config {
	config.Init("config.toml")
	return &Config{
		Mysql: &db.Mysql{
			Host:         config.GetString("mysql.host"),
			Port:         config.GetString("mysql.port"),
			User:         config.GetString("mysql.username"),
			Pass:         config.GetString("mysql.password"),
			DbName:       config.GetString("mysql.dbname"),
			MaxIdleConns: config.GetInt("mysql.maxIdleConns"),
			MaxOpenConns: config.GetInt("mysql.maxOpenConns"),
		},
		Http: &web.Http{
			Port:  config.GetString("http.port"),
			Debug: config.GetBool("http.debug"),
		},
		Static: &web.Static{
			Port: config.GetString("static.port"),
			Dir:  config.GetString("static.dir"),
		},
		Log: &xlog.Config{
			LogPath:     config.GetString("log.logPath"),
			LogFileName: config.GetString("log.logFileName"),
			LogLevel:    config.GetString("log.logLevel"),
			LogOutFile:  config.GetBool("log.logOutFile"),
		},
	}
}
