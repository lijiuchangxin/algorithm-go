package main

import (
	"github.com/astaxie/beego"
	_ "lzsc/models"
	_ "lzsc/routers"
)


func main() {

	beego.Run()
}

