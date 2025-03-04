package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")
	db.Model(&product).Update("Price", 200)

	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// db.Delete(&product, 1)

	var products []Product
	db.Find(&products)

	fmt.Println("Updated Product List:")
	for _, p := range products {
		fmt.Printf("ID: %d, Code: %s, Price: %d, CreatedAt: %s, UpdatedAt: %s\n", p.ID, p.Code, p.Price, p.CreatedAt, p.UpdatedAt)
	}

}
