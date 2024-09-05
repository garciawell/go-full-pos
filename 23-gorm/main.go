package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductInfo struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&ProductInfo{})

	// create
	// db.Create(&ProductInfo{
	// 	Name:  "Iphone",
	// 	Price: 1000.50,
	// })

	// // Create Batch
	// products := []ProductInfo{
	// 	{Name: "Notebook", Price: 120},
	// 	{Name: "Android", Price: 20},
	// 	{Name: "Keyboard", Price: 50},
	// }
	// db.Create(&products)

	// Select One

	var product ProductInfo
	// db.First(&product, 1)
	// fmt.Println(product)

	db.First(&product, "name = ?", "Iphone")
	fmt.Println(product)

	var products []ProductInfo
	db.Find(&products)

	for _, p := range products {
		fmt.Println(p)
	}
}
