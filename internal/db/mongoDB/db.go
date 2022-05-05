package mongoDB

import (
	"clean/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

type Database struct {
	Db  *mongo.Client
	log *zap.Logger
}

func NewDatabase(log *zap.Logger) *Database {
	log.Info("GetUserDataByName db called")
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
		Db:  client,
		log: log,
	}
}

func (db *Database) GetUserDataByName(username string) *models.User {
	db.log.Info("GetUserDataByNamed db called")
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
	db.log.Info("GetUsers db called")
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

func (db *Database) GetLogger() *zap.Logger {
	return db.log
}
