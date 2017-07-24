package main

import (
	"crypto/rand"
	"encoding/base64"
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

func loginHandler(c *gin.Context) {
	state := randToken()
	session := sessions.Default(c)
	session.Options(sessions.Options{Path: "/random", Domain: "localhost", MaxAge: 2592000})
	session.Set("state", state)
	session.Save()
	c.Redirect(http.StatusTemporaryRedirect, githuboauth.GetLoginURL(state))
}

func (app *App) auth(c *gin.Context) {
	info := c.MustGet("user").(*github.User)
	user := User{}
	app.DB.Where(User{GitHubId: *info.ID}).Attrs(User{Model: Model{CreatedAt: time.Now()}, Name: *info.Name, Avatar: *info.AvatarURL}).FirstOrCreate(&user)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
