package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("settin up new database connection")

	dbUsername := "postgres" //os.Getenv("DB_USERNAME")
	dbPassword := "postgres" //os.Getenv("DB_PASSWORD")
	dbHost := "localhost"    //os.Getenv("DB_HOST")
	dbTable := "postgres"    //os.Getenv("DB_TABLE")
	dbPort := "5432"         //os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbTable, dbPassword)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
