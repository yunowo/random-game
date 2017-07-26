package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
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
	app.DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres")
	if err != nil {
		log.Fatalf("Got error when connect database, '%v'", err)
	}

	app.DB.AutoMigrate(&User{}, &NameList{})

	app.DB.LogMode(true)

	//namelist := NameList{}
	//namelist.Names = pq.StringArray{"222", "2332", "3333"}
	//namelist.Title = "This"
	//user := User{}
	//user.Name = "?.?"
	//user.Password = "?.?"
	//user.CreatedAt = time.Now()
	//user.NameLists = []NameList{namelist}
	//app.DB.Create(&user)
}
