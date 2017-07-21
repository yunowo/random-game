package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/zalando/gin-oauth2/github"
)

func (app *App) configRoutes() {
	r := app.Engine

	secret := []byte("a secret")
	scopes := []string{"repo", "user:email", }
	sessionName := "randomgosession"
	github.Setup(RedirectURL, ConfigFile, scopes, secret)

	r.Use(github.Session(sessionName))
	r.POST("/random/login", github.LoginHandler)

	random := r.Group("/random")
	random.Use(github.Auth())
	{
		random.GET("/auth",app.auth)

		random.GET("/namelist/:id", app.getNameList)
		random.POST("/namelist", app.createNameList)
		random.PATCH("/namelist/:id", app.updateNameList)

		random.GET("/user/:id", app.getUser)
		random.PATCH("/user/:id", app.updateUser)
		random.PATCH("/user/:id/relationships/namelists", app.updateUserNameLists)
	}
}

func (app *App) auth(c *gin.Context) {

}

func (app *App) getNameList(c *gin.Context) {
	id := c.Param("id")
	var nameList NameList
	if err := app.DB.First(&nameList, id).Error; err == nil {
		responseOK(c, nameList)
	} else {
		responseBadRequest(c)
	}
}

func (app *App) createNameList(c *gin.Context) {
	var json NameList
	if c.BindJSON(&json) == nil {
		app.DB.Model(&json).Create(&json)
		responseOK(c, json)
	} else {
		responseBadRequest(c)
	}
}

func (app *App) updateNameList(c *gin.Context) {
	var json NameList
	if c.BindJSON(&json) == nil {
		app.DB.Model(&json).Update(&json)
		responseOK(c, json)
	} else {
		responseBadRequest(c)
	}
}

func (app *App) getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	var nameLists []NameList
	if err := app.DB.First(&user, id).Error; err == nil {
		app.DB.Model(&user).Related(&nameLists, "NameLists")
		user.NameLists = nameLists
		responseOK(c, user)
	} else {
		responseBadRequest(c)
	}
}

func (app *App) updateUser(c *gin.Context) {
	var json User
	if c.BindJSON(&json) == nil {
		app.DB.Model(&json).Update(&json)
		responseOK(c, json)
	} else {
		responseBadRequest(c)
	}
}

func (app *App) updateUserNameLists(c *gin.Context) {
	var json User
	if c.BindJSON(&json) == nil {
		app.DB.Model(&json).Update(&json)
		responseOK(c, json)
	} else {
		responseBadRequest(c)
	}
}

func responseOK(c *gin.Context, json interface{}) {
	c.JSON(http.StatusOK, Response{json, nil, nil})
}

func responseBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{nil, gin.H{"status": "Bad request"}, nil})
}

func responseUnauthorized(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{nil, gin.H{"status": "Unauthorized"}, nil})
}
