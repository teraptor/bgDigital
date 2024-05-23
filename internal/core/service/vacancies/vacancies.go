package vacancies

import (
	"gitlab.com/bat.ggl/bgDigital/internal/core/entity"
	"gitlab.com/bat.ggl/bgDigital/internal/core/interfaces"
)

type VacanciesService struct {
	vacRepo interfaces.VacanciesRepository
}

func CreateVacanciesService(vacRepo interfaces.VacanciesRepository) *VacanciesService {
	return &VacanciesService{
		vacRepo: vacRepo,
	}
}

func (v *VacanciesService) CreateVacancies(e entity.Vacancies) error {
	err := v.vacRepo.CreateVacancies(e)
	if err != nil {
		return err
	}

	return nil // todo
}
