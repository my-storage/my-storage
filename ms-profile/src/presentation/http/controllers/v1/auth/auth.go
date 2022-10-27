package auth

import (
	"github.com/gin-gonic/gin"

	protocols "github.com/my-storage/ms-profile/src/shared/infra/http/gin"
)

func Register(router *gin.RouterGroup) {
	group := router.Group("/auth")

	group.POST("/", protocols.GinAdapter(makeLoginController()))
}
