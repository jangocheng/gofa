package resources

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var Token = ""

func init() {
	initLog()
	if err := getAuth(); err != nil {
		panic("get baidu auth error")
	}
}

func initLog() {

	beego.SetLogger("file",
		`{"filename":"logs/gofa`+time.Now().Format("20060102")+`.log"}`)

	// 打印文件名 行
	beego.SetLogFuncCall(true)

	// 异步打印
	logs.Async(1e3)

}

// 获取 Token
func getAuth() error {
	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	v.Add("client_id", beego.AppConfig.String("baiduKey::APIKey"))
	v.Add("client_secret", beego.AppConfig.String("baiduKey::SecretKey"))

	resp, err := http.PostForm(beego.AppConfig.String("baiduKey::TokenUrl"), v)
	if err != nil {
		beego.Warning()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("get baidu token error, %+v" + err.Error())
	}

	var session struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &session); err != nil {
		panic("get baidu token error, %+v" + err.Error())
	}
	Token = session.AccessToken

	return nil
}
