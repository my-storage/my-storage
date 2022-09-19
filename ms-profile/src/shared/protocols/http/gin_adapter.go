package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinAdapter[Body, Query any](controller Controller[Body, Query]) func(c *gin.Context) {
	return func(c *gin.Context) {
		param := c.Param

		body := getBody[Body](c)
		if body == nil {
			return
		}

		query := getQuery[Query](c)
		if query == nil {
			return
		}

		httpRequest := Request[Body, Query]{
			Body:  body,
			Query: query,
			Param: param,
		}

		httpResponse := controller.Handler(httpRequest)
		c.JSON(httpResponse.StatusCode, httpResponse.Data)
	}
}

func getBody[Body any](c *gin.Context) *Body {
	body := new(Body)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return body
}

func getQuery[Query any](c *gin.Context) *Query {
	query := new(Query)

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return query
}
