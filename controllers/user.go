package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxwellhealth/bongo"
	"gitlab.com/gints/backend/models"
	"gitlab.com/gints/backend/utils"
	"gitlab.com/gints/backend/web/middleware"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserImpl struct {
	Router    *gin.Engine
	DB        *mgo.Database
	Generator *utils.JWTGenerator
	AuthWare  *middleware.AuthMiddleware
}

func NewUserImpl(router *gin.Engine, db *mgo.Database,
	generator *utils.JWTGenerator, auth *middleware.AuthMiddleware) *UserImpl {
	return &UserImpl{Router: router, DB: db, Generator: generator, AuthWare: auth}
}

func (ui *UserImpl) Initialize() {
	userGroup := ui.Router.Group("/user", ui.AuthWare.UsersMiddleware())
	userGroup.POST("/gint", ui.addGint)
	userGroup.DELETE("/gint", ui.deleteGint)
	userGroup.PUT("/profile", ui.updateUser)
}

func (ui *UserImpl) updateUser(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	email := input["email"].(string)

	update := bson.M{"$set": input}
	ui.DB.C("profile").UpdateId(email, update)
	c.JSON(http.StatusOK, gin.H{
		"status": "profile updated successfully",
	})
}

func (ui *UserImpl) deleteGint(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)
	var gint models.Gint
	GintID := input["gintid"]

	ui.DB.C("gints").FindId(GintID).One(&gint)

	ui.DB.C("gints").RemoveId(GintID)
	ui.DB.C(gint.Game + "_gints").RemoveId(GintID)

	selector := bson.M{"_id": gint.Email}
	update := bson.M{"$pull": bson.M{"gints": gint.GetId()}}
	ui.DB.C("profile").Update(selector, update)
}

func (ui *UserImpl) addGint(c *gin.Context) {
	var input map[string]interface{}
	c.BindJSON(&input)

	gint := models.Gint{
		Email:         c.Value("profile").(utils.JWTPayload).Email,
		Hint:          input["gintdesc"].(string),
		Game:          input["gamename"].(string),
		Approves:      0,
		Declines:      0,
		Inappropriate: false,
		Cheat:         false,
	}

	//Add to gints Collection
	err := ui.DB.C("gints").Insert(&gint)
	if err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			fmt.Println("Validation errors are:", vErr.Errors)
		} else {
			fmt.Println("Got a real error:", err.Error())
		}
	}

	//Add to user's gints Collection
	selector := bson.M{"_id": gint.Email}
	update := bson.M{"$push": bson.M{"gints": gint.GetId()}}
	ui.DB.C("profile").Update(selector, update)

	//Add to game's gints Collection
	gamegint := models.GameGints{}
	gamegint.SetId(gint.GetId())
	err2 := ui.DB.C(gint.Game + "_gint").Insert(&gamegint)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't save doc",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "gint added successfully",
	})
}
