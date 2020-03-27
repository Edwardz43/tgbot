package mongodb_test

import (
	"Edwardz43/tgbot/config"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestPing(t *testing.T) {
	connStr := config.GetMongoConnStr()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	err := client.Ping(ctx, nil)
	assert.Nil(t, err)
}
