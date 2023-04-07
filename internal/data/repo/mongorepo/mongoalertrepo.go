package mongorepo

import (
	"com/anoop/examples/internal/data/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
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
	_, err := r.coll.InsertOne(context.Background(), ent)
	if err != nil {
		return &entity.Alerts{}, err
	}
	//ent.Id = res.InsertedID
	return &ent, nil
}

func (r *MongoAlertRepository) Get(id string) (*entity.Alerts, error) {
	result := &entity.Alerts{}
	err := r.coll.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *MongoAlertRepository) ByDeviceId(id string) (*[]entity.Alerts, error) {
	results := &[]entity.Alerts{}
	err := r.coll.FindOne(context.TODO(), bson.D{{"DeviceId", id}}).Decode(results)
	if err != nil {
		return results, err
	}
	return results, nil
}

func (r *MongoAlertRepository) Delete(id string) error {
	result := r.coll.FindOneAndDelete(context.TODO(), bson.D{{"DeviceId", id}})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
