package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)


func ReadRequest(reader *bufio.Reader) (method, path string, headers map[string]string) {
	
	headers = make(map[string]string)
	
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error handling request line: ", err)
	}

	items := strings.Split(strings.TrimSpace(line), " ")

	if len(items) < 2 {
		return "", "", headers
	}

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == ""{
			break
		}
		items := strings.SplitN(line, ":", 2)
		if len(items) == 2 {
			headers[strings.TrimSpace(items[0])] = strings.TrimSpace(items[1])
		}
	}

	fmt.Println(len(headers))

	return items[0], items[1], headers
}


func writeResponse(conn net.Conn, status, body string) {
	resp := fmt.Sprintf("HTTP/1.1 %s \r\n" +
			"Content-Length: %d\r\n" +
			"Connection: close\r\n\r\n%s",
		status, len(body), body)
	
	conn.Write([]byte(resp))
}