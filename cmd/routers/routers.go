package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"stu_sdk/studentManage"
)

/**
 * @Author: WuNaiChi
 * @Date: 2020/6/22 15:57
 * @Desc:
 */

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "x-requested-with", "no-referrer-when-downgrade"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowOrigins:     beego.AppConfig.Strings("Allowip"),
	}))

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/students",
			beego.NSRouter("/createStu", &studentManage.StuController{}, "post:CommitStudentInfo"),
			beego.NSRouter("/updateStu", &studentManage.StuController{}, "post:UpdateStudentInfo"),
			beego.NSRouter("/", &studentManage.StuController{}, "post:QueryStudentInfoList"),
			beego.NSRouter("/:AcctId", &studentManage.StuController{}, "get:QueryStudentInfo"),
			beego.NSRouter("/deleteStu", &studentManage.StuController{}, "post:DeleteStudentInfo"),
		),
	)
	//  todo:AddNamespace这个为啥呢
	beego.AddNamespace(ns)
}
