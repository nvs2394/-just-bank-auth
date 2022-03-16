package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvs2394/just-bank-auth/common"
	"github.com/nvs2394/just-bank-auth/dto"
	"github.com/nvs2394/just-bank-auth/service"
	"github.com/nvs2394/just-bank-lib/logger"
)

type AuthHandlers struct {
	service service.AuthService
}

func (authHandler *AuthHandlers) Login(context *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := context.BindJSON(&loginRequest); err != nil {
		logger.Error("Error while decoding login request" + err.Error())
		context.JSON(http.StatusBadRequest, nil)
	} else {
		loginResponse, err := authHandler.service.Login(loginRequest)

		if err != nil {
			context.JSON(http.StatusUnauthorized, err.AsMessage())
		} else {
			context.JSON(http.StatusOK, loginResponse)
		}
	}
}

/*
  Sample URL string
 http://localhost:8181/auth/verify?token=somevalidtokenstring&routeName=GetCustomer&customer_id=2000&account_id=95470
*/

func (authHandler *AuthHandlers) Verify(context *gin.Context) {
	urlParams := make(map[string]string)

	queries := []string{"token", "routeName", "customer_id", "account_id"}

	for _, key := range queries {
		urlParams[key] = context.Query(key)
	}

	if urlParams["token"] != "" {
		err := authHandler.service.Verify(urlParams)
		if err != nil {
			context.JSON(err.Code, common.NotAuthorizedResponse(err.Message))
		} else {
			context.JSON(http.StatusOK, common.AuthorizedResponse())
		}

	} else {
		context.JSON(http.StatusForbidden, common.NotAuthorizedResponse("Missing token"))
	}
}
