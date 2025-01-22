package controller

import (
	"github.com/beego/beego"
	"net/http"
	"note-beego/src/repository"
)

type CommandController struct {
	beego.Controller
}

// One 根据id查询命令
func (c *CommandController) One() {
	command := repository.One(*c.Ctx)
	c.Data["json"] = command
	c.ServeJSON()
}

// List 查询命令列表
func (c *CommandController) List() {
	commandArray := repository.List(*c.Ctx)
	c.Data["json"] = commandArray
	c.ServeJSON()
}

// Insert 新增命令
func (c *CommandController) Insert() {
	result, commandName := repository.Insert(*c.Ctx)
	err := c.Ctx.Output.JSON(map[string]interface{}{
		"result":  result,
		"command": commandName,
	}, true, false)
	if err != nil {
		c.Ctx.Abort(http.StatusInternalServerError, err.Error())
	}
	c.ServeJSON()
}

// InsertBatch 批量新增命令
func (c *CommandController) InsertBatch() {
	result := repository.InsertBatch(*c.Ctx)
	c.Data["json"] = result
	c.ServeJSON()
}

// Update 修改命令
func (c *CommandController) Update() {
	result := repository.Update(*c.Ctx)
	c.Data["json"] = result
	c.ServeJSON()
}

// UpdateBatch 批量修改命令
func (c *CommandController) UpdateBatch() {
	result := repository.UpdateBatch(*c.Ctx)
	c.Data["json"] = result
	c.ServeJSON()
}

// Delete 删除命令
func (c *CommandController) Delete() {
	result, objectId := repository.Delete(*c.Ctx)
	err := c.Ctx.Output.JSON(map[string]interface{}{
		"result": result,
		"_id":    objectId,
	}, true, false)
	if err != nil {
		c.Ctx.Abort(http.StatusInternalServerError, err.Error())
	}
	c.ServeJSON()
}

// DeleteBatch 批量删除命令
func (c *CommandController) DeleteBatch() {
	result, objectIds := repository.DeleteBatch(*c.Ctx)
	err := c.Ctx.Output.JSON(map[string]interface{}{
		"result": result,
		"_ids":   objectIds,
	}, true, false)
	if err != nil {
		c.Ctx.Abort(http.StatusInternalServerError, err.Error())
	}
	c.ServeJSON()
}

// Select 查询命令
func (c *CommandController) Select() {
	command := repository.Select(*c.Ctx)
	c.Data["json"] = command
	c.ServeJSON()
}

// NameList 查询命令名称列表
func (c *CommandController) NameList() {
	nameArray := repository.NameList(*c.Ctx)
	c.Data["json"] = nameArray
	c.ServeJSON()
}
