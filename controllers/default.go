package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"regexp"
	"strconv"
)

var (
	ret = make([]string, 0)
	res *UitlStruct
)

type UitlStruct struct {
	Code int
	Msg  interface{}
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Post() {
	tp, err := strconv.Atoi(c.GetString("type"))
	if err != nil {
		c.Ctx.WriteString("类型选择错误")
		return
	}
	url := c.GetString("url")
	matchString, _ := regexp.MatchString(HttpCheck, url)
	if matchString == false {
		c.Ctx.WriteString("请输入以http开头或者https开头的网址")
		return
	}

	data := httplib.Get(url)
	html, _ := data.String()
	//fmt.Println(html)
	switch tp {
	case 1:
		c.Mobile(html)
	case 2:
		c.Email(html)
	case 3:
		c.Http(html)
	case 4:
		c.Card(html)
	case 5:
		c.Image(html)
	default:
		c.Ctx.WriteString("请选择正确类型")
		return
	}
}
