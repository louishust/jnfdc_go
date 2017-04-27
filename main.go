package main

import (
	_ "jnfdc/routers"
	"jnfdc/tasks"
	"jnfdc/utils"

	"github.com/astaxie/beego"
)

func main() {
	go tasks.CrawlJnfdc()

	utils.InitMenuBar()
	beego.Run()
}
