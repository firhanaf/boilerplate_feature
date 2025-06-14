package database

import (
	"boilerplate-feature/app/config"
	Userdata "boilerplate-feature/features/users/data"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)

	Db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return Db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&Userdata.User{})
}
