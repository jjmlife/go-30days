package main

import (
	model "day09/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDb() *gorm.DB {
	dsn := "root:jjm123456@tcp(127.0.0.1:3306)/gin_admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect error !")
	}
	return db
}

func insertDB() {
	/// 建表
	error := db.AutoMigrate(&model.User{})
	if error != nil {
		fmt.Println("Err ", error)
	}

	/// insert
	user := model.User{Username: "lisa", Password: "123456"}
	result := db.Create(&user)

	fmt.Println(result)
}

func queryDB() {
	var user model.User
  db.Where("username = ?","jjm").First(&user)
	fmt.Println("user query", user.Password)
}


func updateDB() {
	db.Model(&model.User{}).Where("username=?","lisa").Update("password","112233")
}



func main() {
	// testsql.SqlTest()
	db = connectDb()

	insertDB()

	queryDB()

	updateDB()
}