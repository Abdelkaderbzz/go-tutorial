package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"myapp/database"
	"myapp/routes"
	"os"
)

func main() {
	database.ConnectMongo()
	defer database.CloseMongo()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	routes.RegisterUserRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
