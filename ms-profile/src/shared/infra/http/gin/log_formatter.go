package gin

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/my-storage/ms-profile/src/app/config"
)

func LogFormatter() gin.HandlerFunc {
	config := config.GetInstance()

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}
		return fmt.Sprintf("[%v] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			config.AppName,
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	})
}
