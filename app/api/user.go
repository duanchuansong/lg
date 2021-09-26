package api

import (
	"github.com/gin-gonic/gin"
	"zj/app/model"
	"zj/pkg/web"
)

func (a *Api) Login(c *gin.Context) {
	user := new(model.User)
	if err := c.ShouldBind(user); err != nil {
		web.Json(c, err)
		return
	}
	data, err := a.s.Login(user)
	if err != nil {
		web.Json(c, err)
		return
	}
	web.Json(c, data)
}

func (a *Api) CreateUser(c *gin.Context) {
	user := new(model.User)
	if err := c.ShouldBind(user); err != nil {
		web.Json(c, err)
		return
	}
	data, err := a.s.CreateUser(a.c, user)
	if err != nil {
		web.Json(c, err)
		return
	}
	web.Json(c, data)
}
