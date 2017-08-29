package main

import (
	"fmt"
	"net/rpc"
)

type req2 struct {
	UserName string
	Password string
	Path     string
}

type resp2 struct {
	Success bool
	Msg     string
}

func main() {
	client, e := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	fmt.Println(client, e)

	req := req2{
		"aa",
		"paaa",
		"pp",
	}
	r := new(resp2)
	e = client.Call("PayCenter.Buy", req, r)
	fmt.Println(e)
	fmt.Println(r.Msg)
	client.Close()
}
