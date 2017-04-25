package utils

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	connstr := MySQLUsername + ":" + MySQLPassword + "@tcp(" + MySQLHost + ":" + MySQLPort + ")/" + MySQLDatabase

	orm.DefaultTimeLoc = time.Local
	err := orm.RegisterDataBase("default", "mysql", connstr, MySQLMaxIdleConns)
	if err != nil {
		panic(err.Error())
	}

	/** ignore limit 1000 limitation*/
	orm.DefaultRowsLimit = -1

	logs.Info("Register MySQL completely!")
}
