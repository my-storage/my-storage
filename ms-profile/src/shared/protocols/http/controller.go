package http

type Controller[Body any] interface {
	Handler(request Request[Body]) Response
}
