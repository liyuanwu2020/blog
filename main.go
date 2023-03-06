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

func main() {
	log.Println("main start")

	//http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	_, err := fmt.Fprintf(w, "%s 欢迎来到马神之路goweb教程", "lyw.com")
	//	log.Println("HandleFunc ", err)
	//	if err != nil {
	//		return
	//	}
	//})
	//err := http.ListenAndServe(":8088", nil)
	//if err != nil {
	//	log.Fatal("启动失败", err)
	//}

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

	g.Get("/home", func(ctx *msgo.Context) {
		var err error
		query, _ := ctx.GetMapQuery("user")
		//err = ctx.HTML(http.StatusOK, "<h2>HTML</h2>")
		//err = ctx.Template("index.html", &tplData{Title: "个人中心"})
		err = ctx.JSON(http.StatusOK, query)
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

	//g.Post("/account/:id/edit", func(ctx *msgo.Context) {
	//	log.Println("exec main start")
	//	_, err := fmt.Fprintf(ctx.W, "%s 欢迎来到马神之路goweb教程", "lyw.com")
	//	if err != nil {
	//		return
	//	}
	//	log.Println("exec main end")
	//}, func(handlerFunc msgo.HandlerFunc) msgo.HandlerFunc {
	//	return func(ctx *msgo.Context) {
	//		log.Println("test router middle")
	//		handlerFunc(ctx)
	//	}
	//})
	//g.Get("/index", func(ctx *msgo.Context) {
	//	err := ctx.HTMLTemplate("index.html", &tplData{Title: "首页"}, "tpl/index.html", "tpl/header.html")
	//	log.Println(err)
	//})
	//g.Get("/login", func(ctx *msgo.Context) {
	//	err := ctx.HTMLTemplateGlob("login.html", &tplData{Title: "登录页"}, "tpl/*.html")
	//	log.Println(err)
	//})
	//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
	//git remote add origin https://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@github.com/liyuanwu2020/msgo.git
	//g.Post("/hello/:id/edit", func(ctx *msgo.Context) {
	//	_, err := fmt.Fprintf(ctx.W, "%s Post 欢迎来到马神之路goweb教程", "lyw.com")
	//	log.Println("HandleFunc ", err)
	//	if err != nil {
	//		return
	//	}
	//})
	//g.Get("/hello", func(ctx *msgo.Context) {
	//	_, err := fmt.Fprintf(ctx.W, "%s get 欢迎来到马神之路goweb教程", "lyw.com")
	//	log.Println("HandleFunc ", err)
	//	if err != nil {
	//		return
	//	}
	//})

	engine.Run()
}
