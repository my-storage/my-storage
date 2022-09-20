package mongodb

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/my-storage/ms-profile/src/service/config"
)

type MongoDB struct {
	Client *mongo.Client
}

var once sync.Once
var instance *MongoDB

func GetInstance() *MongoDB {
	if instance == nil {
		once.Do(func() {
			db, err := createConnection()
			if err != nil {
				log.Fatal(err)
			}

			instance = db
		})
	}

	return instance
}

func createConnection() (*MongoDB, error) {
	config := config.GetInstance()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDbUrl))
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		Client: client,
	}, nil
}
