package main

import (
	"go-http-server/handlers"
	"go-http-server/server"
)

func main() {
	s := server.NewServer("localhost:8080")

	s.AddRoute("GET", "/", handlers.GreetHandler)
	s.AddRoute("GET", "/greet", handlers.GreetHandler)
	s.AddRoute("POST", "/echo", handlers.EchoHandler)

	s.Start()
}