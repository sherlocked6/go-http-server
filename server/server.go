package server

import (
	"fmt"
	"net"
	"net/url"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Println("Failed to bind adrress:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on server: ", s.addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error: ", err)
			continue
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()

	method, rawPath := ReadRequest(conn)

	fmt.Printf("Recieved request : %s %s\n", method, rawPath)

	var body string
	var status string

	parsedUrl, err := url.Parse(rawPath)

	if err != nil {
		writeResponse(conn, "400 Bad Request", "Malformed URL")
		return
	}

	path := parsedUrl.Path
	query := parsedUrl.Query()

	if method == "GET" {
		switch path {
		case "/":
			body = "hello from my server!"
			status = "200 OK"
		case "/health":
			body = "OK"
			status = "200 OK"
		case "/greet":
			name := query.Get("name")
			if name == ""{
				name = "stranger"
			}
			body = "hello there " + name
			status = "200 OK"
		default:
			body = "404 Not Found"
			status = "404 Not Found"
		}
	} else {
		body = "405 Method Not Allowed"
		status = "405 Method Not Allowed"
	}
	writeResponse(conn, status, body)
}