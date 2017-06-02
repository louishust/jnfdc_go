package main

import (
	_ "jnfdc/routers"
	"jnfdc/tasks"
	"jnfdc/utils"

	"github.com/astaxie/beego"
)

func main() {
	go tasks.CrawlJnfdc()
	go tasks.CrawlQDfdc()

	utils.InitMenuBar()
	beego.Run()
}
