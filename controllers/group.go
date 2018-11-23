package controllers

import (
	"github.com/astaxie/beego"
	"gofa/models"
)

type GroupController struct {
	beego.Controller
}

// @Title userlist
// @Description 通过groupID获取人物列表
// @Param	group_id		query 	string	true		"The group_id for search"
// @Success 200 {array} ["uid1","uid2"]
// @Failure 403 uid is empty
// @router /userlist [post]
func (g *GroupController) UserList() {
	groupID := g.GetString("group_id")

	list, err := models.GetUserListByGroupID(groupID)
	if err != nil {
		g.Data["json"] = err.Error()
	} else {
		g.Data["json"] = list
	}

	g.ServeJSON()
}
