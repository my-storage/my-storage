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
	protocols.Controller[RegisterProfileControllerBody]
}

func makeRegisterProfileController() protocols.Controller[RegisterProfileControllerBody] {
	return &RegisterProfileController{}
}

func (a *RegisterProfileController) Handler(request protocols.Request[RegisterProfileControllerBody]) protocols.Response {
	return protocols.Response{
		Data:       request.Body,
		StatusCode: http.StatusOK,
	}
}
