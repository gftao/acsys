package main

import (
	_ "acsys/routers"
	_ "acsys/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
