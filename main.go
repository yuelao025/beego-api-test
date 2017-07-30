package main

import (
	_ "api-test/docs"
	_ "api-test/routers"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
    // _ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// 打印sql语句；成功的
	// orm.Debug = true
	beego.Run()
}
