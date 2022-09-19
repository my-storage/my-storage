package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/my-storage/ms-profile/src/presentation/http/controllers/v1/auth"
	"github.com/my-storage/ms-profile/src/presentation/http/controllers/v1/profile"
)

func Register(router *gin.RouterGroup) {
	v1 := router.Group("v1")

	auth.Register(v1)
	profile.Register(v1)
}
