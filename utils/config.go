package utils

import (
	"path"
	"runtime"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var (
	MySQLHost         string
	MySQLPort         string
	MySQLUsername     string
	MySQLPassword     string
	MySQLDatabase     string
	MySQLMaxIdleConns int
)

func initlogs() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"jnfdc.log","level":7,"maxlines":0,"maxsize":33554432,"daily":true,"maxdays":7}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Async()
}
func init() {
	initlogs()

	_, filename, _, _ := runtime.Caller(1)
	confpath := path.Join(path.Dir(filename), "../conf/fdc.cnf")
	conf, err := config.NewConfig("ini", confpath)
	if err != nil {
		logs.Error(err)
	}

	MySQLHost = conf.String("mysql::host")
	MySQLPort = conf.String("mysql::port")
	MySQLUsername = conf.String("mysql::username")
	MySQLPassword = conf.String("mysql::password")
	MySQLDatabase = conf.String("mysql::database")
	MySQLMaxIdleConns, _ = conf.Int("mysql::max_idle")
}
