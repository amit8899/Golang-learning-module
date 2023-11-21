package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=true&loc=Local" // sample
var dsn = "root:toor@tcp(127.0.0.1:3306)/gormdb?charset=utf8mb4&parseTime=true&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type GoTestModel struct {
	gorm.Model
	Name string
	Year string
}

func main() {
	http.HandleFunc("/createStuff", GoDatabaseCreate)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	GoTestModel := GoTestModel{
		Name: "Mike",
		Year: "2021",
	}

	db.Create(&GoTestModel)
	if err := db.Create(&GoTestModel).Error; err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(w).Encode(GoTestModel)

}
