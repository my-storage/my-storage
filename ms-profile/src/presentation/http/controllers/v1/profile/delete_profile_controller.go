package profile

import (
	"net/http"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

type DeleteProfileControllerBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DeleteProfileController struct {
	protocols.Controller[DeleteProfileControllerBody]
}

func makeDeleteProfileController() protocols.Controller[DeleteProfileControllerBody] {
	return &DeleteProfileController{}
}

func (a *DeleteProfileController) Handler(request protocols.Request[DeleteProfileControllerBody]) protocols.Response {
	return protocols.Response{
		Data:       request.Body,
		StatusCode: http.StatusOK,
	}
}
