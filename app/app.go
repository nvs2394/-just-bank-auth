package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nvs2394/just-bank-auth/common"
	"github.com/nvs2394/just-bank-auth/domain"
	"github.com/nvs2394/just-bank-auth/service"
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

	dbClient := common.GetDBClient()

	authRepositoryDB := domain.NewAuthRepositoryDb(dbClient)

	authHandlers := AuthHandlers{
		service: service.NewAuthService(authRepositoryDB, domain.GetRolePermissions()),
	}

	router := gin.Default()

	router.POST("/login", authHandlers.Login)
	router.GET("/verify", authHandlers.Verify)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)

	if err != nil {
		log.Fatal("Can not start server")
	}

}
