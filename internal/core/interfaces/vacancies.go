package interfaces

import "gitlab.com/bat.ggl/bgDigital/internal/core/entity"

type VacanciesRepository interface {
	CreateVacancies(e entity.Vacancies) error
}
