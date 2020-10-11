package router

import (
	"gf-decoration/app/api/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlUser := new(user.Controller)
		group.POST("/login", ctlUser, "Login")
		//group.Middleware(middleware.Auth)
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.GET("/info", ctlUser, "Info")
			group.GET("/logout", ctlUser, "logout")
		})
	})
}
