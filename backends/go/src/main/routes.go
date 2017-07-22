package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	githuboauth "github.com/zalando/gin-oauth2/github"
)

func (app *App) configRoutes() {
	r := app.Engine

	secret := []byte("a secret")
	scopes := []string{"user:email", }
	sessionName := "session"
	githuboauth.Setup(RedirectURL, ConfigFile, scopes, secret)

	r.Use(githuboauth.Session(sessionName))
	r.GET("/random/login", githuboauth.LoginHandler)

	random := r.Group("/random")
	random.Use(githuboauth.Auth())
	{
		random.GET("/auth", app.auth)

		random.GET("/namelist/:id", app.getNameList)
		random.POST("/namelist", app.createNameList)
		random.PATCH("/namelist/:id", app.updateNameList)

		random.GET("/user", app.getUser)
		random.PATCH("/user", app.updateUser)
		random.PATCH("/user/import", app.updateUserNameLists)
	}
}

func (app *App) auth(c *gin.Context) {
	info := c.MustGet("user").(*github.User)
	user := User{}
	app.DB.Where(User{GitHubId: *info.ID}).Attrs(User{Model: Model{CreatedAt: time.Now()}, Name: *info.Name}).FirstOrCreate(&user)
	c.Redirect(http.StatusTemporaryRedirect, "/")
	//c.JSON(http.StatusOK, gin.H{"user": c.MustGet("user")})
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
	//id, _ := strconv.Atoi(c.Param("id"))
	id := c.MustGet("user").(*github.User).ID
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
	id := c.MustGet("user").(*github.User).ID
	var json User
	if c.BindJSON(&json) == nil {
		if *id != json.ID {
			responseForbidden(c)
			return
		}
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
	c.JSON(http.StatusUnauthorized, Response{nil, gin.H{"status": "Unauthorized"}, nil})
}

func responseForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{nil, gin.H{"status": "Forbidden"}, nil})
}
