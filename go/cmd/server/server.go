package main

import (
	"fmt"
	"net"
	"sonicServer/go/pkg/server"
	"time"
)

func main() {
	pomServer := server.NewPomServer("localhost:9001")

	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("This is the %v loop\n", i)
		// @TODO Going to need channels for this, to push message to the go routines.
		go processConnection(pomServer.Conn, msg)
		time.Sleep(time.Second * 3)
	}
}

func processConnection(conn net.Conn, msg string) {
	defer conn.Close()
	// For every connection to the listening socket, create a "while" loop
	for {
		emit(conn, msg)
	}
}

func emit(conn net.Conn, message string) {
	fmt.Fprint(conn, message)
}
