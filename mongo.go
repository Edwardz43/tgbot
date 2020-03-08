package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conn represents a MongoDB client connection
type Conn struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (c *Conn) query() error {

	return nil
}

func (c *Conn) setup(client *mongo.Client, db, collection string) {
	c.Client = client
	c.Collection = c.Client.Database(db).Collection(collection)
}

func (c *Conn) insertMany(docs []interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := c.Collection.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	return nil
}

func (c *Conn) insertOne(doc interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	r, err := c.Collection.InsertOne(ctx, doc)

	if err != nil {
		return "", err
	}

	return r.InsertedID, nil
}

func (c *Conn) update(filter, data interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	r, err := c.Collection.UpdateOne(ctx, filter, bson.M{
		"$set": data,
	})

	if err != nil {
		return "", err
	}

	return r.UpsertedID, nil
}

func (c *Conn) upsert(filter, data interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	r, err := c.Collection.UpdateOne(ctx, filter, bson.M{
		"$set": data,
	}, options.Update().SetUpsert(true))

	if err != nil {
		return "", err
	}

	return r.UpsertedID, nil
}

func (c *Conn) updateMany(filter, data interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	r, err := c.Collection.UpdateOne(ctx, filter, bson.M{
		"$set": data,
	})

	if err != nil {
		return "", err
	}

	return r.UpsertedCount, nil
}

func (c *Conn) contains(doc interface{}) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := c.Collection.Find(ctx, doc)
	if err != nil {
		return false, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cur.Close(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	return cur.Next(ctx), nil
}
