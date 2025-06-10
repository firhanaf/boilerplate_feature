package database

import (
	"boilerplate-feature/app/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	Db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return Db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate()
}
