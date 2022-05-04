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
	utils.ClientMap.Range(func(key, value any) bool {
		w := value.(*utils.Client).Writer
		w.WriteString(utils.OKChunked(request.Session.MustId(), msg))
		w.Flush()
		return true
	})
}

func register(request *ghttp.Request) {
	client, _ := utils.ClientMap.Load(request.Session.MustId())
	if client == nil {
		request.Response.WriteExit("fail")
	}
	hj, w, _ := request.Response.Hijack()
	c := client.(*utils.Client)
	c.Writer = w
	c.Conn = &hj
	c.ClientId = request.Session.MustId()
	w.WriteString("HTTP/1.1 200 OK\r\n")
	w.WriteString("Connection: keep-alive\r\n")
	w.WriteString(utils.CorsHeader(request.Header.Get("Origin")))
	for k, v := range headers {
		w.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	w.WriteString("\r\n")
	w.Flush()
}

func quit(request *ghttp.Request) {
	id := request.Session.MustId()
	client, ok := utils.ClientMap.Load(id)
	if !ok {
		return
	}
	c := client.(*utils.Client)
	c.Writer.WriteString(utils.EndChunked())
	c.Writer.Flush()
	(*c.Conn).Close()
	utils.ClientMap.Delete(id)
}
