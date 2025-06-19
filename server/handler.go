package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"
)


func ReadRequest(reader *bufio.Reader) (*Request, error) {	
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error handling request line: ", err)
	}

	items := strings.Split(strings.TrimSpace(line), " ")
	if len(items) < 2 {
		return nil, fmt.Errorf("invalid request")
	}
	method, rawPath := items[0], items[1]

	headers := make(map[string]string)
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

	parsedUrl, _ := url.Parse(rawPath)

	req := &Request{
		Method: method,
		Path: parsedUrl.Path,
		Query: parsedUrl.Query(),
		Headers: headers,
	}

	if method == "POST" {
		cl := headers["Content-Length"]
		length, _ := strconv.Atoi(cl)

		buf := make([]byte, length)
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			return nil, err
		}
		req.Body = buf
	}

	return req, nil
}


func WriteResponse(conn net.Conn, status, body string) {
	resp := fmt.Sprintf("HTTP/1.1 %s \r\n" +
			"Content-Length: %d\r\n" +
			"Connection: close\r\n\r\n%s",
		status, len(body), body)
	
	conn.Write([]byte(resp))
}