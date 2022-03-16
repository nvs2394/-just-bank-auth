package common

import (
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

func GetDBClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connectionString := dbUser + ":" + dbPassword + "@/" + dbName

	client, err := sqlx.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
