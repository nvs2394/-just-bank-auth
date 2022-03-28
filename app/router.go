package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nvs2394/just-bank-auth/common"
	"github.com/nvs2394/just-bank-auth/domain"
	"github.com/nvs2394/just-bank-auth/service"
)

func NewRouter() *gin.Engine {
	dbClient := common.GetDBClient()
	authRepositoryDB := domain.NewAuthRepositoryDb(dbClient)

	authHandlers := AuthHandlers{
		service: service.NewAuthService(authRepositoryDB, domain.GetRolePermissions()),
	}

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.POST("/login", authHandlers.Login)
	v1.GET("/verify", authHandlers.Verify)

	return router
}
