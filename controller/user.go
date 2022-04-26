package controller

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"simple_chat_room/utils"
)

func init() {
	group := g.Server().Group("/user")
	group.POST("/getid", genUser)
	group.POST("/hearbeat", heartBeat)
}

func genUser(request *ghttp.Request) {
	s, _ := request.Session.Get("id")
	id := request.Session.MustId()
	if s == nil {
		request.Session.Set("id", id)
	}
	request.Session.Close()
	utils.ClientMap[id] = &utils.Client{ClientId: id}
	request.Response.WriteJsonExit(utils.OK("ok", id))
}

func heartBeat(request *ghttp.Request) {

	request.Response.WriteJsonExit(utils.OK("ok", nil))
}
