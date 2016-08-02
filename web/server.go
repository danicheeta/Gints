package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/gints/backend/controllers"
	"gitlab.com/gints/backend/utils"
	"gitlab.com/gints/backend/web/middleware"
	"gopkg.in/mgo.v2"
)

func NewServer(host string, port string, router *gin.Engine,
	db *mgo.Database, generator *utils.JWTGenerator) (err error) {
	auth := middleware.NewAuthMiddleware(generator)
	authctl := controllers.NewAuth(router, db, generator, auth)
	authctl.Initialize()
	oauthctl := controllers.NewOAuth(router, generator)
	oauthctl.Initialize()
	Guestimpl := controllers.NewGuestImpl(router, db)
	Guestimpl.Initialize()
	UserImpl := controllers.NewUserImpl(router, db, generator, auth)
	UserImpl.Initialize()
	AdminImpl := controllers.NewAdminImpl(router, db, generator, auth)
	AdminImpl.Initialize()
	MasterImpl := controllers.NewMasterImpl(router, db, generator, auth)
	MasterImpl.Initialize()

	addr := host + ":" + port

	log.Printf("Listening on %s...", addr)
	router.Run(addr)

	return
}
