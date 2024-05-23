package datasource

import (
	"context"
	"gitlab.com/bat.ggl/bgDigital/internal/core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	mgClient *mongo.Client
}

func (u *UsersRepo) Login(login, password string) (*entity.User, error) {
	user := &entity.User{}
	filter := bson.M{
		"login":    login,
		"password": password,
	}
	err := u.mgClient.Database(dbName).Collection(users, nil).FindOne(
		context.Background(),
		filter,
		nil,
	).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUsersRepo(client *mongo.Client) *UsersRepo {
	return &UsersRepo{mgClient: client}
}

func (u *UsersRepo) Register(email, password string) (*entity.User, error) {
	return nil, nil
}

func (u *UsersRepo) FindUsers() {

}

func (u *UsersRepo) BlockUsers() {

}

//
// // создаем нового рекрутера
// func (u *UsersRepo) CreateRecruiter(recruiter entity.Recruiter) (err error) {
// 	opts := options.Update().SetUpsert(true)
// 	_, err = u.mgClient.Database(dbName).
// 		Collection(recruiterTable, nil).
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
