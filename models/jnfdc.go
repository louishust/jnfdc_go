package models

import (
	_ "jnfdc/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	TB_JNNetSignRegion string = "jn_net_sign_region"
)

type JNNetSignRegion struct {
	Id             uint64    `orm:"column(id)"`
	RegionId       uint64    `orm:"column(region_id)"`
	RegionName     string    `orm:"column(region_name)"`
	OnsaleNum      uint64    `orm:"column(onsale_num)"`
	HouseOnsaleNum uint64    `orm:"column(house_onsale_num)"`
	SignNum        uint64    `orm:"column(sign_number)"`
	SignArea       float64   `orm:"column(sign_area);digits(20);decimals(2)"`
	Created        time.Time `orm:"column(created);type(timestamp)"`
}

func (c *JNNetSignRegion) TableName() string {
	return TB_JNNetSignRegion
}

func init() {
	orm.RegisterModel(new(JNNetSignRegion))
}
