package controllers

import (
	"regexp"
	"strings"
)

func (c *MainController) Card(html string) {
	compile := regexp.MustCompile(Card)
	submatch := compile.FindAllStringSubmatch(html, -1)
	for _, v := range submatch {
		t := strings.TrimSpace(v[0])
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
