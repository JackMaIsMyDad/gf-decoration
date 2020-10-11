package boot

import (
	_ "gf-decoration/packed"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
)

func init() {
	initGToken()
}

func initGToken() {
	AdminGfToken = &gtoken.GfToken{
		CacheMode:        g.Cfg().GetInt8("gToken.CacheMode"),
		CacheKey:         g.Cfg().GetString("gToken.CacheKey"),
		Timeout:          g.Cfg().GetInt("gToken.Timeout"),
		MaxRefresh:       g.Cfg().GetInt("gToken.MaxRefresh"),
		TokenDelimiter:   g.Cfg().GetString("gToken.TokenDelimiter"),
		EncryptKey:       g.Cfg().GetBytes("gToken.EncryptKey"),
		AuthFailMsg:      g.Cfg().GetString("gToken.AuthFailMsg"),
		MultiLogin:       g.Cfg().GetBool("gToken.MultiLogin"),
		LoginPath:        "/login",
		LoginBeforeFunc:  service.AdminLogin,
		LogoutPath:       "/user/logout",
		AuthPaths:        g.SliceStr{"/user/*"},
		AuthExcludePaths: g.SliceStr{"login"},
	}
	AdminGfToken.Start()
}
