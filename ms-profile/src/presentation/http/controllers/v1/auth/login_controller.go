package auth

import (
	"net/http"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

type LoginControllerBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginController struct {
	protocols.Controller[LoginControllerBody]
}

func makeLoginController() protocols.Controller[LoginControllerBody] {
	return &LoginController{}
}

func (a *LoginController) Handler(request protocols.Request[LoginControllerBody]) protocols.Response {
	return protocols.Response{
		Data:       request.Body,
		StatusCode: http.StatusOK,
	}
}
