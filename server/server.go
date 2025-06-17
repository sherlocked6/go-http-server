package server

import (
	"fmt"
	"net"
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

	method, path := ReadRequest(conn)

	fmt.Printf("Recievd request : %s %s\n", method, path)

	body := "hello from my server!"
	resp := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n" +
			"Content-Length: %d\r\n" +
			"Connection: close\r\n\r\n%s",
		len(body), body)
	
	//fmt.Println(resp)
	conn.Write([]byte(resp))
}