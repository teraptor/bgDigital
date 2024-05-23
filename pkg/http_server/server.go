package http_server

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gitlab.com/bat.ggl/bgDigital/internal/core/service/users"
	"gitlab.com/bat.ggl/bgDigital/internal/core/service/vacancies"
	"net/http"
)

// интерфейс для тестирования
type IServer interface {
	MountainHandlers()
}

// контейнер с пробросом сущностей в handlers
type Container struct {
	UsersService *users.UserService
	VacService   *vacancies.VacanciesService
}

type Server struct {
	container *Container
}

/*
Server handlers
*/
func NewContainer() *Container {
	return &Container{}
}

func CreateNewServer(c *Container) *Server {
	return &Server{
		container: c,
	}
}

func (s *Server) Run() {
	r := chi.NewRouter()
	r.Use()
	MountainHandlers(r, s.container)

	http.ListenAndServe(":3000", r)
}

func MountainHandlers(r chi.Router, c *Container) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Hello World!"))
		})
	})
	r.Route("/user", func(r chi.Router) {
		r.Post("/register", c.Register)
		r.Post("/login", c.Login) // post
		r.Get("/logout", c.Logout)
		r.Post("/update", c.Logout)
	})
}

func Response(w http.ResponseWriter, data []byte) {
	type success struct {
		Data interface{}
	}
	s := success{
		Data: data,
	}
	d, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(d)
	if err != nil {
		return
	}
}
func ErrorResponse(w http.ResponseWriter, data []byte, code int) {
	type errorResponse struct {
		Code    int
		Message string
	}
	eResp := errorResponse{
		Code:    code,
		Message: string(data),
	}
	d, err := json.Marshal(eResp)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(d)
	if err != nil {
		return
	}
}
