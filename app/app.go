package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nvs2394/just-bank-auth/domain"
	"github.com/nvs2394/just-bank-auth/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func getDBClient() *sqlx.DB {
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

func Start() {

	sanityCheck()
	router := mux.NewRouter()

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	dbClient := getDBClient()

	authRepositoryDB := domain.NewAuthRepositoryDb(dbClient)

	authHandlers := AuthHandlers{
		service: service.NewAuthService(authRepositoryDB, domain.GetRolePermissions()),
	}

	router.HandleFunc("/login", authHandlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/verify", authHandlers.Verify).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
