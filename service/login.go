package service

import (
	"errors"
	"github.com/liyuanwu2020/msgo/mspool"
	"log"
	"sync"
	"time"
)

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

func Login() (*BlogResponse, error) {

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
