package models

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url, reader string) ([]byte, error) {

	cli := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(reader))
	if err != nil {
		beego.Warning("get ", url, "error: ", err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := cli.Do(req)
	if err != nil {
		beego.Warning("get ", url, "error: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err

}
