package main

import (
	"net/http"
	"strconv"
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
		random.POST("/user/import", app.importNameList)
		random.POST("/user/remove", app.removeNameList)
	}
}

func (app *App) auth(c *gin.Context) {
	info := c.MustGet("user").(*github.User)
	user := User{}
	app.DB.Where(User{GitHubId: *info.ID}).Attrs(User{Model: Model{CreatedAt: time.Now()}, Name: *info.Name}).FirstOrCreate(&user)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (app *App) getNameList(c *gin.Context) {
	id := c.Param("id")
	var nameList NameList
	if err := app.DB.First(&nameList, id).Error; err == nil {
		ok(c, nameList)
	} else {
		badRequest(c)
	}
}

func (app *App) createNameList(c *gin.Context) {
	id := c.MustGet("user").(*github.User).ID
	var user User
	app.DB.First(&user, id)
	var json NameList
	if c.BindJSON(&json) == nil {
		json.Creator = user
		app.DB.Model(&json).Create(&json)
		ok(c, json)
	} else {
		badRequest(c)
	}
}

func (app *App) updateNameList(c *gin.Context) {
	var json NameList
	if c.BindJSON(&json) == nil {
		app.DB.Model(&json).Select("Title", "Names", "Visibility").Update(&json)
		ok(c, json)
	} else {
		badRequest(c)
	}
}

func (app *App) getUser(c *gin.Context) {
	id := c.MustGet("user").(*github.User).ID
	var user User
	var nameLists []NameList
	if err := app.DB.First(&user, id).Error; err == nil {
		app.DB.Model(&user).Related(&nameLists, "NameLists")
		user.NameLists = nameLists
		ok(c, user)
	} else {
		badRequest(c)
	}
}

func (app *App) updateUser(c *gin.Context) {
	id := c.MustGet("user").(*github.User).ID
	var json User
	if c.BindJSON(&json) == nil {
		if *id != json.ID {
			forbidden(c)
			return
		}
		app.DB.Model(&json).Select("Name").Update(&json)
		ok(c, json)
	} else {
		badRequest(c)
	}
}

// params: id int, fork bool
func (app *App) importNameList(c *gin.Context) {
	id := c.MustGet("user").(*github.User).ID
	nid, err1 := strconv.Atoi(c.Param("id"))
	fork, err2 := strconv.ParseBool(c.Param("fork"))
	var user User
	var nameList NameList
	app.DB.First(&user, id)
	app.DB.First(&nameList, nid)
	if app.DB.Error == nil && err1 != nil && err2 != nil {
		tx := app.DB.Begin()
		if !fork {
			nameList.ID = nil
			tx.Create(&nameList)
		}
		tx.Model(&user).Update(map[string]interface{}{"NameLists": append(user.NameLists, nameList)})
		if tx.Error == nil {
			tx.Commit()
			ok(c, user)
		} else {
			tx.Rollback()
			badRequest(c)
		}
	} else {
		badRequest(c)
	}
}

func (app *App) removeNameList(c *gin.Context) {
	id := c.MustGet("user").(*github.User).ID
	nid, err := strconv.Atoi(c.Param("id"))
	var user User
	app.DB.First(&user, id)
	if app.DB.Error == nil && err != nil {
		n := user.NameLists[:0]
		for _, x := range user.NameLists {
			if x.ID != nid {
				n = append(n, x)
			}
		}
		app.DB.Model(&user).Update(map[string]interface{}{"NameLists": n})
		if app.DB.Error == nil {
			ok(c, user)
		} else {
			badRequest(c)
		}
	} else {
		badRequest(c)
	}
}

func ok(c *gin.Context, json interface{}) {
	c.JSON(http.StatusOK, Response{json, nil, nil})
}

func badRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{nil, gin.H{"status": "Bad request"}, nil})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{nil, gin.H{"status": "Unauthorized"}, nil})
}

func forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{nil, gin.H{"status": "Forbidden"}, nil})
}
