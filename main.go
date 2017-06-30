package main

import (
	_ "jnfdc/routers"
	"jnfdc/tasks"

	"github.com/astaxie/beego"
)

func main() {
	go tasks.CrawlJnfdc()
	go tasks.CrawlQDfdc()
	go tasks.JnfdcStat()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
