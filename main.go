package main

import (
	"errors"
	"github.com/liyuanwu2020/msgo"
	mslog "github.com/liyuanwu2020/msgo/log"
	"html/template"
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

	engine := msgo.Default()
	funcMap := template.FuncMap{"showTime": ShowTime}
	engine.SetFuncMap(funcMap)
	engine.LoadTemplate("tpl/*.html")
	g := engine.Group("user")
	//先进后出
	//g.Use(msgo.Logging, msgo.Recovery)
	engine.Logger.Formatter = mslog.JsonFormat
	//engine.Logger.SetLogPath("./logs")
	engine.RegisterErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *BlogResponse:
			return http.StatusOK, e.Response()
		default:
			return http.StatusInternalServerError, "Internal Server Error"
		}
	})
	g.Any("/home", func(ctx *msgo.Context) {
		var err error
		//业务主体
		user, err := login()
		ctx.HandleWithError(http.StatusOK, user, err)

	}, func(handlerFunc msgo.HandlerFunc) msgo.HandlerFunc {
		return func(ctx *msgo.Context) {
			handlerFunc(ctx)
			//log.Println("方法级别 MiddleHandler")
		}
	})

	engine.Logger.Info("engine start")
	engine.Run()
}

type BlogResponse struct {
	Success bool
	Code    int
	Data    any
	Msg     string
}

type BlogNoDataError struct {
	Success bool
	Code    int
	Msg     string
}

func (r *BlogResponse) Error() string {
	return r.Msg
}

func (r *BlogResponse) Response() any {
	if r.Data == nil {
		return &BlogNoDataError{
			Success: r.Success,
			Code:    r.Code,
			Msg:     r.Msg,
		}
	}
	return r
}

func login() (*BlogResponse, error) {
	return &BlogResponse{
		Success: false,
		Code:    1003,
		Data:    nil,
		Msg:     "ok",
	}, errors.New("账号密码错误")
}
