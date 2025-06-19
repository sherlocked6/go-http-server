package handlers

import (
	"go-http-server/server"
	"net"
)

func GreetHandler(conn net.Conn, req *server.Request) {
	server.ServeStatic(conn, req.Path, req.Query)
}