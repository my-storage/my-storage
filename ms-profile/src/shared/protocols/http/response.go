package http

import netHttp "net/http"

type Response struct {
	Data       any
	Redirect   string
	Headers    netHttp.Header
	StatusCode int
}
