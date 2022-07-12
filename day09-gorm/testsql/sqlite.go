package testsql

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}


func SqlTest() {
	db, err := gorm.Open(sqlite.Open("test.db"),&gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect sqlite")
		return
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code:"12",Price: 100})

	var product Product

	db.First(&product,1)
	fmt.Println(product)



}
