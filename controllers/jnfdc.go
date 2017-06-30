package controllers

import (
	"fmt"
	"jnfdc/models"
	"time"

	"github.com/astaxie/beego"
)

type JnfdcController struct {
	beego.Controller
}

// @Title Get NetSign
// @Summary 获取今日昨日网签数量
// @router /netsign [get]
func (c *JnfdcController) NetSign() {
	c.Data["Title"] = "济南房产今昨网签数"

	c.Data["Dimension"] = []string{"今日网签", "昨日网签"}

	c.Data["XToday"], c.Data["YToday"] = models.JnfdcFetchToday()
	c.Data["XYestoday"], c.Data["YYestoday"] = models.JnfdcFetchYestoday()
	c.TplName = "jnfdc_netsign.tpl"
}

// @Title Get NetSign Statistic
// @Summary 获取今日昨日网签数量
// @router /netsignstat [get]
func (c *JnfdcController) NetSignStat() {
	c.Data["Title"] = "济南房产网签趋势图"

	date := models.JnfdcFetchStat()
	t := time.Now().AddDate(0, 0, -30)
	start := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())

	c.Data["X"] = date.SignDate
	c.Data["YName"] = date.SignType
	c.Data["YData"] = date.SignNumber
	c.Data["Start"] = start

	c.TplName = "jnfdc_netsignstat.tpl"
}
