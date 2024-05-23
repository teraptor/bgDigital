package datasource

import (
	"context"
	"gitlab.com/bat.ggl/bgDigital/internal/core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
вакансии которые надо взять в работу...
*/

type VacanciesRepository struct {
	db *mongo.Client
}

func (v VacanciesRepository) CreateVacancies(e entity.Vacancies) error {
	opts := options.Update().SetUpsert(true)
	_, err := v.db.Database(dbName).
		Collection(vacancies, nil).
		UpdateOne(
			context.Background(),
			bson.D{
				{"postID", e.MessageID},
			},
			bson.D{
				{"$set",
					bson.D{
						{"message", e.Message},
						{"created_at", e.CreatedAt},
						{"deleted_at", nil},
					},
				},
			},
			opts)
	if err != nil {
		panic(any(err))
	}

	return nil
}

func CreateVacanciesRepo(db *mongo.Client) *VacanciesRepository {
	return &VacanciesRepository{
		db: db,
	}
}
