package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/my-storage/ms-profile/src/app/config"
	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

func NewGinEngine() *gin.Engine {
	config := config.GetInstance()

	if config.ApiMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	return router
}

func GinAdapter[Body any](controller protocols.Controller[Body]) func(c *gin.Context) {
	validate := validator.New()

	return func(c *gin.Context) {
		param := c.Param
		query := c.Query
		body := GetBody[Body](c, validate)

		httpRequest := protocols.Request[Body]{
			Body:  body,
			Query: query,
			Param: param,
		}

		httpResponse := controller.Handler(httpRequest)

		if httpResponse.Headers != nil {
			SetHeaders(c, httpResponse.Headers)
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
