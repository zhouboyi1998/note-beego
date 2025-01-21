package controller

import "github.com/beego/beego"

type Controller struct {
	beego.Controller
}

func (c *Controller) Hello() {
	c.Data["json"] = "Hello Beego"
	c.ServeJSON()
}
