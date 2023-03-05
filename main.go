package main

import (
	"fmt"
	"github.com/liyuanwu2020/msgo"
	"log"
)

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
	g := engine.Group("user")
	g.Use(func(handlerFunc msgo.HandlerFunc) msgo.HandlerFunc {
		return func(ctx *msgo.Context) {
			log.Println("test preHandler")
			handlerFunc(ctx)
			log.Println("test PostHandler")
		}
	})

	g.Post("/account/:id/edit", func(ctx *msgo.Context) {
		log.Println("exec main start")
		_, err := fmt.Fprintf(ctx.W, "%s 欢迎来到马神之路goweb教程", "lyw.com")
		if err != nil {
			return
		}
		log.Println("exec main end")
	}, func(handlerFunc msgo.HandlerFunc) msgo.HandlerFunc {
		return func(ctx *msgo.Context) {
			log.Println("test router middle")
			handlerFunc(ctx)
		}
	})
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
