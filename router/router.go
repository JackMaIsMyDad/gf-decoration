package router

import (
	"gf-decoration/app/api/user_controller"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func bindTouter() {
	baseUrl := g.Config().GetString("base-url")
	s := g.Server()
	s.Group(baseUrl+"/user", func(g *ghttp.RouterGroup) {
		ctlUser := new(user_controller.Controller)
		g.POST("/signup", ctlUser.SignUp)
		g.GET("/info", ctlUser.Info)
	})
}

func init() {
	glog.Info("--------- router start -------------")
	bindTouter()
	glog.Info("--------- router finish -------------")
}
