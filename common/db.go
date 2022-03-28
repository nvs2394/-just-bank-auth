package common

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBClient() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connectionString := dbUser + ":" + dbPassword + "@/" + dbName

	client, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return client
}
