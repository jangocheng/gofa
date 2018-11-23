package controllers

import (
	"github.com/astaxie/beego"
	"gofa/entities"
	"gofa/models"
	"strings"
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

// @Title addface
// @Description 向祖中添加face
// @Param	group_id	query 	string	true		"groupID"
// @Param	image		query 	string	true		"image"
// @Param	image_type	query	string	true 		"base64/url/face_token"
// @Param	user_id		query	string	true 		"user name"
// @Success 200 {string} "2fa64a88a9d5118916f9a303782a97d3"
// @Failure 403 uid is empty
// @router /addface [post]
func (g *GroupController) AddFace() {
	gStorage := entities.AddFaceStorage{
		Image:     g.GetString("image"),
		ImageType: strings.ToUpper(g.GetString("image_type")),
		GroupID:   g.GetString("group_id"),
		UserID:    g.GetString("user_id"),
	}
	if err := g.ParseForm(&gStorage); err != nil {
		g.Data["json"] = "add error"
	}

	re, err := models.AddUserImageToGroup(gStorage)
	if err != nil {
		g.Data["json"] = err.Error()
	} else {
		g.Data["json"] = re
	}

	g.ServeJSON()
}
