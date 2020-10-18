package router

import (
	"gf-decoration/app/api/user_controller"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func bindRouter() {
	baseUrl := g.Config().GetString("base-url")
	s := g.Server()
	s.Group(baseUrl+"/api", func(g *ghttp.RouterGroup) {
		g.Group("/user", func(g *ghttp.RouterGroup) {
			ctlUser := new(user_controller.Controller)
			g.POST("/signup", ctlUser.SignUp)
			g.GET("/info", ctlUser.Info)
		})
		g.Group("/portal", func(g *ghttp.RouterGroup) {

		})
	})
}

func init() {
	glog.Info("--------- router start -------------")
	bindRouter()
	glog.Info("--------- router finish -------------")
}
