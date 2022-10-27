package helpers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/my-storage/ms-profile/src/shared/aggregators/errors"
)

func getStack() []byte {
	buffer := make([]byte, 1024)
	for {
		level := runtime.Stack(buffer, false)
		if level < len(buffer) {
			return buffer[:level]
		}

		buffer = make([]byte, 2*len(buffer))
	}
}

func ErrorHandler() gin.HandlerFunc {
	logger := log.New(gin.DefaultWriter, "", log.LstdFlags|log.Llongfile)

	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(*errors.AppError); ok {
					ctx.JSON(*e.GetStatusCode(), gin.H{
						"error": map[string]any{
							"name":    e.Name,
							"message": e.Description,
							"details": e.Details,
						},
					})

					ctx.Abort()
					return
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				logger.Printf("|ErrorHandler| panic recovered:\n%s\n%s\n%s", string(httpRequest), err, getStack())

				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": map[string]any{
						"name":    errors.InternalServerError,
						"message": "Internal server error",
					},
				})

				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}
