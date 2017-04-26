package tasks

import (
	"jnfdc/models"
	_ "jnfdc/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func fetchJnfdc() error {
	http.DefaultClient.Timeout = 30 * time.Second
	doc, err := goquery.NewDocument("http://www.jnfdc.gov.cn")
	if err != nil {
		logs.Error(err)
		return err
	}

	o := orm.NewOrm()
	// Find the review items
	doc.Find("#todayview").Find(".col_bg").Find("ul").Each(func(i int, s *goquery.Selection) {
		s.Find("ul").Each(func(i int, s *goquery.Selection) {
			var ns models.JNNetSign
			s.Find("li").Each(func(i int, s *goquery.Selection) {
				switch i {
				case 0:
					ns.Type = s.Text()
				case 1:
					ns.Num, _ = strconv.ParseUint(s.Text(), 10, 64)
				case 2:
					ns.Area, _ = strconv.ParseFloat(s.Text(), 32)
				}
			})
			_, err := o.Insert(&ns)
			if err != nil {
				logs.Error(err)
			}
		})
	})

	return nil
}

func CrawlJnfdc() {
	for {
		fetchJnfdc()
		time.Sleep(5 * time.Minute)
	}
}
