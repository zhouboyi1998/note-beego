package router

import (
	"github.com/beego/beego"
	"note-beego/src/controller"
)

// init 初始化路由规则
func init() {
	// 新建命令相关路由命名空间
	ns := beego.NewNamespace("/command",
		// 添加命令相关路由
		beego.NSRouter("/:commandId", &controller.CommandController{}, "get:One"),
		beego.NSRouter("/", &controller.CommandController{}, "get:List"),
		beego.NSRouter("/", &controller.CommandController{}, "post:Insert"),
		beego.NSRouter("/batch", &controller.CommandController{}, "post:InsertBatch"),
		beego.NSRouter("/", &controller.CommandController{}, "put:Update"),
		beego.NSRouter("/batch", &controller.CommandController{}, "put:UpdateBatch"),
		beego.NSRouter("/:commandId", &controller.CommandController{}, "delete:Delete"),
		beego.NSRouter("/batch", &controller.CommandController{}, "delete:DeleteBatch"),
		beego.NSRouter("/select/:commandName", &controller.CommandController{}, "get:Select"),
		beego.NSRouter("/name-list", &controller.CommandController{}, "get:NameList"),
	)
	// 加载命名空间
	beego.AddNamespace(ns)
}
