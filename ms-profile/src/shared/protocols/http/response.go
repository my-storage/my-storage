package http

import netHttp "net/http"

type Response struct {
	Data       any
	Header     netHttp.Header
	StatusCode int
}
