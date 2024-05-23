package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogout(t *testing.T) {
	// Create a New Server with Handlers
	s := CreateNewServer()

	// Create a New Request
	req, _ := http.NewRequest("GET", "/user/login", nil)
	// Execute Request
	response := executeRequest(req, s)

	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)

	// We can use testify/require to assert values, as it is more convenient
	require.Equal(t, "Hello World!", response.Body.String())
}
