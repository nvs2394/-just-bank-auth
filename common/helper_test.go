package common

import (
	"testing"
)

func TestNotAuthorizedResponse(t *testing.T) {
	errorMessage := "Missing token"
	result := NotAuthorizedResponse(errorMessage)
	var expected = map[string]interface{}{
		"isAuthorized": false,
		"message":      errorMessage,
	}
	if result["message"] != expected["message"] {
		t.Error("result should be Missing token, got", result["message"])
	}

	if result["isAuthorized"] != expected["isAuthorized"] {
		t.Error("result should be false, got", result["isAuthorized"])
	}
}
