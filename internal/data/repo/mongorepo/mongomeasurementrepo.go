package mongorepo

import (
	"com/anoop/examples/internal/data/entity"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMeasurementRepository struct {
	coll *mongo.Collection
}

func NewMeasurementRepository(db *mongo.Database) *MongoMeasurementRepository {
	coll := db.Collection("measurements")
	return &MongoMeasurementRepository{coll: coll}
}

func (r *MongoMeasurementRepository) Save(ent entity.Measurements) (*entity.Measurements, error) {
	ent.ID = primitive.NewObjectID()
	_, err := r.coll.InsertOne(context.Background(), ent)
	if err != nil {
		return &entity.Measurements{}, err
	}
	return &ent, nil
}

func (r *MongoMeasurementRepository) Get(id string) (*entity.Measurements, error) {
	result := &entity.Measurements{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := r.coll.FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *MongoMeasurementRepository) ByDeviceId(id string) (*[]entity.Measurements, error) {
	results := []entity.Measurements{}
	cursor, err := r.coll.Find(context.TODO(), bson.D{{"deviceid", id}})
	if err != nil {
		return &results, err
	}
	for cursor.Next(context.TODO()) {
		var result entity.Measurements
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

func (r *MongoMeasurementRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result := r.coll.FindOneAndDelete(context.TODO(), bson.D{{"_id", objID}})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
