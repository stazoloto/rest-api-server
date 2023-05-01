package database

import (
	"fmt"
	"github.com/stazoloto/rest-api-server/internal/model"
	"log"
	"strconv"

	"github.com/stazoloto/rest-api-server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("error with port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("connection opened to database")

	err = DB.AutoMigrate(&model.Book{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Database migrated")
}
