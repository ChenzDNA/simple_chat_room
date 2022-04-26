package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

func init() {
	s := g.Server()
	s.BindHookHandler("/**", ghttp.HookBeforeServe, func(r *ghttp.Request) {
		if r.IsFileRequest() && strings.HasSuffix(r.StaticFile.Path, ".js") {
			r.Response.Header().Set("Content-Type", "text/javascript;charset=utf-8")
		}
	})
	s.BindHookHandler("/chat/*", ghttp.HookAfterServe, func(r *ghttp.Request) {
		r.Response.Header().Set("Connection", "keep-alive")
	})
}
