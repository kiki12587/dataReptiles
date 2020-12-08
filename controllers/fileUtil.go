/*
@Time : 2020/12/2 10:50
@Author : HP
@File : fileUtil
@Software: GoLand
*/
package controllers

import (
	"dataReptiles/common"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	chSem = make(chan int, 5)
)

//异步下载图片
func downloadImageAsync(url, filename string) {
	go func() {
		chSem <- 123
		downloadImage(url, filename)
		<-chSem
	}()
}

//同步下载
func downloadImage(url, filename string) {
	data, _ := http.Get(url)
	defer data.Body.Close()
	bytes, _ := ioutil.ReadAll(data.Body)
	fileConfigPath := beego.AppConfig.String("filename")
	retFileName := fileConfigPath + filename
	err := ioutil.WriteFile(retFileName, bytes, 0644)
	if err != nil {
		fmt.Println(retFileName, "下载失败")
	} else {
		fmt.Println(retFileName, "下载成功")
	}
}

/**
从img 标签中 提取文件名
有alt使用alt加时间戳_随机数做文件名,没有alt用时间戳_随机数做文件名
*/
func GetImgNameFormTag(imgTag string) string {
	//尝试从img中提取alt
	re := regexp.MustCompile(reImageAlt)
	rets := re.FindAllStringSubmatch(imgTag, 1)
	if len(rets) > 0 {
		return rets[0][1] + common.RandFileName("jpg")
	} else {
		return common.RandFileName("jpg")
	}
}
