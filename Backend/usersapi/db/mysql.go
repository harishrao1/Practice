package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"userapi/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQL(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("MySQL connection error: %v", err)
	}

	// Connection pool settings
	database.SetMaxOpenConns(25)
	database.SetMaxIdleConns(25)
	database.SetConnMaxLifetime(5 * time.Minute)

	if err := database.Ping(); err != nil {
		log.Fatalf("MySQL ping error: %v", err)
	}

	DB = database
	log.Println("✅ MySQL connected")
}
