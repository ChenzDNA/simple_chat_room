package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func Cors(r *ghttp.Request) {
	if !r.IsFileRequest() {
		r.Response.CORSDefault()
	}
	r.Middleware.Next()
}
