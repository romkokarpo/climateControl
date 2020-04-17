package repositories

import (
	"climateControl/DAL"
	"climateControl/models"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	dbContext *DAL.DbContext
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		dbContext: DAL.NewDbContext(),
	}
}

func (repository *UserRepository) GetAllUsers() []*models.User {
	ctx := repository.dbContext

	var users []*models.User
	cursor, err := ctx.Users.Find(*ctx.MongoContext, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(*ctx.MongoContext, &users); err != nil {
		panic(err)
	}

	return users
}
