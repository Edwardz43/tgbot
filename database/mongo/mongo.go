package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo represents a MongoDB client
type Mongo struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (m *Mongo) query() error {

	return nil
}

func (m *Mongo) Connect(connStr string) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	Client, _ := mongo.Connect(ctx, options.Client().ApplyURI(connStr))

	m.Client = Client

	m.Collection = m.Client.Database("Test").Collection("tgbot")
}

func (m *Mongo) InsertMany(docs []interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := m.Collection.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) InsertOne(doc interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	r, err := m.Collection.InsertOne(ctx, doc)

	if err != nil {
		return "", err
	}

	return r.InsertedID, nil
}

func (m *Mongo) Update(filter, data interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	r, err := m.Collection.UpdateOne(ctx, filter, bson.M{
		"$set": data,
	})

	if err != nil {
		return "", err
	}

	return r.UpsertedID, nil
}

func (m *Mongo) Upsert(column, data interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	f := bson.M{
		column.(string): data,
	}

	r, err := m.Collection.UpdateOne(ctx, f, bson.M{
		"$set": f,
	}, options.Update().SetUpsert(true))

	if err != nil {
		return "", err
	}

	return r.UpsertedID, nil
}

func (m *Mongo) UpdateMany(filter, data interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	r, err := m.Collection.UpdateOne(ctx, filter, bson.M{
		"$set": data,
	})

	if err != nil {
		return "", err
	}

	return r.UpsertedCount, nil
}

func (m *Mongo) Contains(doc interface{}) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := m.Collection.Find(ctx, doc)
	if err != nil {
		return false, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cur.Close(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	return cur.Next(ctx), nil
}
