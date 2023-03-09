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
	Name    string   `json:"name" xml:"name" required:"true"`
	Age     int      `json:"age" xml:"age" validate:"required,max=50,min=18"`
	Address []string `json:"address" xml:"address"`
	Class   string   `json:"class" xml:"class"`
}

func main() {
	//1.创建引擎
	//1.1 创建上下文.参数处理
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
		user := &User{}
		//ctx.StructValidator = validator.Validator
		//err = ctx.BindXML(&user)
		//log.Println(user)
		//if err != nil {
		//	log.Println("处理XML错误", err)
		//}

		//user := make([]*User, 0)
		//ctx.IsValidate = true
		//ctx.StructValidator = validator.Validator
		//err = ctx.BindJson(&user)
		//if err != nil {
		//	log.Println("处理JSON错误", err)
		//}
		query, _ := ctx.GetAllPost()
		log.Println(query)
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
		//err = ctx.XML(http.StatusOK, &user)
		//ctx.File("tpl/2023课程表.xlsx")
		//ctx.FileAttachment("tpl/2023课程表.xlsx", "myCourse.xlsx")
		//ctx.FileFromFS("2023课程表.xlsx", http.Dir("tpl"))
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
	//先进后出
	g.Use(msgo.Logging)

	log.Println("engine start")
	engine.Run()
}
