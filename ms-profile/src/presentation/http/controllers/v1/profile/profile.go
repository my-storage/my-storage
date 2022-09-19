package profile

import (
	"github.com/gin-gonic/gin"

	protocols "github.com/my-storage/ms-profile/src/shared/protocols/http"
)

func Register(router *gin.RouterGroup) {
	profile := router.Group("/profile")

	profile.POST("/", protocols.GinAdapter(makeRegisterProfileController()))
	profile.PUT("/:id", protocols.GinAdapter(makeUpdateProfileController()))
	profile.DELETE("/:id", protocols.GinAdapter(makeDeleteProfileController()))
	profile.GET("/:id", protocols.GinAdapter(makeFindProfileControllerById()))
}
