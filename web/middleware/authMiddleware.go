package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gints/backend/utils"
)

type AuthMiddleware struct {
	Generator *utils.JWTGenerator
}

func NewAuthMiddleware(generator *utils.JWTGenerator) *AuthMiddleware {
	return &AuthMiddleware{Generator: generator}
}

func (middleware *AuthMiddleware) UsersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		jwt := authHeader[7:len(authHeader)]
		fmt.Println("xxx", jwt)
		if middleware.Generator.ValidateJWT(jwt) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid jwt",
			})
			return
		}
		payload := middleware.Generator.Decode(jwt)
		if middleware.Generator.CheckExpire(payload.Expire) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "expired jwt",
			})
			return
		}
		c.Set("profile", payload)
		c.Next()
	}
}

func (middleware *AuthMiddleware) AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		jwt := authHeader[7:len(authHeader)]
		if !middleware.Generator.ValidateJWT(jwt) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid jwt",
			})
			return
		}
		payload := middleware.Generator.Decode(jwt)
		if middleware.Generator.CheckExpire(payload.Expire) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "expired jwt",
			})
			return
		}
		if payload.Admin == "user" { // payload.Admin != "admin" ???!?!?
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "not admin acess",
			})
			return
		}
		c.Set("profile", payload)
		c.Next()
	}
}

func (middleware *AuthMiddleware) MasterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		jwt := authHeader[7:len(authHeader)]
		if !middleware.Generator.ValidateJWT(jwt) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid jwt",
			})
			return
		}
		payload := middleware.Generator.Decode(jwt)
		if middleware.Generator.CheckExpire(payload.Expire) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "expired jwt",
			})
			return
		}
		if payload.Admin != "master" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "not master acess",
			})
			return
		}
		c.Set("profile", payload)
		c.Next()
	}
}
