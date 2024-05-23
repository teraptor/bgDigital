package shedulled

import (
	"fmt"
	"gitlab.com/bat.ggl/bgDigital/internal/core/entity"
	"gitlab.com/bat.ggl/bgDigital/internal/core/service/vacancies"
	"gitlab.com/bat.ggl/bgDigital/pkg/tg_a_c_i_su"
	"time"
)

const (
	repeatebleTask time.Duration = 10
)

type TelegramParser struct {
	vacService *vacancies.VacanciesService
}

func CreateTgParser(vacService *vacancies.VacanciesService) *TelegramParser {
	return &TelegramParser{
		vacService: vacService,
	}
}

func (t *TelegramParser) ParseChannels() bool {
	ticker := time.NewTicker(repeatebleTask * time.Minute)
	go func() {
		for {
			select {
			// case <-r.gf.Done():
			// 	return
			case _ = <-ticker.C:
				client := tg_a_c_i_su.CreateACIClient()
				data, err := client.Run("skillstaff_platform")
				if err != nil {
					_ = fmt.Errorf("%s", "http code more 200... delay later") // todo
					continue
				}
				for _, v := range data.Messages {
					err := t.vacService.CreateVacancies(
						entity.Vacancies{
							MessageID: v.PostID,
							Message:   v.Message,
							CreatedAt: time.Now(),
						},
					)
					fmt.Sprint("done")
					if err != nil {
						_ = fmt.Errorf("%s", "http code more 200... delay later") // todo
						continue
					}
				}
			}
		}
	}()

	return true
}
