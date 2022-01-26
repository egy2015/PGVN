package main

import (
	"log"
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct{
	gorm.Model

	Name string
	Email string `gorm:"typevarchar(100); unique_index"`
	Book []Book

}

type Book struct{
	gorm.Model

	Title string
	Author string
	BookId int `gorm:"typevarchar(100); "unique_index"`
	PersonID int
}

var db *gorm.DB
var err error

func main() {
	// Access ENV var
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//DB CONNECT string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s  port=%s ", host, user, dbname, password, dbport)

	//OPEN COnnection to db

	db, err := gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database")
	}

	//close connection
	defer db.Close()

	//make migrations to the db if not exist
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})



}