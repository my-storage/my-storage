package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/my-storage/ms-profile/src/shared/infra/http/gin/helpers"
	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

func GinAdapter[Body any](controller protocols.Controller[Body]) func(c *gin.Context) {
	validate := validator.New()

	return func(c *gin.Context) {
		param := c.Param
		query := c.Query
		body := helpers.GetBody[Body](c, validate)

		httpRequest := protocols.Request[Body]{
			Body:  body,
			Query: query,
			Param: param,
		}

		httpResponse := controller.Handler(httpRequest)

		if httpResponse.Headers != nil {
			helpers.SetHeaders(c, httpResponse.Headers)
		}

		if httpResponse.Data != nil {
			c.JSON(httpResponse.StatusCode, httpResponse.Data)
			return
		}

		if httpResponse.StatusCode == http.StatusPermanentRedirect || httpResponse.StatusCode == http.StatusTemporaryRedirect {
			c.Redirect(httpResponse.StatusCode, httpResponse.Redirect)
			return
		}

		c.AbortWithStatus(httpResponse.StatusCode)
	}
}
