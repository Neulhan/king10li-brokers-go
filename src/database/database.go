package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Broker struct {
	gorm.Model
	Code           string `json:"code"`           // 등록번호
	Name           string `json:"name"`           // 상호
	Address        string `json:"address"`        // 소재지
	Representative string `json:"representative"` // 대표자
	PhoneNumber    string `json:"phoneNumber"`    // 전화번호
	Status         string `json:"status"`         // 상태
}

func DBLoader() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Broker{})

	// db.Create(&Broker{})

	// var broker Broker
	// db.First(&broker, 1)                 // find broker with integer primary key
	// db.First(&broker, "code = ?", "D42") // find broker with code D42

	// db.Model(&broker).Update("Price", 200)

	// db.Model(&broker).Updates(Broker{}) // non-zero fields
	// db.Model(&broker).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// db.Delete(&broker, 1)

	return db
}
