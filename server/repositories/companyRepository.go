package repositories

import (
	"github.com/romkokarpo/climateControl/DAL"
	"github.com/romkokarpo/climateControl/DTO"
	"github.com/romkokarpo/climateControl/models"

	"github.com/PeteProgrammer/go-automapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyRepository struct {
	dbContext *DAL.DbContext
}

func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{
		dbContext: DAL.NewDbContext(),
	}
}

func (repository *CompanyRepository) GetAllCompanies() []*models.Company {
	ctx := repository.dbContext
	var result []*models.Company

	cursor, err := ctx.Companies.Find(*ctx.MongoContext, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(*ctx.MongoContext, &result); err != nil {
		panic(err)
	}

	return result
}

func (repository *CompanyRepository) GetCompanyById(id string) *models.Company {
	ctx := repository.dbContext
	hexedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	queryResult := ctx.Users.FindOne(*ctx.MongoContext, bson.M{"_id": hexedId})
	result := models.Company{}
	queryResult.Decode(result)

	return &result
}

func (repository *CompanyRepository) AddNewOffices(officesDto *DTO.EmbeddedOfficeDto) bool {
	ctx := repository.dbContext
	var offices []*models.EmbeddedOffice

	automapper.MapLoose(officesDto.Offices, &offices)

	insertResult, err := ctx.Companies.UpdateOne(
		*ctx.MongoContext,
		bson.M{"_id": officesDto.CompanyID},
		bson.M{
			"$push": bson.M{
				"offices": bson.M{
					"$each": offices,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	return insertResult.ModifiedCount > 0
}

func (repository *CompanyRepository) AddNewGreenhouses(greenhouses *DTO.EmbeddedGreenhouseDto) bool {
	ctx := repository.dbContext
	var greenhouseModels []*models.EmbeddedOffice

	automapper.MapLoose(greenhouses.Greenhouses, &greenhouseModels)

	insertResult, err := ctx.Companies.UpdateOne(
		*ctx.MongoContext,
		bson.M{"_id": greenhouses.CompanyID},
		bson.M{
			"$push": bson.M{
				"greenhouses": bson.M{
					"$each": greenhouseModels,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	return insertResult.ModifiedCount > 0
}
