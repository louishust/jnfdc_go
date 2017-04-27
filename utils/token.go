package utils

import (
	"crypto/tls"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

type AccessToken struct {
	Token  string `json:"access_token"`
	Expire int    `json:"expires_in"`
}

func GetAccessToken() (string, error) {
	var at AccessToken
	req := httplib.Get("https://api.weixin.qq.com/cgi-bin/token")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Param("grant_type", "client_credential")
	req.Param("appid", "wx9252ad6f802fcb20")
	req.Param("secret", "14cff4105b5b68f0b5c2de960ba5d805")

	err := req.ToJSON(&at)
	if err != nil {
		logs.Error(err)
		return "", err
	} else {
		return at.Token, nil
	}
}
