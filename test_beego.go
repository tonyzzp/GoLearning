package main

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	this.Ctx.WriteString("hello beego")
}

func main() {
	beego.Router("/", &HelloController{})
	beego.Run()
}
