package dao

import (
	"gorm.io/gorm"
	"zj/app/config"
	"zj/app/model"
	"zj/pkg/db"
)

type Dao struct {
	db *gorm.DB
}

func (d *Dao) Init() {
	d.db.Set("gorm:table_options", "CHARSET=utf8  AUTO_INCREMENT=1     COMMENT='用户';").AutoMigrate(&model.User{})
}

func New(cfg *config.Config) *Dao {
	d := &Dao{
		db: db.NewDb(cfg.Mysql),
	}
	d.Init()
	return d
}
