package router

import (
	"github.com/beego/beego"
	"note-beego/src/controller"
)

// init 初始化路由规则
func init() {
	beego.Router("/hello", &controller.Controller{}, "*:Hello")
}
