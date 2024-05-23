package tests

import (
	"github.com/go-chi/chi/v5"
	"gitlab.com/bat.ggl/bgDigital/internal/appplication"
	"gitlab.com/bat.ggl/bgDigital/pkg/http_server"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Server struct {
	Router *chi.Mux
}

/*
*
создаем новый сервер
*/
func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	c := appplication.NewContainer()
	http_server.MountainHandlers(s.Router, c)

	go func() {
		err := http.ListenAndServe(":3000", s.Router)
		if err != nil {
			panic(err.Error())
		}
	}()
	return s
}

/*
*
выполнение запроса
*/
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

/*
*
проверка кода от сервера
*/
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
