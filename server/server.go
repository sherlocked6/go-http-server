package server

import (
	"bufio"
	"fmt"
	"net"
)

type HandlerFunc func(net.Conn, *Request)

type Server struct {
	addr string
	routes map[string]map[string]HandlerFunc
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (s *Server) AddRoute(method, path string, handler HandlerFunc) {
	if s.routes[method] == nil {
		s.routes[method] = make(map[string]HandlerFunc)
	}
	s.routes[method][path] = handler
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
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, err := ReadRequest(reader)
	if err != nil {
		WriteResponse(conn, "400 Bad Request", "invalid request")
		return
	}

	fmt.Printf("Recieved request : %s %s\n", req.Method, req.Path)

	if handlers, ok := s.routes[req.Method]; ok {
		if handler, ok := handlers[req.Path]; ok {
			handler(conn, req)
		}
	}

	WriteResponse(conn, "404 Not Found", "Route not found")
}