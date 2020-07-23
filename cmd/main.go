package main

import (
	"github.com/astaxie/beego"
	_ "stu_sdk/cmd/routers"
)

/**
 * @Author: WuNaiChi
 * @Date: 2020/7/1 17:06
 * @Desc:
 */

var cfgPath = "cmd/config/beego.conf"

func init() {
	if err := beego.LoadAppConfig("ini", cfgPath); err != nil {
		panic(err)
	}
}

func main() {
	beego.Run()
}
