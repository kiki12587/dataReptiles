package controllers

var (
	Mobile = `1[23456789]\d{9}`
	//检查用户端输入的是否是网址链接
	HttpCheck = `http[s]{0,1}:\/\/([\w.]+\/?)\S*`
	Http      = `<a[\s\S]+?href="(http[\s\S]+?)"`
	Email     = `[\w\.]+@\w+\.[a-z]{2,3}(\.[a-z]{2,3})?`
	//362428-1999-1014-4114
	//Card = `[1-6]\d{5}-((19\d{2})|(20\d{2}))-(((0[1-9])|(1[0123])))-((0[1-9])|([12]\d)|(3[01]))-\d{3}[\dX]`
	Card = `[1-6]\d{5}((19\d{2})|(20\d{2}))(((0[1-9])|(1[012])))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dX]`
	//<img alt="" src="static2/img/gbsxbzxrmd.png">
	Image    = `<img.+?src="(http.+?)".*?>`
	ImageAlt = `<img.+?alt="(.+?)"`
	/*img标签中的alt*/
	reImageAlt = `alt="(.+?)"`
)
