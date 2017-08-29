package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/rpc"
)

type Req struct {
	UserName string
	Password string
	Path     string
}

type Resp struct {
	Success bool
	Msg     string
}

type PayCenter int

func (p *PayCenter) Buy(req *Req, resp *Resp) error {
	fmt.Println("收到请求", req)
	if req.Password == "" {
		return errors.New("请输入密码")
	}
	resp.Success = true
	resp.Msg = "购买" + req.Path + "成功"
	return nil
}

func main() {
	pc := new(PayCenter)
	rpc.Register(pc)
	rpc.HandleHTTP()

	http.ListenAndServe("127.0.0.1:1234", nil)
}
