package repositories

import (
	"climateControl/DAL"
	"climateControl/DTO"
	"climateControl/models"
	"time"

	"github.com/PeteProgrammer/go-automapper"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (repository *UserRepository) GetUserById(id string) *models.User {
	ctx := repository.dbContext
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	queryResult := ctx.Users.FindOne(*ctx.MongoContext, bson.M{"_id": userId})
	user := models.User{}
	queryResult.Decode(user)

	return &user
}

func (repository *UserRepository) GetUserByEmail(email string) *models.User {
	ctx := repository.dbContext
	queryResult := ctx.Users.FindOne(*ctx.MongoContext, bson.M{"email": email})
	user := models.User{}
	queryResult.Decode(user)

	return &user
}

func (repository *UserRepository) RegisterUser(model DTO.UserDto) *mongo.InsertOneResult {
	var user models.User
	automapper.MapLoose(model, user)

	hash, err := hashPassword(model.Password)
	if err != nil {
		panic(err)
	}

	user.PasswordHash = hash
	user.RegistrationDate = time.Now()
	user.LastLoginDate = time.Now()

	ctx := repository.dbContext
	insertResult, err := repository.dbContext.Users.InsertOne(*ctx.MongoContext, user)
	if err != nil {
		panic(err)
	}

	return insertResult
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
