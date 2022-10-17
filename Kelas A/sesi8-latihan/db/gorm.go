package db

import (
	"fmt"
	"sesi8-latihan/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = "8881"
	user   = "nbcamp-user"
	pass   = "nbcamp-pass"
	dbname = "nbcamp-be"
)

func ConnectGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Person{})

	return db
}
