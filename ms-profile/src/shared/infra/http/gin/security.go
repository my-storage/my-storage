package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my-storage/ms-profile/src/app/config"
)

const REQUEST_MAX_SIZE int64 = 5000000

func ContentSizeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, REQUEST_MAX_SIZE)
	}
}

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-XSS-Protection", "0")
		ctx.Writer.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		ctx.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Writer.Header().Set("X-Download-Options", "noopen")
		ctx.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		ctx.Writer.Header().Set("Strict-Transport-Security", "max-age=15552000; includeSubDomains")
		ctx.Writer.Header().Set("Referrer-Policy", "no-referrer")
		ctx.Writer.Header().Set("Origin-Agent-Cluster", " ?1")
		ctx.Writer.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		ctx.Writer.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		ctx.Writer.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		ctx.Writer.Header().Set("Content-Security-Policy", "default-src 'self';base-uri 'self';font-src 'self' https: data:;form-action 'self';frame-ancestors 'self';img-src 'self' data:;object-src 'none';script-src 'self';script-src-attr 'none';style-src 'self' https: 'unsafe-inline';upgrade-insecure-requests")
	}
}

func CorsMiddleware() gin.HandlerFunc {
	config := config.GetInstance()

	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", config.HttpAllowOrigin)
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
