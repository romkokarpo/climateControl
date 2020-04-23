package repositories

import (
	"climateControl/server/DAL"
	"climateControl/server/DTO"
	"climateControl/server/customErrors"
	"climateControl/server/models"
	"log"
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
	queryResult.Decode(&user)

	return &user
}

func (repository *UserRepository) GetUserByEmail(email string) *models.User {
	ctx := repository.dbContext
	queryResult := ctx.Users.FindOne(*ctx.MongoContext, bson.M{"email": email})
	user := models.User{}
	queryResult.Decode(&user)

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
	insertResult, err := ctx.Users.InsertOne(*ctx.MongoContext, user)
	if err != nil {
		panic(err)
	}

	return insertResult
}

func (repository *UserRepository) CheckUserCredentials(email string, password string) bool {
	user := repository.GetUserByEmail(email)

	if err := compareHashAndPassword(user.PasswordHash, password); err == nil {
		ctx := repository.dbContext
		updateResult, err := ctx.Users.UpdateOne(*ctx.MongoContext,
			bson.M{"_id": user.ID},
			bson.D{
				{"$set", bson.D{{"lastLoginDate", time.Now()}}},
			},
		)
		if err != nil {
			panic(err)
		} else if updateResult.ModifiedCount != 0 {
			log.Println("Successful update")
		}

		return true
	} else {
		err := customErrors.UserPasswordIncorrectError{}.New()
		panic(err)
	}
}

func (repository *UserRepository) UpdateUserEmail(currentEmail string, newEmail string) bool {
	ctx := repository.dbContext
	updateResult, err := ctx.Users.UpdateOne(
		*ctx.MongoContext,
		bson.M{"email": currentEmail},
		bson.D{
			{"$set", bson.D{{"email", newEmail}}},
		},
	)

	if err != nil {
		panic(err)
	}
	if updateResult.ModifiedCount > 0 {
		return true
	}

	return false
}

func (repository *UserRepository) UpdateUserPassword(email string, currentPassword string, newPassword string) bool {
	ctx := repository.dbContext
	newPasswordHash, err := hashPassword(newPassword)
	if err != nil {
		panic(err)
	}

	if repository.CheckUserCredentials(email, currentPassword) {
		updateResult, err := ctx.Users.UpdateOne(
			*ctx.MongoContext,
			bson.M{"email": email},
			bson.D{
				{"$set", bson.D{{"passwordHash", newPasswordHash}}},
			},
		)
		if err != nil {
			panic(err)
		}
		if updateResult.ModifiedCount > 0 {
			return true
		}

		return false
	} else {
		err := customErrors.UserPasswordIncorrectError{}.New()
		panic(err)
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	return string(bytes), err
}

func compareHashAndPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}
