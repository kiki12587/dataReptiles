package routers

import (
	"demo01/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Get;post:Post")
	//beego.Router("/mobile", &controllers.MobileController{})
	//beego.Router("/member", &controllers.MemberController{})
	//beego.Router("/member/add", &controllers.MemberController{},"get:AddMember")
	//beego.Router("/member/edit", &controllers.MemberController{},"get:EditMember")
}
