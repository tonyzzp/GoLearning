package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	ip, _ := net.ResolveTCPAddr("tcp", "14.215.177.38:80")
	conn, e := net.DialTCP("tcp", nil, ip)
	fmt.Println(conn, e)
	fmt.Println(conn.RemoteAddr())
	fmt.Println(conn.LocalAddr())

	req := "GET / HTTP/1.1\r\n\r\n"
	conn.Write([]byte(req))
	conn.CloseWrite()

	content, e := ioutil.ReadAll(conn)
	fmt.Println(e)
	fmt.Println(string(content))
	conn.Close()
}
