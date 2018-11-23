package controllers

import (
	"github.com/astaxie/beego"
)

type GroupController struct {
	beego.Controller
}

// @Title userlist
// @Description 通过groupID获取人物列表
// @Param	group_id		query 	string	true		"The group_id for search"
// @Success 200 {array} login success
// @router /userlist [get]
func (g *GroupController) UserList() {
	groupID := g.GetString("group_id")

	g.Data["json"] = groupID
	g.ServeJSON()
}
