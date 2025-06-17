package main

import (
	"go-http-server/server"
)

func main() {
	s := server.NewServer("localhost:8080")
	s.Start()
}