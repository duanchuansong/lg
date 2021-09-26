package service

import (
	"zj/app/config"
	"zj/app/dao"
)

type Service struct {
	Dao *dao.Dao
	C   *config.Config
}

func New(cfg *config.Config) *Service {

	return &Service{
		C:   cfg,
		Dao: dao.New(cfg),
	}
}
