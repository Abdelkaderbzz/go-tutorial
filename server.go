package main

import (
	"context"
	"fmt"
	"log"
	"myapp/database"
	"net/http"
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
	// Access database and collection
	db := client.Database("godatabase")
	usersCollection := db.Collection("users")

	// Insert new document
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
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println("Received name:", name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})
	router.POST("/data", func(c *gin.Context) {
		var body map[string]string
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"received": body})
	})
	router.Run(":" + os.Getenv("PORT"))
}
