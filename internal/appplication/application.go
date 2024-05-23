package appplication

import (
	"gitlab.com/bat.ggl/bgDigital/internal/core/datasource"
	"gitlab.com/bat.ggl/bgDigital/internal/core/service/users"
	"gitlab.com/bat.ggl/bgDigital/internal/core/service/vacancies"
	"gitlab.com/bat.ggl/bgDigital/internal/shedulled"
	"gitlab.com/bat.ggl/bgDigital/pkg/http_server"
	"gitlab.com/bat.ggl/bgDigital/pkg/mgo"
	"sync"
)

func Run() {
	c := http_server.NewContainer()
	client := mgo.Connect("mongodb://root:example@localhost:27017")

	// репозиторий и сервис для работы с вакансиями
	vacRepo := datasource.CreateVacanciesRepo(client)
	c.VacService = vacancies.CreateVacanciesService(vacRepo)

	// репозиторий и сервис для работы с пользователями
	userRepo := datasource.CreateUsersRepo(client)
	c.UsersService = users.CreateUserService(userRepo)

	// run http server
	go func() {
		s := http_server.CreateNewServer(c)
		s.Run()
	}()

	// shedulled task
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// /  прокинуть контекст
		tgShedulled := shedulled.CreateTgParser(c.VacService)
		tgShedulled.ParseChannels()
	}()
	wg.Wait()
}
