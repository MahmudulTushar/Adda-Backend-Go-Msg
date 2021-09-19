package connection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type mongoDatabase struct {
	Client *mongo.Client
}

func NewMongoDatabaseInstance() *mongoDatabase {
	mongoDb := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoDb)
	clientOptions = clientOptions.SetMaxPoolSize(50)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to Db!")
	return &mongoDatabase{
		Client: dbClient,
	}
}

var MongoDBInstance = NewMongoDatabaseInstance()
