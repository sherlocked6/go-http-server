# Go HTTP Server (from scratch)

A minimal HTTP server built from scratch using only the Go standard library and raw TCP sockets — no external libraries or frameworks.

## Features

- Built from scratch using `net` package
- Parses HTTP request lines, headers, and body
- Supports:
  - `GET` requests with query parameters
  - `POST` requests with content-length parsing
  - Static file serving from `public/` directory
  - Basic templating with `{{placeholders}}` using query params
- Simple router system with method/path matching
- Concurrency using goroutines (each connection handled independently)
- Clean and modular folder structure

## How It Works

1. Listens on a TCP port using `net.Listen`
2. Accepts and reads each connection using a goroutine
3. Parses method, path, headers, and body
4. Dispatches request to the correct handler based on path/method
5. Supports static files with basic template replacement
6. Writes a valid HTTP/1.1 response back to the client

## Example Requests

```bash
curl "localhost:8080/greet?name=Sherlocked"
curl -X POST localhost:8080/echo -d "hello there"
curl localhost:8080/index.html
```

## Requirements

- Go 1.20+
- No external dependencies
