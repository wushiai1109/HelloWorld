package main

import (
	"fmt"
	"net"
)

/**
 * @Author: Bruce
 * @Date: 2021/3/14 4:15 下午
 * @Description:
 */
type Server struct {
	Ip   string
	Port int
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

//启动服务器的接口
func (t *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	//accept

	//close listen socket

	//启动监听Message的goroutine
}
