package main

import (
	_ "jnfdc/routers"
	"jnfdc/tasks"
	_ "jnfdc/utils"

	"github.com/astaxie/beego"
)

func main() {
	go tasks.CrawlJnfdc()
	beego.Run("127.0.0.1:8888")
}
