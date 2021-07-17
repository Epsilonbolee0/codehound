package connections

import (
	"fmt"
	"os"

	"../models/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetConnection(role string) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	username := os.Getenv("db_login_" + role)
	password := os.Getenv("db_password_" + role)

	fmt.Println(role)
	fmt.Println(username)

	dbURI := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbName, username, password)
	if err != nil {
		panic(err)
	}

	conn, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn.Debug().AutoMigrate(&domain.Account{})
	return conn
}
