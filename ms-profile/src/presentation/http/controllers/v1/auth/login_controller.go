package auth

import (
	"net/http"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

type LoginControllerBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// type LoginControllerQuery struct {
// 	Ids string `form:"ids" binding:"required"`
// }

type LoginController struct {
	protocols.Controller[LoginControllerBody, any]
}

func makeLoginController() protocols.Controller[LoginControllerBody, any] {
	return &LoginController{}
}

func (a *LoginController) Handler(request protocols.Request[LoginControllerBody, any]) protocols.Response {
	return protocols.Response{
		Data:       request.Body,
		StatusCode: http.StatusOK,
	}
}
