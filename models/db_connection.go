package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB                *gorm.DB
	DATABASE_URL      = os.Getenv("DATABASE_URL")
	DATABASE_USERNAME = os.Getenv("DATABASE_USERNAME")
	DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	DATABASE_NAME     = os.Getenv("DATABASE_NAME")
)

func migrateDatabase() {
	DB.AutoMigrate(&User{})
}

func InitDatabase() {
	dataSourceName := fmt.Sprintf("host=%v user=%v password=%v port=5432", DATABASE_URL, DATABASE_USERNAME, DATABASE_PASSWORD) //sslmode=disable
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	isExistDatabaseNameStatement := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s'", DATABASE_NAME)
	result := db.Raw(isExistDatabaseNameStatement)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	// If database not exist
	var rec = make(map[string]interface{})
	if result.Find(rec); len(rec) == 0 {
		createDatabaseStatement := fmt.Sprintf("CREATE DATABASE %s;", DATABASE_NAME)
		if result := db.Exec(createDatabaseStatement); result.Error != nil {
			log.Fatal(result.Error)
		}
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			log.Fatal(err)
		}
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

	// Connect to selected database
	dataSourceName = fmt.Sprintf("host=%v user=%v password=%v port=5432 database=%v", DATABASE_URL, DATABASE_USERNAME, DATABASE_PASSWORD, DATABASE_NAME) //sslmode=disable
	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migrateDatabase()
	DB = db
}
