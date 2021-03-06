package controllers

import (
	"regexp"
	"strings"
)

//超链接爬取
func (c *MainController) Http(html string) {
	compile := regexp.MustCompile(Http)
	submatch := compile.FindAllStringSubmatch(html, -1)
	for _, v := range submatch {
		t := strings.TrimSpace(v[1])
		if t == "" {
			continue
		}
		ret = append(ret, t)
	}

	boolStatus := CheckSlice(ret)
	if boolStatus == true {
		res = &UitlStruct{-1, []string{"暂无数据匹配"}}
	} else {
		res = &UitlStruct{0, ret}
	}
	c.Data["json"] = res
	c.ServeJSON()
}
