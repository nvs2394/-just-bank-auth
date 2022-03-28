package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func Start() {

	sanityCheck()

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	router := NewRouter()

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)

	if err != nil {
		log.Fatal("Can not start server")
	}

}
