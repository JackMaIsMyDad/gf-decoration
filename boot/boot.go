package boot

import (
	"gf-decoration/app/api/user_controller"
	"gf-decoration/library/base"
	_ "gf-decoration/packed"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
)

func init() {
	initGToken()
}

func initGToken() {
	base.GfToken = &gtoken.GfToken{
		CacheMode: g.Cfg().GetInt8("gToken.CacheMode"),
		//CacheKey:         g.Cfg().GetString("gToken.CacheKey"),
		Timeout:          g.Cfg().GetInt("gToken.Timeout"),
		MaxRefresh:       g.Cfg().GetInt("gToken.MaxRefresh"),
		TokenDelimiter:   g.Cfg().GetString("gToken.TokenDelimiter"),
		EncryptKey:       g.Cfg().GetBytes("gToken.EncryptKey"),
		AuthFailMsg:      g.Cfg().GetString("gToken.AuthFailMsg"),
		MultiLogin:       g.Cfg().GetBool("gToken.MultiLogin"),
		LoginPath:        "/login",
		LoginBeforeFunc:  user_controller.Login,
		LogoutPath:       "/user/logout",
		LogoutBeforeFunc: user_controller.Logout,
		AuthPaths:        g.SliceStr{"/user", "/system"},
		AuthExcludePaths: g.SliceStr{"/user/signup"},
	}
	base.GfToken.Start()
}
