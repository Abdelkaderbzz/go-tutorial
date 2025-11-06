package models

import "go.mongodb.org/mongo-driver/bson/primitive"
type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name      string             `bson:"name" json:"name"`
    Email     string             `bson:"email" json:"email"`
    Password  string             `bson:"password,omitempty" json:"password,omitempty"`
    Age       int                `bson:"age,omitempty" json:"age,omitempty"`
    CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}