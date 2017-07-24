package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	githuboauth "github.com/zalando/gin-oauth2/github"
)

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func setOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Options(sessions.Options{Path: "/random", Domain: "localhost", MaxAge: 2592000})
		c.Next()
	}
}

func loginHandler(c *gin.Context) {
	state := randToken()
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, githuboauth.GetLoginURL(state))
}

func (app *App) auth(c *gin.Context) {
	info := c.MustGet("user").(*github.User)
	user := User{}
	app.DB.Where(User{GitHubId: *info.ID}).Attrs(User{Model: Model{CreatedAt: time.Now()}, Name: *info.Name, Avatar: *info.AvatarURL}).FirstOrCreate(&user)

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (app *App) check() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session %s", userID))
			return
		}

		var user User
		if app.DB.First(&user, userID).Error == nil {
			c.Set("user", user)
		} else {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid user %s", userID))
		}
		c.Next()
	}
}
