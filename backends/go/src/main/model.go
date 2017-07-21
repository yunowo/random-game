package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"log"
)

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type User struct {
	Model
	Name      string  `json:"name"`
	Password  string `json:"-"`
	NameLists []NameList `json:"name_lists" gorm:"many2many:user_name_lists"`
}

type NameList struct {
	Model
	Title      string `json:"title"`
	Names      pq.StringArray  `json:"names" gorm:"type:varchar(100)[]"`
	Visibility int    `json:"visibility"`
	Users      []User `json:"-" gorm:"many2many:user_name_lists"`
}

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
	Meta  interface{} `json:"meta"`
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
