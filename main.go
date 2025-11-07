package main

import (
	"log"
	"myapp/database"
	"myapp/routes"
	"os"
	_ "myapp/docs" 
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Example User API
// @version 1.0
// @description This is a sample server for users.
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	database.ConnectMongo()
	defer database.CloseMongo()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterUserRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
