package main

import (
	"context"
	"fmt"
	"log"
	"myapp/database"
	"myapp/routes"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	client := database.ConnectMongo()
	defer database.CloseMongo()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	db := client.Database("godatabase")
	usersCollection := db.Collection("users")

	user := map[string]interface{}{
		"name": "Kader",
		"role": "developer",
	}

	insertResult, err := usersCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Inserted document with ID:", insertResult.InsertedID)
	router := gin.Default()
	routes.RegisterUserRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
