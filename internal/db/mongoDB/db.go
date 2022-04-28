package mongoDB

import (
	"clean/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Database struct {
	Db *mongo.Client
}

func NewDatabase() *Database {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil
	}

	return &Database{
		Db: client,
	}
}

func (db *Database) GetUserDataByName(username string) *models.User {
	var result *models.User
	collection := db.Db.Database("udsf").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"name": username}).Decode(&result)
	if err != nil {
		return nil
	} else {
		return result
	}
}

func (db *Database) GetUsers() []*models.User {
	var result []*models.User
	collection := db.Db.Database("udsf").Collection("users")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil
	}
	for cur.Next(context.TODO()) {
		var k *models.User
		_ = cur.Decode(&k)
		result = append(result, k)
	}
	return result
}
