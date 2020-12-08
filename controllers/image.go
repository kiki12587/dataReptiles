package controllers

import (
	"dataReptiles/common"
	"fmt"
	"github.com/astaxie/beego"
	"regexp"
	"strings"
)

func (c *MainController) Image(html string) {
	imginfos := make([]map[string]string, 0)
	compile := regexp.MustCompile(Image)
	submatch := compile.FindAllStringSubmatch(html, -1)
	for _, v := range submatch {
		imginfo := make(map[string]string)
		imgurl := strings.TrimSpace(v[0])
		if imgurl == "" {
			continue
		}
		imgurl = common.ConvertToString(imgurl, "gbk", "utf-8")

		imginfo["url"] = v[1]
		imginfo["filename"] = GetImgNameFormTag(imgurl)

		imginfos = append(imginfos, imginfo)
	}

	boolStatus := CheckMap(imginfos)
	if boolStatus == true {
		res = &UitlStruct{-1, []string{"暂无数据匹配"}}
	} else {
		fmt.Println("图片张数:", len(imginfos))
		downloadImgOpen := beego.AppConfig.String("downloadImgOpen")
		if downloadImgOpen == "1" { //"1"代表开启图片下载
			//遍历图片
			for _, v := range imginfos {
				downloadImageAsync(v["url"], v["filename"])
			}
		}

		res = &UitlStruct{0, imginfos}
	}

	c.Data["json"] = res
	c.ServeJSON()
}
