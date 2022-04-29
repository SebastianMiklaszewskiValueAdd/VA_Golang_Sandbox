package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(r *gin.Engine) {
	helloWorld(r)
}

func helloWorld(r *gin.Engine) {
	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
}
