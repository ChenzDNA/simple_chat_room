package controller

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"simple_chat_room/utils"
)

var headers = map[string]string{
	"Transfer-Encoding": "chunked",
	"Content-Type":      "text/html; charset=utf-8",
}

func init() {
	s := g.Server()
	group := s.Group("/chat")
	group.POST("/send", send)
	group.GET("/register", register)
	group.POST("/quit", quit)
}

func send(request *ghttp.Request) {
	msg := request.Get("msg").String()
	for _, v := range utils.ClientMap {
		w := *v.Writer
		w.WriteString(utils.OKChunked(request.Session.MustId(), msg))
		w.Flush()
	}
}

func register(request *ghttp.Request) {
	c := utils.ClientMap[request.Session.MustId()]
	if c == nil {
		request.Response.WriteExit("fail")
	}
	hj, w, _ := request.Response.Hijack()
	c.Writer = w
	c.Conn = &hj
	c.ClientId = request.Session.MustId()
	w.WriteString("HTTP/1.1 200 OK\n")
	w.WriteString("Connection: keep-alive\n")
	w.WriteString(utils.CorsHeader(request.Header.Get("Origin")))
	for k, v := range headers {
		w.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	w.WriteString("\n")
	w.Flush()
}

func quit(request *ghttp.Request) {
	id := request.Session.MustId()
	client := utils.ClientMap[id]
	if client == nil {
		return
	}
	client.Writer.WriteString(utils.EndChunked())
	client.Writer.Flush()
	(*client.Conn).Close()
	delete(utils.ClientMap, id)
}
