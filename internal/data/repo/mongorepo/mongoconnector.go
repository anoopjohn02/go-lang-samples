package mongorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	client *mongo.Client
}

func NewMongoRepo(uri string) *MongoRepo {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &MongoRepo{client: client}
}

func (m *MongoRepo) Db() *mongo.Database {
	return m.client.Database("admin")
}

func (m *MongoRepo) Disconnect() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
