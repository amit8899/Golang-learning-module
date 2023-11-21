package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string `gorm:"name"`
	Class string `gorm:"class"`
}

func main() {
	dsn := "root:toor@tcp(127.0.0.1:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// st1 := Student{Name: "Amit Rai", Class: "A"}

	db.AutoMigrate(&Student{})

	// res := db.Create(&st1)
	// fmt.Println(res.RowsAffected)

	// db.Select("Name", "Class").Create(&Student{Name: "Shiva", Class: "A"})

	var user Student
	db.First(&user)
	// db.Take(&user, 3)
	db.Last(&user)
	// fmt.Println(user)

	// var sts []Student
	// db.Find(&sts) // Get All
	// fmt.Println(sts)

	result := map[string]interface{}{}
	db.Model(&Student{}).Find(&result)
	fmt.Println(result)
}
