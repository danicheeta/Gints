package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormChecker(keys ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, key := range keys {
			if c.Request.PostFormValue(key) == "" {
				c.Abort()
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "invalid jwt",
				})
                return
			}
		}
        c.Next()
	}
}
