package main

import (
	"github.com/beego/beego"
	_ "note-beego/src/router"
)

func main() {
	// 启动服务
	beego.Run()
}
