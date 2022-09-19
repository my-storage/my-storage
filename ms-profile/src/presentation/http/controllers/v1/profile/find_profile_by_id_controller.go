package profile

import (
	"net/http"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

type FindProfileControllerBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type FindProfileControllerById struct {
	protocols.Controller[FindProfileControllerBody, any]
}

func makeFindProfileControllerById() protocols.Controller[FindProfileControllerBody, any] {
	return &FindProfileControllerById{}
}

func (a *FindProfileControllerById) Handler(request protocols.Request[FindProfileControllerBody, any]) protocols.Response {
	return protocols.Response{
		Data:       request.Body,
		StatusCode: http.StatusOK,
	}
}
