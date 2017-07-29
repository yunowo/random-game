package main

import (
	"log"
	"time"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"os"
)

type Model struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type User struct {
	Model
	GitHubID  int        `json:"-" gorm:"column:github_id"`
	Avatar    string     `json:"avatar"`
	Name      string     `json:"name"`
	NameLists []NameList `json:"name_lists" gorm:"many2many:user_name_lists"`
}

type NameList struct {
	Model
	Title      string         `json:"title"`
	Names      pq.StringArray `json:"names" gorm:"type:varchar(100)[]"`
	Visibility int            `json:"visibility"`
	Users      []User         `json:"-" gorm:"many2many:user_name_lists"`
	CreatorID  int            `json:"creator_id" gorm:"column:creator_id"`
}

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
	Meta  interface{} `json:"meta"`
}

type Request struct {
	Data NameList `json:"data"`
}

func (app *App) initDB() {
	var err error
	arg := fmt.Sprintf("host=%v user=%v dbname=%v sslmode=disable password=%v", "localhost", os.Getenv("RND_DB_USER"),os.Getenv("RND_DB_NAME"), os.Getenv("RND_DB_PWD"))
	app.DB, err = gorm.Open("postgres", arg)
	if err != nil {
		log.Fatalf("Got error when connect database, '%v'", err)
	}

	app.DB.AutoMigrate(&User{}, &NameList{})

	app.DB.LogMode(true)
}
