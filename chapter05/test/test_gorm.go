package main

import (
	"chapter05/common"
	"github.com/dzwvip/oracle"
	"gorm.io/gorm"
)

type Product struct {
	common.Model
	Code  string `gorm:"primarykey"`
	Price uint
}

func main() {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db, err := gorm.Open(oracle.Open("pdbadmin/oracle@192.168.31.23:1521/pdb1"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	//db.First(&product, 1)                  // find product with integer primary key
	db.First(&product, "code = :1", "F42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, "code = :1", "F42")
}
