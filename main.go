package main

import (
	"github.com/liyuanwu2020/msgo"
	"html/template"
	"log"
	"net/http"
	"time"
)

func ShowTime() string {
	return time.Now().String()
}

type User struct {
	Name    string   `json:"name" required:"true"`
	Age     int      `json:"age"`
	Address []string `json:"address"`
	Class   string   `json:"class"`
}

func main() {
	//1.创建引擎 (创建上下文)
	//2.添加模板函数 && 解析模板
	//3.使用引擎创建组
	//4.使用组创建路由
	//5.启动引擎

	engine := msgo.New()
	funcMap := template.FuncMap{"showTime": ShowTime}
	engine.SetFuncMap(funcMap)
	engine.LoadTemplate("tpl/*.html")
	g := engine.Group("user")

	g.Get("/toRedirect", func(ctx *msgo.Context) {
		log.Println("重定向开始")
		err := ctx.Redirect(http.StatusFound, "/user/home")
		if err != nil {
			log.Println("重定向error", err)
		}
		log.Println("重定向不执行?")
	})

	g.Any("/home", func(ctx *msgo.Context) {
		var err error
		user := make([]*User, 0)
		ctx.IsValidate = true
		err = ctx.DealJson(&user)
		if err != nil {
			log.Println("处理JSON错误", err)
		}
		//query, _ := ctx.GetAllPost()
		//log.Println(query)
		//avatar, fErr := ctx.FormFile("avatar")
		//if fErr != nil {
		//	log.Println("获取上传文件错误", err)
		//}
		//err = ctx.SaveUploadedFile(avatar, "./upload/"+avatar.Filename)
		//if err != nil {
		//	log.Println("上传文件错误", err)
		//}

		//err = ctx.HTML(http.StatusOK, "<h2>HTML</h2>")
		//err = ctx.Template("index.html", &tplData{Title: "个人中心"})
		err = ctx.JSON(http.StatusOK, user)
		//err = ctx.XML(http.StatusOK, &tplData{Title: "个人中心", Age: 20})
		//ctx.File("tpl/2023课程表.xlsx")
		//ctx.FileAttachment("tpl/2023课程表.xlsx", "myCourse.xlsx")
		//ctx.FileFromFS("2023课程表.xlsx", http.Dir("tpl"))
		//err = ctx.String(http.StatusOK, "%s的%d课程表.xlsx", "liyuanwu", 2023)
		//err = ctx.String(http.StatusOK, "%s 是由 %s 制作", "goweb框架", "go微服务框架")
		if err != nil {
			log.Println(err)
		}
	}, func(handlerFunc msgo.HandlerFunc) msgo.HandlerFunc {
		return func(ctx *msgo.Context) {
			handlerFunc(ctx)
			//log.Println("方法级别 MiddleHandler")
		}
	})

	g.Use(func(handlerFunc msgo.HandlerFunc) msgo.HandlerFunc {
		return func(ctx *msgo.Context) {
			//log.Println("组级别 PreMiddleHandler")
			handlerFunc(ctx)
			//log.Println("组级别 PostMiddleHandler")
		}
	})
	log.Println("engine start")
	engine.Run()
}
