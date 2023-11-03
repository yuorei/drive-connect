package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"drive-connect/db/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Conn *gorm.DB
}

func ConnectionDB() *DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	TimeZone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, TimeZone)
	fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		Conn: db,
	}
}

func (db *DB) CreateTable() {
	if err := db.Conn.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Database create table failed")
	}

	if err := db.Conn.AutoMigrate(&model.Comment{}); err != nil {
		log.Fatalf("Database create table failed")
	}

	if err := db.Conn.AutoMigrate(&model.Board{}); err != nil {
		log.Fatalf("Database create table failed")
	}
}
