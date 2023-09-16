package inits

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")
	HOST := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)

	sqlDB, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
