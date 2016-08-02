package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hamed1soleimani/goauth"
	"gitlab.com/gints/backend/models"
	"gitlab.com/gints/backend/utils"
)

type OAuthImpl struct {
	Router    *gin.Engine
	Generator *utils.JWTGenerator
}

func NewOAuth(router *gin.Engine, generator *utils.JWTGenerator) *OAuthImpl {
	return &OAuthImpl{Router: router, Generator: generator}
}

func (oauth *OAuthImpl) Initialize() {
	auth := goauth.NewGOAuth()
	auth.Providers["google"] = models.Google
	oauth.Router.GET("/auth/:provider", auth.AuthHandler)
	oauth.Router.GET("/auth/:provider/oauth2callback", oauth.OauthMiddleware(), auth.CallbackHandler)
}

func (oauth *OAuthImpl) OauthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		profile := c.Value("profile").(goauth.Profile)
		c.JSON(http.StatusOK, oauth.Generator.GenerateJWT(profile.Email, "user"))
	}
}
