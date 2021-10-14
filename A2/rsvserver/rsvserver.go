package main

import (
	"net"
	"net/rpc"

	"net4005/A2/reservation"
)

func main() {

	// create a `*Flight` object
	flight := reservation.NewFlight()

	// create a custom RPC server
	server := rpc.NewServer()

	// register `flight` object with `server`
	server.Register(flight)

	listener, _ := net.Listen("tcp", "127.0.0.1:5999")

	for {

		// get connection from the listener when client connects
		conn, _ := listener.Accept() // Accept blocks until next connection is received

		// serve connection in a separate goroutine
		go server.ServeConn(conn)
	}

}
