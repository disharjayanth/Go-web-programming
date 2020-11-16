package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	dsn := "user=gwp dbname=gwp password=123 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{
		Content: "Hello World",
		Author:  "SauSheong Chang",
	}
	Db.Create(&post)

	fmt.Println("After insertion of Post:", post)

	comment := Comment{
		Content: "Good one sir",
		Author:  "Student",
	}
	Db.Model(&post).Association("Comments").Append(&comment)

	var readPost Post
	Db.Model(&post).Where("author = ?", "SauSheong Chang").First(&readPost)
	fmt.Println("Reading post from Post TABLE:", readPost)

	// Need to study GORM
}
