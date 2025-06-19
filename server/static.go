package server

import (
	"fmt"
	"io"
	"strings"
	//"mime"
	"net"
	"net/url"
	"os"
	"path/filepath"
)

func ServeStatic(conn net.Conn, requestPath string, query url.Values) {
	if requestPath == "/" {
		requestPath = "index"
	}
	cleanPath := filepath.Clean(requestPath) + ".html"
	filePath := filepath.Join("./public", cleanPath)

	// fmt.Println(filePath)
	//file, err := os.Open("./public/index.html")
	file, err := os.Open(filePath)
	if err != nil {
		WriteResponse(conn, "404 Not Found", "File not found")
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		WriteResponse(conn, "500 Internal Server Error", "Failed to read file")
		return
	}

	body := string(data)
	if len(query) == 0 {
		body = strings.ReplaceAll(body, "{{name}}", "Stranger")
	}
	for key, value := range query {
		placeholder := fmt.Sprintf("{{%s}}", key)
		body = strings.ReplaceAll(body, placeholder, value[0])
	}

	WriteResponse(conn, "200 OK", body)

	// info, err := file.Stat()
	// if err != nil {
	// 	WriteResponse(conn, "404 Not Found", "File not found")
	// 	return
	// }

	// contentType := mime.TypeByExtension(filepath.Ext(filePath))
	// if contentType == "" {
	// 	contentType = "application/octet-stream"
	// }

	// header := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: %s\r\nonnection: close\r\n\r\n", info.Size(), contentType)

	// conn.Write([]byte(header))

	// io.Copy(conn, file)
}