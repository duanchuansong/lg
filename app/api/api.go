package api

import (
	"context"
	"zj/app/service"
)

type Api struct {
	c context.Context
	s *service.Service
}

func New(s *service.Service) *Api {
	return &Api{
		s: s,
		c: context.Background(),
	}
}
