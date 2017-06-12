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
	http.DefaultClient.Timeout = 10 * time.Second
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

func fetchJnfdcRegion() error {
	http.DefaultClient.Timeout = 10 * time.Second
	doc, err := goquery.NewDocument("http://www.jnfdc.gov.cn/saletoday/index.shtml")
	if err != nil {
		logs.Error(err)
		return err
	}

	o := orm.NewOrm()
	// Find the review items
	project_table := doc.Find(".project_table")
	last_idx := len(project_table.Nodes) - 1
	doc.Find(".project_table").Each(func(i int, s *goquery.Selection) {
		if i == last_idx {
			s.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
				if i != 0 {
					var data models.JNNetSignRegion
					s.Find("td").Each(func(i int, s *goquery.Selection) {
						switch i {
						case 0:
							data.RegionId, _ = strconv.ParseUint(s.Text(), 10, 64)
						case 1:
							data.RegionName = s.Text()
						case 2:
							data.OnsaleNum, _ = strconv.ParseUint(s.Text(), 10, 64)
						case 3:
							data.HouseOnsaleNum, _ = strconv.ParseUint(s.Text(), 10, 64)
						case 4:
							data.SignNum, _ = strconv.ParseUint(s.Text(), 10, 64)
						case 5:
							data.SignArea, _ = strconv.ParseFloat(s.Text(), 32)
						}
					})
					_, err := o.Insert(&data)
					if err != nil {
						logs.Error(err)
					}
				}
			})
		}
	})

	return nil
}

func CrawlJnfdc() {
	for {
		fetchJnfdc()
		fetchJnfdcRegion()
		time.Sleep(1 * time.Minute)
	}
}
