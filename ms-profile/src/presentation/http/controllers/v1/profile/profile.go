package profile

import (
	"github.com/gin-gonic/gin"

	protocols "github.com/my-storage/ms-profile/src/shared/infra/http/gin"
)

func Register(router *gin.RouterGroup) {
	group := router.Group("/profile")

	group.POST("/", protocols.GinAdapter(makeRegisterProfileController()))
	group.PUT("/:id", protocols.GinAdapter(makeUpdateProfileController()))
	group.DELETE("/:id", protocols.GinAdapter(makeDeleteProfileController()))
	group.GET("/:id", protocols.GinAdapter(makeFindProfileControllerById()))
}
