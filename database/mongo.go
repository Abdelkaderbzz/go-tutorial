// database/mongo.go
package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client 
func ConnectMongo() *mongo.Client {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := godotenv.Load()
    if err != nil {
        log.Fatal("❌ Error loading .env file")
    }

    clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("❌ Could not connect to MongoDB:", err)
    }

    fmt.Println("✅ Connected to MongoDB!")

    MongoClient = client
    return client
}


func CloseMongo() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if MongoClient != nil {
        if err := MongoClient.Disconnect(ctx); err != nil {
            log.Fatal(err)
        }
        fmt.Println("🔌 Disconnected from MongoDB.")
    }
}