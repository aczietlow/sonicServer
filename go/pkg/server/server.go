package server

import (
	"net"
)

type Server struct {
	Conn     net.Conn
	Listener net.Listener
}

func NewPomServer(addr string) Server {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	conn, _ := listener.Accept()

	return Server{
		Conn:     conn,
		Listener: listener,
	}

}
