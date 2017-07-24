package main

import (
	"time"

	"github.com/gin-contrib/cors"
	githuboauth "github.com/zalando/gin-oauth2/github"
)

func (app *App) configRoutes() {
	r := app.Engine

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://app.liuyun.me"},
		AllowMethods:     []string{"GET", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	secret := []byte("a secret")
	scopes := []string{"user:email", }
	sessionName := "rnd_session"
	githuboauth.Setup(RedirectURL, ConfigFile, scopes, secret)
	r.Use(githuboauth.Session(sessionName))

	r.GET("/random/login", loginHandler)

	random := r.Group("/random")
	random.Use(githuboauth.Auth())
	{
		random.GET("/auth", app.auth)

		random.GET("/namelist/:id", app.getNameList)
		random.POST("/namelist", app.createNameList)
		random.PATCH("/namelist/:id", app.updateNameList)

		random.GET("/user", app.getUser)
		random.POST("/user/import", app.importNameList)
		random.POST("/user/remove", app.removeNameList)
	}
}
