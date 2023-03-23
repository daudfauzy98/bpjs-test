package main

import (
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Person struct {
	ID        uint    `gorm:"size:4;primaryKey" json:"id"`
	Customer  string  `gorm:"size:50" json:"customer"`
	Quantity  uint    `gorm:"size:4" json:"quantity"`
	Price     float64 `gorm:"size:6" json:"price"`
	Timestamp string  `json:"timestamp"`
}

func init() {
	dns := strings.Join([]string{
		"user=postgres",
		"password=postgres",
		"dbname=my_db",
		"port=5432",
		"sslmode=disable",
		"TimeZone=Asia/Jakarta",
	}, " ")
	d, err := gorm.Open(postgres.Open(dns), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	db = d
}

func migrate() {
	log.Println("[MESSAGE] migrating table..")
	err := db.Debug().AutoMigrate(&Person{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("[MESSAGE] migrating success!")
}

func dropTable() {
	log.Println("[MESSAGE] droping table..")
	err := db.Debug().Migrator().DropTable(&Person{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("[MESSAGE] droping success!")
}
