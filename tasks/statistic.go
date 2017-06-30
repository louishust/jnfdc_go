package tasks

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func JnfdcStatNetSign() {
	o := orm.NewOrm()
	o.Begin()
	defer o.Commit()
	sql := "replace into jn_net_sign_static(sign_type, sign_number, sign_date)  select sign_type, max(sign_number) as sign_number, date(created) as sign_date FROM jn_net_sign where date(created) = current_date() group by sign_date, sign_type;"

	_, err := o.Raw(sql).Exec()
	if err != nil {
		logs.Error(err)
	}
}

func JnfdcStat() {
	for {
		JnfdcStatNetSign()
		time.Sleep(1 * time.Minute)
	}
}
