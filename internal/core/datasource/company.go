package datasource

import (
	"go.mongodb.org/mongo-driver/mongo"
)

/**
репозиторий в которох храним компании с заключеным договором о сотрудничистве
*/

type CompanyRepository struct {
	mgClient *mongo.Client
}

func CreateCompanyRepo(client *mongo.Client) *CompanyRepository {
	return &CompanyRepository{
		mgClient: client,
	}
}

//
// // создаем новую компанию
// func (c *CompanyRepository) CreateCompany() (uuid, error) {
// 	opts := options.Update().SetUpsert(true)
// 	_, err = c.mgClient.Database(dbName).
// 		Collection(companyCollection, nil).
// 		UpdateOne(
// 			context.Background(),
// 			bson.D{{"chatID", recruiter.ChatID}},
// 			bson.D{
// 				{"$set",
// 					bson.D{
// 						{"active", false},
// 						{"name", time.Now().Unix()},
// 						{"email", recruiter.Email},
// 						{"phone", recruiter.Phone},
// 						{"created_at", recruiter.CreatedAt},
// 						{"deleted_at", recruiter.DeletedAt},
// 						{"chatID", recruiter.ChatID},
// 						{"subscribeExpired", nil},
// 					},
// 				},
// 			},
// 			opts)
// 	if err != nil {
// 		panic(any(err))
// 	}
// 	return nil
// }
//
// // создаем новую компанию
// func (c *CompanyRepository) UpdateCompany() {
//
// }
