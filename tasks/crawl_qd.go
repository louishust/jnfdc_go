package tasks

import (
	"jnfdc/models"
	"jnfdc/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func FetchQDfdcFirst() error {
	http.DefaultClient.Timeout = 30 * time.Second
	doc, err := goquery.NewDocument("http://www.qdfd.com.cn/qdweb/realweb/fh/fhNewsale.jsp")
	if err != nil {
		logs.Error(err)
		return err
	}

	o := orm.NewOrm()
	// Find the review items
	doc.Find(".ysjrcj_table").Each(func(i int, s *goquery.Selection) {
		if i >= 2 && i <= 16 {
			var ns models.QDNetSignFirst
			nodes := s.Find("td").Nodes
			region := nodes[0].FirstChild.Data
			b := []byte(region[4:])
			b2, err := utils.GbkToUtf8(b)
			if err != nil {
				logs.Debug(err)
			}
			region = string(b2)

			house_num, _ := strconv.ParseUint(nodes[1].FirstChild.Data, 10, 64)
			house_area, _ := strconv.ParseFloat(nodes[2].FirstChild.Data, 32)
			office_num, _ := strconv.ParseUint(nodes[4].FirstChild.Data, 10, 64)
			office_area, _ := strconv.ParseFloat(nodes[5].FirstChild.Data, 32)
			business_num, _ := strconv.ParseUint(nodes[7].FirstChild.Data, 10, 64)
			business_area, _ := strconv.ParseFloat(nodes[8].FirstChild.Data, 32)
			other_num, _ := strconv.ParseUint(nodes[10].FirstChild.Data, 10, 64)
			other_area_str := nodes[11].FirstChild.Data
			other_area_str = other_area_str[0 : len(other_area_str)-4]
			other_area, _ := strconv.ParseFloat(other_area_str, 32)

			ns.Region = region

			ns.Id = 0
			ns.Type = "住宅"
			ns.Num = house_num
			ns.Area = house_area
			_, err = o.Insert(&ns)
			if err != nil {
				logs.Error(err)
			}

			ns.Id = 0
			ns.Type = "办公"
			ns.Num = office_num
			ns.Area = office_area
			_, err = o.Insert(&ns)
			if err != nil {
				logs.Error(err)
			}

			ns.Id = 0
			ns.Type = "商业"
			ns.Num = business_num
			ns.Area = business_area
			_, err = o.Insert(&ns)
			if err != nil {
				logs.Error(err)
			}

			ns.Id = 0
			ns.Type = "其它"
			ns.Num = other_num
			ns.Area = other_area
			_, err = o.Insert(&ns)
			if err != nil {
				logs.Error(err)
			}
		}
	})
	return nil
}

func FetchQDfdcSecond() error {
	http.DefaultClient.Timeout = 30 * time.Second
	doc, err := goquery.NewDocument("http://www.qdfd.com.cn/qdweb/realweb/sh/shindex.jsp")
	if err != nil {
		logs.Error(err)
		return err
	}

	o := orm.NewOrm()
	// Find the review items
	doc.Find(".jrcj_table4").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(i int, s *goquery.Selection) {
			if i >= 2 && i <= 16 {
				var ns models.QDNetSignSecond
				nodes := s.Find("td").Nodes
				region := nodes[0].FirstChild.Data
				b := []byte(region)
				b2, err := utils.GbkToUtf8(b)
				if err != nil {
					logs.Debug(err)
				}
				region = string(b2)

				house_num, _ := strconv.ParseUint(nodes[1].FirstChild.Data, 10, 64)
				house_area, _ := strconv.ParseFloat(nodes[2].FirstChild.Data, 32)
				office_num, _ := strconv.ParseUint(nodes[4].FirstChild.Data, 10, 64)
				office_area, _ := strconv.ParseFloat(nodes[5].FirstChild.Data, 32)
				business_num, _ := strconv.ParseUint(nodes[7].FirstChild.Data, 10, 64)
				business_area, _ := strconv.ParseFloat(nodes[8].FirstChild.Data, 32)
				other_num, _ := strconv.ParseUint(nodes[10].FirstChild.Data, 10, 64)
				other_area_str := nodes[11].FirstChild.Data
				other_area, _ := strconv.ParseFloat(other_area_str, 32)

				ns.Region = region

				ns.Id = 0
				ns.Type = "住宅"
				ns.Num = house_num
				ns.Area = house_area
				_, err = o.Insert(&ns)
				if err != nil {
					logs.Error(err)
				}

				ns.Id = 0
				ns.Type = "办公"
				ns.Num = office_num
				ns.Area = office_area
				_, err = o.Insert(&ns)
				if err != nil {
					logs.Error(err)
				}

				ns.Id = 0
				ns.Type = "商业"
				ns.Num = business_num
				ns.Area = business_area
				_, err = o.Insert(&ns)
				if err != nil {
					logs.Error(err)
				}

				ns.Id = 0
				ns.Type = "其它"
				ns.Num = other_num
				ns.Area = other_area
				_, err = o.Insert(&ns)
				if err != nil {
					logs.Error(err)
				}
			}
		})
	})
	return nil
}

func CrawlQDfdc() {
	for {
		FetchQDfdcFirst()
		FetchQDfdcSecond()
		time.Sleep(5 * time.Minute)
	}
}
