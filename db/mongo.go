package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var dbName = "minemind_production"
func New() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", "admin", "MineMindProject2019", "207.148.71.222", "27017")
	client, _ := mongo.NewClient(options.Client().ApplyURI(uri))
	if err := client.Connect(ctx); err != nil {
		panic(err.Error())
	}
	db := client.Database(dbName)
	return db
}
