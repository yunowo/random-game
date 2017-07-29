package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type App struct {
	Engine *gin.Engine
	DB     *gorm.DB
	Config Config
}

func main() {
	config := loadConfig()

	gin.SetMode(gin.ReleaseMode)
	app := App{}
	app.Config = config
	app.initDB()
	app.Engine = gin.New()
	r := app.Engine
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	app.configRoutes()

	srv := &http.Server{
		Addr:    ":7000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	app.DB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
