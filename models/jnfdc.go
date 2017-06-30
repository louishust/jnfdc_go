package models

import (
	"fmt"
	_ "jnfdc/utils"
	"time"

	"github.com/astaxie/beego/logs"
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

type EcharData struct {
	X string `orm:column(x)`
	Y uint64 `orm:column(y)`
}

func JnfdcFetchToday() ([]string, []uint64) {
	var ret = []EcharData{}
	var x = []string{}
	var y = []uint64{}
	o := orm.NewOrm()
	o.Begin()
	defer o.Commit()
	sql := "select sign_type as x, max(sign_number) as y from jn_net_sign where date(created) = current_date() group by sign_type;"
	_, err := o.Raw(sql).QueryRows(&ret)
	if err != nil {
		logs.Error(err)
		return x, y
	}

	i := 0
	for i < len(ret) {
		if ret[i].X == `储藏室（地下室）` {
			x = append(x, `储藏室`)
		} else {
			x = append(x, ret[i].X)
		}
		y = append(y, ret[i].Y)
		i = i + 1
	}

	return x, y
}

func JnfdcFetchYestoday() ([]string, []uint64) {
	var ret = []EcharData{}
	var x = []string{}
	var y = []uint64{}
	o := orm.NewOrm()
	o.Begin()
	defer o.Commit()
	sql := "select sign_type as x, max(sign_number) as y from jn_net_sign where date(created) = date_sub(current_date(), interval 1 day) group by sign_type;"
	_, err := o.Raw(sql).QueryRows(&ret)
	if err != nil {
		logs.Error(err)
		return x, y
	}

	i := 0
	for i < len(ret) {
		if ret[i].X == `储藏室（地下室）` {
			x = append(x, `储藏室`)
		} else {
			x = append(x, ret[i].X)
		}
		y = append(y, ret[i].Y)
		i = i + 1
	}

	return x, y
}

type SignTrendData struct {
	SignDate   []string   `orm:column(sign_date)`
	SignType   []string   `orm:column(sign_type)`
	SignNumber [][]uint64 `orm:column(sign_number)`
}

func JnfdcFetchStat() SignTrendData {
	var ret SignTrendData

	o := orm.NewOrm()
	o.Begin()
	defer o.Commit()

	sql := "SELECT distinct(sign_date) as sign_date from jn_net_sign_static order by sign_date asc"
	_, err := o.Raw(sql).QueryRows(&ret.SignDate)
	if err != nil {
		logs.Error(err)
		return ret
	}

	sql = "SELECT distinct(sign_type) as sign_type from jn_net_sign_static order by sign_type asc"
	_, err = o.Raw(sql).QueryRows(&ret.SignType)
	if err != nil {
		logs.Error(err)
		return ret
	}

	for _, st := range ret.SignType {
		var signNumber []uint64
		sql = fmt.Sprintf("SELECT max(sign_number) as sign_number from jn_net_sign_static where sign_type = '%s' group by sign_date asc", st)
		logs.Debug(sql)
		_, err := o.Raw(sql).QueryRows(&signNumber)
		if err != nil {
			logs.Error(err)
			return ret
		}

		ret.SignNumber = append(ret.SignNumber, signNumber)
	}

	return ret
}
