package main

import "gitlab.com/bat.ggl/bgDigital/internal/appplication"

func main() {
	// создаем репозитории для сохранения настроек и клиентов
	// usersRepo, err := data
	appplication.Run()

	// сделать gracefull shutdown
	shutdown()
}

func shutdown() {

}
