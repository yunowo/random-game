package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/google/go-github/github"
	"net/http"
)

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
		json.ID = -1
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
			nameList.ID = -1
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