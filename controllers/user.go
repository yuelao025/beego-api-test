package controllers

import (
	// "api-test/models"
	// "encoding/json"
	// "fmt"


	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)


func init() {
	
 	orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:123456@/blog?charset=utf8")

}




// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get   http://localhost:8081/v1/user/show  注意这个是注解路由；看好注释的最后一行的写法！！ 成功的 否则404路由不对
// @Description get user by uid
// @Success 200 
// @Failure 403  
// @router /show [get]
func (this *UserController) Show() {


    o := orm.NewOrm()
    var maps []orm.Params
    o.Raw("SELECT * from user as u inner join comment as c on u.id = c.user_id").Values(&maps)
    // fmt.Println(maps)
	// name := "lvxx"
	// fmt.Fprint(this.Ctx.ResponseWriter, name)

	this.Data["json"] = maps
    this.ServeJSON()
}


// @Title Get   http://localhost:8081/v1/user/show  注意这个是注解路由；看好注释的最后一行的写法！！ 成功的 否则404路由不对
// @Description get user by uid
// @Success 200 
// @Failure 403  
// @router /show1 [get]
func (this *UserController) Show1() {

	//get参数
	sku := this.GetString("sku")
	// fmt.Println(sku)
	// fmt.Fprint(this.Ctx.ResponseWriter, sku)

    o := orm.NewOrm()
    var maps []orm.Params
    o.Raw("SELECT * from user as u inner join comment as c on u.id = c.user_id and u.id=?",sku).Values(&maps)
    // fmt.Println(maps)
	// name := "lvxx"
	// fmt.Fprint(this.Ctx.ResponseWriter, name)

	this.Data["json"] = maps
    this.ServeJSON()
}



// // @Title Get
// // @Description get user by uid
// // @Param	uid		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.User
// // @Failure 403 :uid is empty
// // @router /:uid [get]
// func (u *UserController) Get() {
// 	uid := u.GetString(":uid")
// 	fmt.Println(uid)
// 	name := "lvxx"
// 	fmt.Fprint(u.Ctx.ResponseWriter, name)
// 	return 
// 	// if uid != "" {
// 	// 	user, err := models.GetUser(uid)
// 	// 	if err != nil {
// 	// 		u.Data["json"] = err.Error()
// 	// 	} else {
// 	// 		u.Data["json"] = user
// 	// 	}
// 	// }
// 	// u.ServeJSON()
// }

