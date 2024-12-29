package api_test

import (
	"testing"

	"frontendmasters.com/go/crypto/api"
)

func TestAPICall(t *testing.T) {
	// Passing empty string to the GetRate func
	_, err := api.GetRate("")

	// If no error, test should fail
	if err == nil {
		t.Error(("Error was not found"))
	}
}
