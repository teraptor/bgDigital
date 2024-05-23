package http_server

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var (
	jwtSecretKey = "cybuthc"
)

/*
структура для регистрации нового пользователя
*/
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Container) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := RegisterRequest{}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err) // todo - обработать корректно ошибку
	}
	_, err = c.UsersService.Register(data.Email, data.Password)
	if err != nil {
		ErrorResponse(w, []byte(err.Error()), http.StatusInternalServerError)
	}
	Response(w, []byte("a"))
}

/*
create jwt token
*/
func (c *Container) Login(w http.ResponseWriter, r *http.Request) {
	resp, err := c.UsersService.Login("1", "1")
	if err != nil {
		ErrorResponse(w, []byte(err.Error()), http.StatusInternalServerError)
	}

	payload := jwt.MapClaims{
		"id":                  resp.ID,
		"login":               resp.Login,
		"emailConfirmed":      resp.EmailConfirmed,
		"parentOrganithation": resp.ParentOrganithation,
		"role":                resp.Role,
		"status":              resp.Status,
		"licenceExpiredAt":    resp.LicenceExpiredAt,
		"exp":                 time.Now().Add(time.Minute * 5),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		panic(err)
	}

	type ResponseJSON struct {
		Token string `json:"token"`
	}
	object := ResponseJSON{
		Token: t,
	}

	data, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(data)
	if err != nil {
		return
	}
}

/*
logout and erase token
*/
func (c *Container) Logout(w http.ResponseWriter, r *http.Request) {
	user := LogoutResponse{
		Status: true,
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	Response(w, data)
}
