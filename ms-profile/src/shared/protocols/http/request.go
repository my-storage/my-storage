package http

type Request[Body, Query any] struct {
	Body      *Body
	Query     *Query
	Param     func(key string) string
	Url       string
	IpAddress string
}
