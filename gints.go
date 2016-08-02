package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"

	"gitlab.com/gints/backend/models"
	"gitlab.com/gints/backend/utils"
	"gitlab.com/gints/backend/web"
	"gitlab.com/gints/backend/web/middleware"
)

func main() {
	router := gin.Default()
	router.Use(middleware.FilterPathMiddleware())

	session, err := mgo.Dial(models.MongoURI)
	if err != nil {
		log.Println(err)
		log.Fatalln("can't connect to database server")
	}
	defer session.Close()

	db := session.DB("gints")

	generator := utils.NewJWTGenerator(models.Secret)
	err = web.NewServer(models.Host, models.Port, router, db, generator)
	if err != nil {
		log.Println(err)
		log.Fatalln("cann't bind to host port")
	}
}
