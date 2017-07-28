package main

import (
	"time"

	"github.com/gin-contrib/cors"
	oauth2 "github.com/zalando/gin-oauth2/github"
)

func (app *App) configRoutes() {
	r := app.Engine

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{app.Config.AppEndpoint},
		AllowMethods:     []string{"GET", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	secret := []byte("a secret")
	scopes := []string{"user:email", }
	sessionName := "rnd_session"
	oauth2.Setup(app.Config.ApiEndpoint+"/random/auth", ClientConfigFile, scopes, secret)
	r.Use(oauth2.Session(sessionName))
	r.Use(app.setOptions())

	r.GET("/random/login", loginHandler)

	auth := r.Group("/random")
	auth.Use(oauth2.Auth())
	auth.GET("/auth", app.auth)

	random := r.Group("/random")
	random.Use(app.check())
	{
		random.GET("/namelist/:id", app.getNameList)
		random.POST("/namelist", app.createNameList)
		random.PATCH("/namelist/:id", app.updateNameList)

		random.GET("/user", app.getUser)
		random.POST("/user/import", app.importNameList)
		random.POST("/user/remove", app.removeNameList)
	}
}
