package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func FilterPathMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bannish(c.Request.URL.String()) {
			c.Abort()
			c.JSON(403, gin.H{
				"status": "no permission",
			})
		}
	}
}

func bannish(s string) bool {
	req := strings.ToLower(s)
	for _, v := range banwords {
		if strings.Contains(req, v) {
			return true
		}
	}
	return false
}

var banwords = []string{
	"$where",
	"$insert",
	"$select",
	"$in",
	"$ne",
}
