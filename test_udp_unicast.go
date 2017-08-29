package main

import (
	"fmt"
	"net"
	"time"
)

func receive() {
	ip, _ := net.ResolveUDPAddr("udp", "127.0.0.1:1234")
	conn, e := net.ListenUDP("udp", ip)
	fmt.Println("listener", e)
	content := make([]byte, 100)
	count, e := conn.Read(content)
	fmt.Println("接收:", count, e, string(content[:count]))
}

func send() {
	ip, _ := net.ResolveUDPAddr("udp", "127.0.0.1:1234")
	conn, e := net.DialUDP("udp", nil, ip)
	fmt.Println("listener", e)
	count, e := conn.Write([]byte("zzp"))
	conn.Close()
	fmt.Println("发送:", count, e)
}

func main() {
	go receive()
	send()
	time.Sleep(time.Second * 1)
}
