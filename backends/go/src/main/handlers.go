package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	user := c.MustGet("user").(User)
	var json Request
	if c.BindJSON(&json) != nil {
		badRequest(c)
		return
	}

	nameList := json.Data
	if nameList.Title == "" || nameList.Names == nil {
		badRequest(c)
	}
	nameList.CreatorID = user.ID

	tx := app.DB.Begin()
	tx.Set("gorm:insert_option", "ON CONFLICT (id) DO NOTHING").Create(&nameList)
	tx.Exec("INSERT INTO user_name_lists (user_id, name_list_id) VALUES (?, ?)", user.ID, nameList.ID)
	if tx.Error == nil {
		tx.Commit()
		ok(c, nameList)
	} else {
		tx.Rollback()
		badRequest(c)
	}
}

func (app *App) updateNameList(c *gin.Context) {
	var json Request
	if c.BindJSON(&json) != nil {
		badRequest(c)
		return
	}

	nameList := json.Data
	app.DB.Model(&nameList).Select("Title", "Names", "Visibility").Update(&nameList)
	ok(c, nameList)
}

func (app *App) getUser(c *gin.Context) {
	user := c.MustGet("user").(User)
	ok(c, user)
}

// params: id int, fork bool
func (app *App) importNameList(c *gin.Context) {
	user := c.MustGet("user").(User)
	nid, err1 := strconv.Atoi(c.PostForm("id"))
	fork, err2 := strconv.ParseBool(c.PostForm("fork"))
	var nameList NameList
	if app.DB.First(&nameList, nid).Error != nil || err1 != nil || err2 != nil {
		notFound(c)
		return
	}

	if nameList.Visibility == 1 {
		forbidden(c)
		return
	}

	tx := app.DB.Begin()
	if fork {
		nameList.ID = 0
		nameList.CreatorID = user.ID
		tx.Set("gorm:insert_option", "ON CONFLICT (id) DO NOTHING").Create(&nameList)
	}
	if tx.Exec("INSERT INTO user_name_lists (user_id, name_list_id) VALUES (?, ?)", user.ID, nameList.ID).Error == nil {
		tx.Commit()
		ok(c, nameList)
	} else {
		tx.Rollback()
		badRequest(c)
	}
}

func (app *App) removeNameList(c *gin.Context) {
	user := c.MustGet("user").(User)
	nid, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		badRequest(c)
		return
	}

	app.DB.Exec("DELETE FROM user_name_lists WHERE user_id = ? AND name_list_id = ?", user.ID, nid)
	if app.DB.Error == nil {
		ok(c, user)
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

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{nil, gin.H{"status": "NotFound"}, nil})
}
