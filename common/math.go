/*
@Time : 2020/12/2 14:39
@Author : HP
@File : math
@Software: GoLand
*/
package common

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	randomMT sync.Mutex
)

//随机数文件名包含扩展名
func RandFileName(ext string) string {
	randomNum := strconv.Itoa(GetRanddomInt(100, 1000))
	name := strconv.Itoa(int(time.Now().UnixNano())) + "_" + randomNum + "." + ext
	return name
}

//生成随机数
func GetRanddomInt(start, end int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randnum := start + r.Intn(end-start)
	return randnum
}
