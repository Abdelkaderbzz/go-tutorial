package controllers

import (
	"context"
	"myapp/database"
	"myapp/models"
	"myapp/services"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

// GreetUser godoc
// @Summary Greet user by name
// @Description Returns a simple greeting message.
// @Tags Users
// @Produce  json
// @Param   name  path  string  true  "User name"
// @Success 200 {object} map[string]string
// @Router /user/{name} [get]
func GreetUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
}
// CreateUser godoc
// @Summary Create a new user
// @Description Create a user with name, email, and age
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   user  body  models.User  true  "User data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "Validation error"
// @Failure 409 {object} map[string]string "Email already registered"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	if database.MongoClient == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database not initialized"})
		return
	}

	usersCollection := database.MongoClient.Database("godatabase").Collection("users")
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 	if err := validate.Struct(user); err != nil {
        errors := make(map[string]string)
        for _, err := range err.(validator.ValidationErrors) {
            errors[err.Field()] = err.Tag() // e.g., "required", "email", "min"
        }
        c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errors})
        return
    }
	var existingUser models.User
	err := usersCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}
	if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	svc := services.NewUserService(database.MongoClient, "godatabase")
	insertedID, err := svc.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"inserted_id": insertedID,
		"success":     true,
		"message":     "user created successfully",
	})
}