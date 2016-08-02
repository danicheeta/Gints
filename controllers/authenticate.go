package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/gints/backend/models"
	"gitlab.com/gints/backend/utils"
	"gitlab.com/gints/backend/web/middleware"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserAuthImpl struct {
	Router    *gin.Engine
	DB        *mgo.Database
	Generator *utils.JWTGenerator
	AuthWare  *middleware.AuthMiddleware
}

func NewAuth(router *gin.Engine, db *mgo.Database,
	generator *utils.JWTGenerator, auth *middleware.AuthMiddleware) *UserAuthImpl {
	return &UserAuthImpl{Router: router, DB: db, Generator: generator, AuthWare: auth}
}

func (ui *UserAuthImpl) Initialize() {

	ui.Router.POST("/login", middleware.FormChecker("email", "password"), ui.login)
	ui.Router.POST("/register", ui.register)
	ui.Router.POST("/auth/renew", ui.renew)
	ui.Router.POST("/auth/rstpass", ui.AuthWare.UsersMiddleware(), ui.resetPassword)
}

func (ui *UserAuthImpl) login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	var user models.User
	err := ui.DB.C("user").FindId(email).One(&user)

	if err == nil {
		match := strings.Compare(user.Password, password)
		if match == 0 && user.Activated != false {
			c.JSON(http.StatusOK, gin.H{
				"jwt": ui.Generator.GenerateJWT(email, "user"),
			})
		} else if user.Activated == true {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "not activated account",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "incorrect password",
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found!",
		})
	}
}

func (ui *UserAuthImpl) resetPassword(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	newpass := c.PostForm("newpass")

	var user models.User
	err := ui.DB.C("user").FindId(email).One(&user)
	if err == nil {
		match := strings.Compare(user.Password, password)
		if match == 0 && user.Activated {
			update := bson.M{"$set": bson.M{"password": newpass, "activated": true}}
			ui.DB.C("user").UpdateId(email, update)
		} else if user.Activated == true {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "not activated account",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "incorrect password",
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found!",
		})
	}
}

func (ui *UserAuthImpl) register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	var user models.User
	err := ui.DB.C("user").FindId(email).One(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user already exists",
		})
		return
	}

	user.Password = password
	user.Activated = false
	ui.DB.C("user").Insert(&user)

	//we have to send activation mail here
	c.JSON(http.StatusUnauthorized, gin.H{
		"activate": "check your email and activate your account",
	})
}

func (ui *UserAuthImpl) renew(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	jwt := authHeader[7:len(authHeader)]
	payload := ui.Generator.Decode(jwt)
	if ui.Generator.CheckReLogin(payload.Expire) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "login agian",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"jwt": ui.Generator.RenewJWT(jwt),
		})
	}
}
