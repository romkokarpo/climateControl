package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	FirstName        string             `bson:"firstName,omitempty"`
	LastName         string             `bson:"lastName,omitempty"`
	Email            string             `bson:"email,omitempty"`
	PasswordSalt     string             `bson:"passwordSalt,omitempty"`
	PasswordHash     string             `bson:"passwordHash,omitempty"`
	LastLoginDate    time.Time          `bson:"lastLoginDate,omitempty"`
	RegistrationDate time.Time          `bson:"registrationDate,omitempty"`
}
