package config

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseInit() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:rakamin@tcp(localhost:3306)/coba")

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	fmt.Println("Connected to database")

	return db, nil
}

func GormDatabaseConn() (*gorm.DB, error) {
	dsn := "root:rakamin@tcp(localhost:3306)/coba?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Cannot connect db", err)
		return nil, err
	}

	return db, nil
}
