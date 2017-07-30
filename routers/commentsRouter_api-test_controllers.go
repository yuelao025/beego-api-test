package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["api-test/controllers:UserController"] = append(beego.GlobalControllerRouter["api-test/controllers:UserController"],
		beego.ControllerComments{
			"Show",
			`/show`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["api-test/controllers:UserController"] = append(beego.GlobalControllerRouter["api-test/controllers:UserController"],
		beego.ControllerComments{
			"Show1",
			`/show1`,
			[]string{"get"},
			nil})

}
