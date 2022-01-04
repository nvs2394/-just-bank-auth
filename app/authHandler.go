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

func (authHandler *AuthHandlers) login(response http.ResponseWriter, r *http.Request) {
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
