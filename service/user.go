package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liyuanwu2020/msgo/orm"
	"log"
	"net/url"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Openid   string `json:"openid"`
	Email    string `json:"email"`
}

func CountUser() {
	dataSourceName := fmt.Sprintf("li:ioio@tcp(8.142.149.58:3306)/msgo?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	msDb, err := orm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	msDb.New()
}

func UpdateUser() {
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
	i, err := msDb.New().Table("user").Where("id", 1).Update("age", "1")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(i)
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

func SaveUserBatch() {
	dataSourceName := fmt.Sprintf("li:ioio@tcp(8.142.149.58:3306)/msgo?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	msDb, err := orm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	user := &User{}
	user.Username = "mszlu1"
	user.Openid = "123456"
	user.Email = "test@123456.com"
	user.Age = 1
	var users []any
	user1 := &User{}
	user1.Username = "mszlu2"
	user1.Openid = "123456"
	user1.Email = "test@123456.com"
	user1.Age = 2
	user2 := &User{}
	user2.Username = "mszlu3"
	user2.Openid = "123456"
	user2.Email = "test@123456.com"
	user2.Age = 3
	users = append(users, user, user1, user2)

	insert, i, err := msDb.New().InsertBath(users)
	if err != nil {
		return
	}
	log.Println(insert, i)
}
