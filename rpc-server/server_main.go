package main

import (
	"hello-go/rpc-server/server"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("MathService", new(server.MathService))
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("jsonrpc.Serve: accept:", err.Error())
			return
		}
		go jsonrpc.ServeConn(conn)
	}
}
