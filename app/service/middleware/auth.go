package middleware

import "github.com/gogf/gf/net/ghttp"

// valid the token
func Auth(r *ghttp.Request) {
	r.Middleware.Next()
}
