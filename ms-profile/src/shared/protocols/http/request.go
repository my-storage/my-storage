package http

type Request[Body any] struct {
	Body      *Body
	Query     func(key string) string
	Param     func(key string) string
	Url       string
	IpAddress string
}
