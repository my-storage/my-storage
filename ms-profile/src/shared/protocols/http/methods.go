package http

type HttpMethods string

const (
	Get     HttpMethods = "GET"
	Post    HttpMethods = "POST"
	Put     HttpMethods = "PUT"
	Patch   HttpMethods = "PATCH"
	Delete  HttpMethods = "DELETE"
	Options HttpMethods = "OPTIONS"
)
