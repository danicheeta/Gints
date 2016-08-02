package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/maxwellhealth/bongo"
	"gitlab.com/gints/backend/models"
	"gitlab.com/gints/backend/utils"
	"gitlab.com/gints/backend/web/middleware"
)

type AdminImpl struct {
	Router    *gin.Engine
	DB        *mgo.Database
	Generator *utils.JWTGenerator
	AuthWare  *middleware.AuthMiddleware
}

type GameConf struct {
	bongo.DocumentBase `bson:",inline"` //unique game name
	Name               string           `json: "name,inline"`
	Geners             []string         `json: "geners,inline"` //genre names
	Achivements        []string         `json: "achivements,inline"`
	Banner             string           `json: "banner,inline"` //url
	Thumbnail          string           `json: "thumbnail,inline"`
	Description        string           `json: "description,inline"`
	Pinned             string           `json: "pinned,inline"` //gint id
	Users              int              `json: "users,inline"`
	K                  float32          `json: "k,inline"`
}

func NewAdminImpl(router *gin.Engine, db *mgo.Database,
	generator *utils.JWTGenerator, auth *middleware.AuthMiddleware) *AdminImpl {
	return &AdminImpl{Router: router, DB: db, Generator: generator, AuthWare: auth}
}

func (ai *AdminImpl) Initialize() {
	adminGroup := ai.Router.Group("/admin", ai.AuthWare.AdminMiddleware())
	adminGroup.PUT("/game/config", ai.SetGameConf)
	adminGroup.GET("/game/config", ai.GetGameConf)
}

func (ai *AdminImpl) SetGameConf(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)

	querier := bson.M{"name": input["name"]}
	change := bson.M{"$set": input}

	err2 := ai.DB.C("games").Update(querier, change)
	if err2 != nil {
		fmt.Println(err2)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "game's config updated successfully",
	})
}

func (ai *AdminImpl) GetGameConf(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	gameName := input["game"].(string)

	var game models.Game
	ai.DB.C("games").FindId(gameName).One(&game)

	c.JSON(http.StatusOK, game)
}
