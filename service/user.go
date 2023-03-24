package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liyuanwu2020/msgo/orm"
	"log"
	"net/url"
)

type User struct {
	Id       int64
	Username string
	Age      int
	Openid   string
	Email    string
}

func SaveUser() {
	dataSourceName := fmt.Sprintf("li:ioio@tcp(8.142.149.58:3306)/msgo?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	msDb, err := orm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	user := &User{}
	user.Username = "mszlu"
	user.Openid = "123456"
	user.Email = "test@123456.com"
	user.Age = 30
	insert, i, err := msDb.New().Insert(user)
	if err != nil {
		return
	}
	log.Println(insert, i)
}
