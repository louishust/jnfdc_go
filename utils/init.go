package utils

import (
	"crypto/tls"
	"errors"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

func InitMenuBar() error {
	logs.Info("Init menu bar")
	token, err := GetAccessToken()
	if err != nil {
		logs.Error(err)
		return err
	}

	menu := `{ "button":[ {"type":"click", "name":"今日网签", "key":"V1001_TODAY_SIGN" }] }`

	req := httplib.Post("https://api.weixin.qq.com/cgi-bin/menu/create")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	req.Param("access_token", token)
	req.Body([]byte(menu))
	var res WXResponse
	err = req.ToJSON(&res)
	if res.ErrCode != 0 {
		logs.Error(res.ErrMsg)
		return errors.New(res.ErrMsg)
	}

	return err
}
