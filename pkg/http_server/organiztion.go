package http_server

import (
	"encoding/json"
	"net/http"
	"time"
)

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	user := LoginUserResponse{
		ID:                  "123",
		Login:               "123",
		EmailConfirmed:      false,
		ParentOrganithation: "123",
		Role:                1,
		Status:              1,
		LicenceExpiredAt:    time.Now(),
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
