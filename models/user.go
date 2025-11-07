package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name      string             `bson:"name" binding:"required,min=3,max=20" validate:"required,min=3,max=20" json:"name"`
    Email     string             `bson:"email" binding:"required,email" validate:"required,email" json:"email"`
    Password  string             `bson:"password,omitempty" validate:"required" json:"password,omitempty"`
    Age       int                `bson:"age,omitempty" validate:"gte=0,lte=120" json:"age,omitempty"`
    CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}