/*
@Time : 2020/12/2 14:25
@Author : HP
@File : 字符转换
@Software: GoLand
*/
package common

import (
	"github.com/axgle/mahonia"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
