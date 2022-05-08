package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gsession"
	_ "simple_chat_room/configs"
	_ "simple_chat_room/controller"
	"simple_chat_room/middleware"
)

func main() {
	s := g.Server()
	s.SetSessionStorage(gsession.NewStorageFile("./gsessions"))
	s.BindMiddleware("/*", middleware.Cors)
	s.Run()
}
