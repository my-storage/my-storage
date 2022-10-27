package profile

import (
	"net/http"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

type FindProfileControllerById struct {
	protocols.Controller[any]
}

func makeFindProfileControllerById() protocols.Controller[any] {
	return &FindProfileControllerById{}
}

func (a *FindProfileControllerById) Handler(request protocols.Request[any]) protocols.Response {
	id := request.Param("id")

	return protocols.Response{
		Data:       id,
		StatusCode: http.StatusOK,
	}

}
