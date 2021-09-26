package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zj/app/api"
	"zj/app/service"
	"zj/pkg/web"
)

var srv *service.Service

func StartHttpServer(s *service.Service) *http.Server {
	srv = s
	api := api.New(s)
	if !s.C.Http.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.Use()
	g.GET("/ping", func(c *gin.Context) {
		web.Json(c, gin.H{
			"message": "pong",
		})
	})

	initRoute(g, api)

	srv := &http.Server{
		Addr:    s.C.Http.Port,
		Handler: g,
	}
	return srv
}

func initRoute(g *gin.Engine, api *api.Api) {
	e := g.Group("/zj/user")
	{
		e.POST("/login", api.Login)
		e.POST("/createUser", api.CreateUser)
	}
}
