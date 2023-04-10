package mongorepo

import (
	"com/anoop/examples/internal/data/entity"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAlertRepository struct {
	coll *mongo.Collection
}

func NewAlertRepository(db *mongo.Database) *MongoAlertRepository {
	coll := db.Collection("alerts")
	return &MongoAlertRepository{coll: coll}
}

func (r *MongoAlertRepository) Save(ent entity.Alerts) (*entity.Alerts, error) {
	ent.ID = primitive.NewObjectID()
	_, err := r.coll.InsertOne(context.Background(), ent)
	if err != nil {
		return &entity.Alerts{}, err
	}
	return &ent, nil
}

func (r *MongoAlertRepository) Get(id string) (*entity.Alerts, error) {
	result := &entity.Alerts{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := r.coll.FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *MongoAlertRepository) ByDeviceId(id string) (*[]entity.Alerts, error) {
	results := []entity.Alerts{}
	cursor, err := r.coll.Find(context.TODO(), bson.D{{"deviceid", id}})
	if err != nil {
		return &results, err
	}
	for cursor.Next(context.TODO()) {
		var result entity.Alerts
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return &results, nil
}

func (r *MongoAlertRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result := r.coll.FindOneAndDelete(context.TODO(), bson.D{{"_id", objID}})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
