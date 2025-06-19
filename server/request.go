package server

import "net/url"


type Request struct {
	Method string
	Path string
	Query url.Values
	Headers map[string]string
	Body []byte
}