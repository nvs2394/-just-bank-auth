package app

import (
	"encoding/json"
	"net/http"

	"github.com/nvs2394/just-bank-auth/common"
	"github.com/nvs2394/just-bank-auth/dto"
	"github.com/nvs2394/just-bank-auth/logger"
	"github.com/nvs2394/just-bank-auth/service"
)

type AuthHandlers struct {
	service service.AuthService
}

func (authHandler *AuthHandlers) Login(response http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request" + err.Error())
		common.WriteResponse(response, http.StatusBadRequest, nil)
	} else {
		loginResponse, err := authHandler.service.Login(loginRequest)

		if err != nil {
			common.WriteResponse(response, http.StatusUnauthorized, err.AsMessage())
		} else {
			common.WriteResponse(response, http.StatusOK, loginResponse)
		}
	}

}

/*
  Sample URL string
 http://localhost:8181/auth/verify?token=somevalidtokenstring&routeName=GetCustomer&customer_id=2000&account_id=95470
*/

func (authHandler *AuthHandlers) Verify(response http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	for key := range r.URL.Query() {
		urlParams[key] = r.URL.Query().Get(key)
	}

	if urlParams["token"] != "" {
		err := authHandler.service.Verify(urlParams)
		if err != nil {
			common.WriteResponse(response, err.Code, common.NotAuthorizedResponse(err.Message))
		} else {
			common.WriteResponse(response, http.StatusOK, common.AuthorizedResponse())
		}

	} else {
		common.WriteResponse(response, http.StatusForbidden, common.NotAuthorizedResponse("missing token"))
	}
}
