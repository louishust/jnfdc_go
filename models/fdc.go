package models

import (
	_ "jnfdc/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	JNNetSignTableName string = "jn_net_sign"
	TBQDNetSignFirst   string = "qd_net_sign_first_hand"
	TBQDNetSignSecond  string = "qd_net_sign_second_hand"
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

type QDNetSignFirst struct {
	Id      uint64    `orm:"column(id)"`
	Region  string    `orm:"column(sign_region)"`
	Type    string    `orm:"column(sign_type)"`
	Num     uint64    `orm:"column(sign_number)"`
	Area    float64   `orm:"column(sign_area);digits(20);decimals(2)"`
	Created time.Time `orm:"column(created);type(timestamp)"`
}

func (c *QDNetSignFirst) TableName() string {
	return TBQDNetSignFirst
}

type QDNetSignSecond struct {
	Id      uint64    `orm:"column(id)"`
	Region  string    `orm:"column(sign_region)"`
	Type    string    `orm:"column(sign_type)"`
	Num     uint64    `orm:"column(sign_number)"`
	Area    float64   `orm:"column(sign_area);digits(20);decimals(2)"`
	Created time.Time `orm:"column(created);type(timestamp)"`
}

func (c *QDNetSignSecond) TableName() string {
	return TBQDNetSignSecond
}

func init() {
	orm.RegisterModel(new(JNNetSign))
	orm.RegisterModel(new(QDNetSignFirst))
	orm.RegisterModel(new(QDNetSignSecond))
}
