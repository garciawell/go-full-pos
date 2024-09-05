package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primarykey"`
	Name string
}

type ProductInfo struct {
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&ProductInfo{}, &Category{})

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// db.Create(&ProductInfo{
	// 	Name:       "Notebook",
	// 	Price:      1000,
	// 	CategoryID: category.ID,
	// })

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

	// var product ProductInfo
	// db.First(&product, 1)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Iphone")
	// fmt.Println(product)

	var products []ProductInfo
	// db.Limit(2).Offset(2).Find(&products)
	db.Preload("Category").Find(&products)

	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name)
	}

	// var products []ProductInfo
	// db.Where("price > ?", 1000).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// var p ProductInfo
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 ProductInfo
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)
}
