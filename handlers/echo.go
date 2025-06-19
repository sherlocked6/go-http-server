package handlers

import (
	"go-http-server/server"
	"net"
)

func EchoHandler(conn net.Conn, req *server.Request) {
	message := string(req.Body)
	server.WriteResponse(conn, "200 OK", message)
}