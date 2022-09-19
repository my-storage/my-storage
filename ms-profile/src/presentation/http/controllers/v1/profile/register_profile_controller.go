package profile

import (
	"net/http"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

type RegisterProfileControllerBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterProfileController struct {
	protocols.Controller[RegisterProfileControllerBody, any]
}

func makeRegisterProfileController() protocols.Controller[RegisterProfileControllerBody, any] {
	return &RegisterProfileController{}
}

func (a *RegisterProfileController) Handler(request protocols.Request[RegisterProfileControllerBody, any]) protocols.Response {
	return protocols.Response{
		Data:       request.Body,
		StatusCode: http.StatusOK,
	}
}
