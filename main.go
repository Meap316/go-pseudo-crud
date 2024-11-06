package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Open Database Connection
	dsn := "host=68.183.179.230 user=admin password=0qP2JDV3EVkEMEd dbname=latihan_irvan port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	//
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Database Migration

	r := gin.Default()

	r.Run() // listen and serve on 0.0.0.0:8080
}
