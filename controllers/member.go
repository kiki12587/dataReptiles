package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MemberController struct {
	beego.Controller
}

func (c *MemberController) Get() {

	c.TplName = "member.tpl"
}


func (c *MemberController) AddMember() {
	c.Ctx.WriteString("增加会员")
}



func (c *MemberController) EditMember() {

	id,err:= c.GetInt("id")
	if err != nil {
		fmt.Println(err)
		c.Ctx.WriteString("传入参数错误")
	}
fmt.Printf("%v\n,%s\n",id,id)
	c.Ctx.WriteString("编辑会员")
}
