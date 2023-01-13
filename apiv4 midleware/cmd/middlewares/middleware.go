package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	TOKEN := os.Getenv("TOKEN")
	return func(ctx *gin.Context) {
		receivedToken := ctx.GetHeader("token")
		if TOKEN != receivedToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		ctx.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Request.URL
		host := ctx.Request.Host
		meth := ctx.Request.Method
		t := time.Now().Local()
		size := ctx.Request.ContentLength
		fmt.Println("\nMetodo: ", meth, " URL: ", host, url, " Time: ", t, " Bytes: ", size)
		ctx.Next()
	}
}
