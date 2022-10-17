package config

import (
	"fmt"
	"sesi7/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	if !db.Migrator().HasTable(model.Product{}) {
		db.Debug().AutoMigrate(model.Product{})
	}

	if !db.Migrator().HasColumn(model.Product{}, "title") {
	}
	db.Debug().AutoMigrate(model.Product{})

	return db, nil
}
