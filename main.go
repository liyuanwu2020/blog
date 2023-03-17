package main

import (
	"errors"
	"github.com/liyuanwu2020/msgo"
	"github.com/liyuanwu2020/msgo/mslog"
	"github.com/liyuanwu2020/msgo/mspool"
	"github.com/liyuanwu2020/msgo/token"
	"html/template"
	"log"
	"net/http"
	"sync"
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
	auth := &msgo.Accounts{
		Users: make(map[string]string),
	}
	auth.Users["mszlu"] = "123456"
	//engine.Use(auth.BasicAuth)
	g := engine.Group("user")
	//先进后出
	//g.Use(msgo.Logging, msgo.Recovery)
	engine.Logger.Formatter = mslog.JsonFormat
	engine.RegisterErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *BlogResponse:
			return http.StatusOK, e.Response()
		default:
			return http.StatusInternalServerError, "Internal Server Error"
		}
	})
	data := struct {
		Username string
		Passwd   string
	}{
		"mszlu",
		"123456",
	}
	engine.Logger.Info(msgo.BasicAuth(data.Username, data.Passwd))
	g.Any("/login", func(ctx *msgo.Context) {

		jwt := &token.JwtHandler{}
		jwt.Key = []byte("abc")
		jwt.RefreshKey = "token"
		jwt.SendCookie = true
		jwt.Timeout = 10 * time.Minute
		jwt.Authenticator = func(ctx *msgo.Context) (map[string]any, error) {
			d := make(map[string]any)
			d["userId"] = 100
			return d, nil
		}
		token, err := jwt.LoginHandler(ctx)

		if err != nil {
			ctx.Logger.Error(err)
		} else {
			ctx.Set(jwt.RefreshKey, token.RefreshToken)
			handler, _ := jwt.RefreshTokenHandler(ctx)
			ctx.Logger.Info(handler)
			ctx.JSON(http.StatusOK, token)
		}
	})
	g.Any("/home", func(ctx *msgo.Context) {
		var err error
		//业务主体
		//user, err := login()
		user := &User{
			Name: "xiaoli",
			Age:  20,
		}
		ctx.HandleWithError(http.StatusOK, user, err)

	}, func(handlerFunc msgo.HandlerFunc) msgo.HandlerFunc {
		return func(ctx *msgo.Context) {
			handlerFunc(ctx)
			//mslog.Println("方法级别 MiddleHandler")
		}
	})

	engine.Logger.Info("engine start")
	//engine.Run()
	engine.RunTLS(":8088", "key/server.pem", "key/server.key")
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

	pool, _ := mspool.NewPool(3)
	t := time.Now().UnixMilli()
	var wg sync.WaitGroup
	wg.Add(4)
	pool.Submit(func() {
		time.Sleep(time.Second * 2)
		log.Println("1")
		wg.Done()
	})
	pool.Submit(func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		log.Println("2")
		panic("oh no")
	})
	pool.Submit(func() {
		time.Sleep(time.Second * 2)
		log.Println("3")
		wg.Done()
	})
	pool.Submit(func() {
		time.Sleep(time.Second * 2)
		log.Println("4")
		wg.Done()
	})
	wg.Wait()
	log.Printf("time: %v", time.Now().UnixMilli()-t)

	return &BlogResponse{
		Success: false,
		Code:    1003,
		Data:    nil,
		Msg:     "ok",
	}, errors.New("账号密码错误")
}
