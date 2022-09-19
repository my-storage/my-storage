package http

type Controller[Body, Query any] interface {
	Handler(request Request[Body, Query]) Response
}
