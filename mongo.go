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

//TODO
func (c *Conn) query() error {

	return nil
}

// setup initializes mongodb client connection
func (c *Conn) setup(client *mongo.Client, db, collection string) {
	c.Client = client
	c.Collection = c.Client.Database(db).Collection(collection)
}

// insertMany executes batch insert option and returns error if any
func (c *Conn) insertMany(docs []interface{}) error {
	ctx, f := context.WithTimeout(context.Background(), 5*time.Second)
	defer f()

	_, err := c.Collection.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	return nil
}

// insertOne inserts a single document and returns the data ID if success
func (c *Conn) insertOne(doc interface{}) (interface{}, error) {
	ctx, f := context.WithTimeout(context.Background(), 5*time.Minute)
	defer f()

	r, err := c.Collection.InsertOne(ctx, doc)

	if err != nil {
		return "", err
	}

	return r.InsertedID, nil
}

// update updates the data and returns the data ID if success
func (c *Conn) update(filter, data interface{}) (interface{}, error) {
	ctx, f := context.WithTimeout(context.Background(), 5*time.Minute)
	defer f()

	r, err := c.Collection.UpdateOne(ctx, filter, bson.M{
		"$set": data,
	})

	if err != nil {
		return "", err
	}

	return r.UpsertedID, nil
}

// upsert update the data or insert one if it does not exist
func (c *Conn) upsert(filter, data interface{}) (interface{}, error) {
	ctx, f := context.WithTimeout(context.Background(), 5*time.Minute)
	defer f()

	r, err := c.Collection.UpdateOne(ctx, filter, bson.M{
		"$set": data,
	}, options.Update().SetUpsert(true))

	if err != nil {
		return "", err
	}

	return r.UpsertedID, nil
}

// updateMany executes batch update option and returns the number of rows affected
func (c *Conn) updateMany(filter, data interface{}) (interface{}, error) {
	ctx, f := context.WithTimeout(context.Background(), 5*time.Minute)
	defer f()

	r, err := c.Collection.UpdateOne(ctx, filter, bson.M{
		"$set": data,
	})

	if err != nil {
		return "", err
	}

	return r.UpsertedCount, nil
}

// contains returns true if the document exist
func (c *Conn) contains(doc interface{}) (bool, error) {
	cFind, cancelFind := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFind()

	cur, err := c.Collection.Find(cFind, doc)
	if err != nil {
		return false, err
	}

	cCur, cancelCur := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCur()

	defer cur.Close(cCur)

	cNext, cancelNext := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelNext()

	return cur.Next(cNext), nil
}
