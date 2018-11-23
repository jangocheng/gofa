package models

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"gofa/entities"
	"gofa/resources"
)

var (
	UrlUserListByGroup = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getusers"
	UrlAddUserToGroup  = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add"
	AccessToken        = "?access_token="
)

func GetUserListByGroupID(groupID string) ([]string, error) {
	params, err := json.Marshal(entities.GetUsersListStorage{
		GroupID: groupID,
		Start:   0,
		Length:  100,
	})
	if err != nil {
		beego.Warning("GetUserListByGroupID err", err, groupID)
		return nil, err
	}

	// http 获取
	re, err := HttpPost(UrlUserListByGroup+AccessToken+resources.Token, string(params))
	if err != nil {
		beego.Warning("get userlist by group id  err :" + err.Error())
		return nil, err
	}

	userListResp := entities.UserListResp{}
	if err := json.Unmarshal(re, &userListResp); err != nil {
		beego.Warning("Unmarshal resp error [" + string(re))
		return nil, err
	}
	userList := []string{}
	ok := false
	if userList, ok = userListResp.Result["user_id_list"]; !ok || userListResp.ErrorCode != 0 {
		beego.Warning("call error :", re)
		return nil, errors.New("call err")
	}

	return userList, err

}

func AddUserImageToGroup(gStorage entities.AddFaceStorage) (string, error) {

	params, err := json.Marshal(gStorage)
	if err != nil {
		beego.Warning("AddUserImageToGroup err", err, gStorage)
		return "", err
	}

	re, err := HttpPost(UrlAddUserToGroup+AccessToken+resources.Token, string(params))
	if err != nil {
		beego.Warning("get userlist by group id  err :" + err.Error())
		return "", err
	}
	addResp := entities.AddFaceResp{}
	if err := json.Unmarshal(re, &addResp); err != nil {
		beego.Warning("Unmarshal resp error ["+string(re), err)
		return "", err
	}
	if addResp.ErrorCode != 0 {
		beego.Warning("call error :", string(re))
		return "", errors.New(addResp.ErrorMsg)
	}

	return addResp.Result.FaceToken, nil
}
