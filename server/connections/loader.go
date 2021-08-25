package connections

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"../models/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	dbURI := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbName, username, password)
	if err != nil {
		panic(err)
	}

	customLogger := logger.New(
		log.New(ioutil.Discard, "\r\n", log.LstdFlags),
		logger.Config{},
	)

	conn, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{Logger: customLogger, FullSaveAssociations: true})
	if err != nil {
		panic(err)
	}

	conn.Debug().AutoMigrate(
		&domain.Account{},
		&domain.Version{},
		&domain.Library{},
		&domain.Language{},
		&domain.Tree{},
		&domain.Implementation{},
		&domain.Tag{},
		&domain.Description{},
		&domain.Test{},
	)
	return conn
}
