package DAL

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbContext struct {
	MongoContext   *context.Context
	Users          *mongo.Collection
	Companies      *mongo.Collection
	ControlSystems *mongo.Collection
	Devices        *mongo.Collection
	DeviceData     *mongo.Collection
}

func NewDbContext() *DbContext {
	clOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	cl, err := mongo.Connect(context.TODO(), clOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = cl.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return &DbContext{
		Users:          cl.Database("test").Collection("users"),
		Companies:      cl.Database("test").Collection("companies"),
		ControlSystems: cl.Database("test").Collection("controlSystems"),
		Devices:        cl.Database("test").Collection("devices"),
		DeviceData:     cl.Database("test").Collection("deviceData"),
		MongoContext:   &ctx,
	}
}
