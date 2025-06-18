package server

import (
	"fmt"
	"io"
	"mime"
	"net"
	"net/url"
	"os"
	"path/filepath"
)

func ServeStatic(conn net.Conn, requestPath string, query url.Values) {
	if requestPath == "/" {
		requestPath = "/index.html"
	}

	cleanPath := filepath.Clean(requestPath)
	filePath := filepath.Join("./public", cleanPath)

	//fmt.Println(filePath)

	//file, err := os.Open("./public/index.html")
	file, err := os.Open(filePath)

	if err != nil {
		writeResponse(conn, "404 Not Found", "File not found")
		return
	}
	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		writeResponse(conn, "404 Not Found", "File not found")
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	header := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: %s\r\nonnection: close\r\n\r\n", info.Size(), contentType)

	conn.Write([]byte(header))

	io.Copy(conn, file)
}