package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB 创建连接 MongoDB 的函数
func MongoDB() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://47.113.198.106:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	collection := client.Database("jackaroo").Collection("job")
	return collection, nil
}
