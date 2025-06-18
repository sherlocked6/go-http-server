package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)


func ReadRequest(conn net.Conn) (method, path string) {
	reader := bufio.NewReader(conn)
	
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error handling request line: ", err)
	}

	items := strings.Split(strings.TrimSpace(line), " ")

	if len(items) < 2 {
		return "", ""
	}

	return items[0], items[1]
}


func writeResponse(conn net.Conn, status, body string) {
	resp := fmt.Sprintf("HTTP/1.1 %s \r\n" +
			"Content-Length: %d\r\n" +
			"Connection: close\r\n\r\n%s",
		status, len(body), body)
	
	conn.Write([]byte(resp))
}