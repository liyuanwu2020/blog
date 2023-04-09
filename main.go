package main

import (
	"github.com/liyuanwu2020/blog/service"
	"github.com/liyuanwu2020/micro.service.pb/go/user"
	"github.com/liyuanwu2020/msgo"
	"github.com/liyuanwu2020/msgo/mslog"
	"github.com/liyuanwu2020/msgo/register"
	"github.com/liyuanwu2020/msgo/rpc"
	"github.com/liyuanwu2020/msgo/token"
	"google.golang.org/grpc"
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

	nacosClient, nacosErr := register.CreateNacosClient()
	if nacosErr != nil {
		panic(nacosErr)
	}
	conf := register.NacosServiceConfig{
		Ip:          "localhost",
		Port:        9112,
		ServiceName: "user",
	}
	registerService, nacosErr := register.NacosServiceRegister(nacosClient, conf)
	if nacosErr != nil {
		panic(nacosErr)
	}
	log.Println(registerService)

	grpcServer, grpcErr := rpc.NewGrpcServer(":9112")
	grpcServer.Register(func(g *grpc.Server) {
		user.RegisterUserServiceServer(g, &service.UserService{})
	})
	grpcErr = grpcServer.Run()
	panic(grpcErr)
	return

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
		case *service.BlogResponse:
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
		t, err := jwt.LoginHandler(ctx)

		if err != nil {
			ctx.Logger.Error(err)
		} else {
			ctx.Set(jwt.RefreshKey, t.RefreshToken)
			handler, _ := jwt.RefreshTokenHandler(ctx)
			ctx.Logger.Info(handler)
			_ = ctx.JSON(http.StatusOK, t)
		}
	})
	g.Any("/home", func(ctx *msgo.Context) {
		var err error
		//业务主体
		//user, err := login()
		user := &User{
			Name: "小李",
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
