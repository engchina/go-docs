package main

import (
	"fmt"
	"github.com/engchina/oracle"
	"gorm.io/gorm"
	"time"
)

type Region struct {
	//common.Model
	RegionId    string `gorm:"primarykey"`
	RegionName  string
	UpdateCount int64
}

func main() {
	db, err := gorm.Open(oracle.Open("oracle://pdbadmin:oracle@192.168.31.23:1521/pdb1"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//db.AutoMigrate(&Region{})

	db.Transaction(func(tx *gorm.DB) error {
		// Create
		db.Create(&Region{RegionId: "ap-tokyo-1", RegionName: "Tokyo", UpdateCount: 0})

		time.Sleep(15 * time.Second)

		// Read
		var region Region
		db.Debug().First(&region, "region_id = :1", "ap-tokyo-1")
		fmt.Printf("### regions is: %#v ###", region)

		// Update - update multiple fields
		db.Debug().Model(&region).Updates(Region{RegionName: "TOKYO", UpdateCount: region.UpdateCount + 1}) // non-zero fields

		time.Sleep(15 * time.Second)

		// Read
		db.Debug().First(&region, "region_id = :1", "ap-tokyo-1")
		fmt.Printf("### regions is: %#v ###", region)

		time.Sleep(15 * time.Second)

		// Delete - delete region
		db.Delete(&region, "region_id = :1", "ap-tokyo-1")

		// return nil will commit the whole transaction
		return nil
	})
}
