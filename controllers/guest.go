package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gints/backend/models"
	"gopkg.in/mgo.v2"
)

type GuestImp struct {
	Router *gin.Engine
	DB     *mgo.Database
}

func NewGuestImpl(router *gin.Engine, db *mgo.Database) *GuestImp {
	return &GuestImp{Router: router, DB: db}
}

func (gi *GuestImp) Initialize() {
	gi.Router.GET("/hashtags/:tag", gi.Hashtag)
	gi.Router.GET("/user/:id", gi.User)
	gi.Router.GET("/user/:id/gints", gi.getUsersGints)
	gi.Router.GET("/categories", gi.getCategories)
	gi.Router.GET("/categories/:id", gi.getSubCategories)
	gi.Router.GET("/games/:id/achivements", gi.GameAchivements)
	gi.Router.GET("/games/:id/gints", gi.GameGint)
	gi.Router.GET("/games/:id", gi.Game)
	gi.Router.GET("/games", gi.GetGames)
}

type Input struct {
	Set   int `form:"set" json:"set" binding:"required"`
	Limit int `form:"limit" json:"limit" binding:"required"`
}

func (gi *GuestImp) getUsersGints(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)

	usersName := input["name"]

	var user models.Profile
	var response []models.Gint
	var temp models.Gint

	gi.DB.C("profiles").FindId(usersName).One(&user)

	for _, gint := range user.Gints {
		gi.DB.C("gints").FindId(gint).One(&temp)
		response = append(response, temp)
	}
	c.JSON(http.StatusOK, response)
}

func (gi *GuestImp) getSubCategories(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	name := input["name"]

	var subcat models.Categorie
	var temp models.SubCategorie
	var response map[string][]string

	gi.DB.C("categories").FindId(name).One(&subcat)
	for _, sc := range subcat.Sub {
		gi.DB.C("subcategories").FindId(sc).One(&temp)
		response[subcat.Name] = append(response[subcat.Name], temp.Name)
	}
	c.JSON(http.StatusOK, response)
}

func (gi *GuestImp) getCategories(c *gin.Context) {
	var temp []models.Categorie
	gi.DB.C("categories").Find(nil).All(&temp)
	var categories []string
	for _, cat := range temp {
		categories = append(categories, cat.Name)
	}
	c.JSON(http.StatusOK, categories)
}

func (gi *GuestImp) GetGames(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	set := input["set"].(int)
	from := input["from"].(int)

	var games []models.Game

	gi.DB.C("games").Find(nil).Skip(from).Limit(set).All(&games)

	c.JSON(http.StatusOK, games)
}

func (gi *GuestImp) GameGint(c *gin.Context) {
	GameId := c.Param("id")

	var input map[string]interface{}
	c.BindJSON(&input)
	limit := input["limit"].(int)
	from := input["from"].(int)

	var gameGints []models.GameGints
	var gints []models.Gint
	var temp models.Gint

	gi.DB.C(GameId + "_gint").Find(nil).Skip(from).Limit(limit).All(&gameGints)

	for _, gint := range gameGints {
		gi.DB.C("gints").FindId(gint.GetId()).One(&temp)
		gints = append(gints, temp)
	}

	c.JSON(http.StatusOK, gints)
}

func (gi *GuestImp) Hashtag(c *gin.Context) {
	tag := c.Param("tag")

	var input map[string]interface{}
	c.BindJSON(&input)
	set := input["set"].(int)
	from := input["from"].(int)

	var hashSet []models.Hashtags
	gi.DB.C(tag + "_hashtag").Find(nil).All(&hashSet)

	var gints []models.Gint
	var temp models.Gint

	for i, hash := range hashSet {
		if i > (from-1)*set && i < from*set {
			gi.DB.C("gints").FindId(hash.GetId()).One(&temp)
			gints = append(gints, temp)
		}
	}
	c.JSON(http.StatusOK, gints)
}

func (gi *GuestImp) Game(c *gin.Context) {
	id := c.Param("id")
	var game models.Game
	gi.DB.C("games").FindId(id).One(&game)
	c.JSON(http.StatusOK, game)
}

func (gi *GuestImp) User(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	gi.DB.C("users").FindId(id).One(&user)
	c.JSON(http.StatusOK, user)
}

func (gi *GuestImp) GameAchivements(c *gin.Context) {
	id := c.Param("id")
	var game models.Game
	gi.DB.C("games").FindId(id).One(&game)
	var achivements []models.Achivement
	var temp models.Achivement
	for _, achivement := range game.Achivements {
		gi.DB.C("achivements").FindId(achivement).One(&temp)
		achivements = append(achivements, temp)
	}
	c.JSON(http.StatusOK, achivements)
}
