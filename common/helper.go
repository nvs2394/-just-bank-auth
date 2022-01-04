package common

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(response http.ResponseWriter, code int, data interface{}) {
	response.Header().Add("Content-type", "application/json")
	response.WriteHeader(code)
	if err := json.NewEncoder(response).Encode(data); err != nil {
		panic(err)
	}
}
