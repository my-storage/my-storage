package auth

import (
	"github.com/gin-gonic/gin"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

func Register(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	auth.POST("/", protocols.GinAdapter(makeLoginController()))
}
