package models

import (
	_ "jnfdc/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	JNNetSignTableName string = "jn_net_sign"
)

type JNNetSign struct {
	Id      uint64    `orm:"column(id)"`
	Type    string    `orm:"column(sign_type)"`
	Num     uint64    `orm:"column(sign_number)"`
	Area    float64   `orm:"column(sign_area);digits(20);decimals(2)"`
	Created time.Time `orm:"column(created);type(timestamp)"`
}

func (c *JNNetSign) TableName() string {
	return JNNetSignTableName
}

func init() {
	orm.RegisterModel(new(JNNetSign))
}
