package profile

import (
	"net/http"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

type UpdateProfileControllerBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileController struct {
	protocols.Controller[UpdateProfileControllerBody, any]
}

func makeUpdateProfileController() protocols.Controller[UpdateProfileControllerBody, any] {
	return &UpdateProfileController{}
}

func (a *UpdateProfileController) Handler(request protocols.Request[UpdateProfileControllerBody, any]) protocols.Response {
	return protocols.Response{
		Data:       request.Body,
		StatusCode: http.StatusOK,
	}
}
