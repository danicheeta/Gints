package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"gitlab.com/gints/backend/models"
	"gitlab.com/gints/backend/utils"
	"gitlab.com/gints/backend/web/middleware"
)

type MasterImpl struct {
	Router    *gin.Engine
	DB        *mgo.Database
	Generator *utils.JWTGenerator
	AuthWare  *middleware.AuthMiddleware
}

func NewMasterImpl(router *gin.Engine, db *mgo.Database,
	generator *utils.JWTGenerator, auth *middleware.AuthMiddleware) *MasterImpl {
	return &MasterImpl{Router: router, DB: db, Generator: generator, AuthWare: auth}
}

func (mi *MasterImpl) Initialize() {
	masterGroup := mi.Router.Group("/master" /*, mi.AuthWare.MasterMiddleware()*/)
	masterGroup.POST("/game", mi.addGame)
	masterGroup.PUT("/game", mi.updateGame)
	masterGroup.DELETE("/game", mi.delGame)
	masterGroup.POST("/category", mi.addCat)
	masterGroup.DELETE("/category", mi.delCat)
	masterGroup.GET("/admin", mi.getAdmins)
	masterGroup.PUT("/admin", mi.updateAdmins)
	masterGroup.DELETE("/admin", mi.delAdmins)
}

func (mi *MasterImpl) updateAdmins(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)

	GameName := input["gamename"].(string)
	Admins := input["admins"].([]string)

	update := bson.M{"$set": Admins}
	mi.DB.C("admin").UpdateId(GameName, update)
	c.JSON(http.StatusOK, gin.H{
		"status": "admins updated successfully",
	})
}

func (mi *MasterImpl) delAdmins(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	admin := input["admin"].(string)
	game := input["game"].(string)

	update := bson.M{"$pull": bson.M{"admins": admin}}
	mi.DB.C("admin").UpdateId(game, update)
	c.JSON(http.StatusOK, gin.H{
		"status": "admin deleted successfully",
	})
}

func (mi *MasterImpl) getAdmins(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	game := input["game"].(string)

	var adminModel models.Admin
	err := mi.DB.C("admin").FindId(game).One(&adminModel)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": "game or admin wasn't found",
		})
		return
	}

	c.JSON(http.StatusOK, adminModel.Admins)
}

func (mi *MasterImpl) updateGame(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)

	name := input["name"].(string)

	update := bson.M{"$set": input}
	err := mi.DB.C("games").UpdateId(name, update)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": "game updating had some problems",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "game updated successfully",
	})
}

func (mi *MasterImpl) delGame(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	name := input["name"].(string)

	err := mi.DB.C("games").RemoveId(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "game was not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "game deleted successfully",
	})
}

func (mi *MasterImpl) delCat(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	name := input["name"].(string)

	mi.DB.C("categories").RemoveId(name)
	c.JSON(http.StatusOK, gin.H{
		"status": "category deleted successfully",
	})
}

func (mi *MasterImpl) addCat(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	name := input["name"].(string)

	var cat models.Categorie
	cat.Name = name

	mi.DB.C("categories").Insert(&cat)
	c.JSON(http.StatusOK, gin.H{
		"status": "category added successfully",
	})
}

func (mi *MasterImpl) addGame(c *gin.Context) {
	var game models.Game
	err := c.BindJSON(&game)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unexpected PUT parameters",
		})
		return
	}

	err = mi.DB.C("games").FindId(game.Name).One(nil)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "Game already exists",
		})
		return
	}

	err = mi.DB.C("games").Insert(&game)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Game added successfully",
	})
	an := models.Game{
		Name: "kir",
	}
	mi.DB.C("games").Insert(&an)
}
