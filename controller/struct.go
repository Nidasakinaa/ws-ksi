package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Ingredients string             `bson:"ingredients,omitempty" json:"ingredients,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Calories    float64            `bson:"calories,omitempty" json:"calories,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"` 
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
}

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    FullName string             `bson:"fullname,omitempty" json:"fullname,omitempty"`
    Phone    string             `bson:"phonenumber,omitempty" json:"phonenumber,omitempty"`
    Username string             `bson:"username,omitempty" json:"username,omitempty"`
    Password string             `bson:"password,omitempty" json:"password,omitempty"`
    Role     string             `bson:"role,omitempty" json:"role,omitempty"` // "admin" or "customer"
}

type ReqMenuItem struct {
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Ingredients string             `bson:"ingredients,omitempty" json:"ingredients,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Calories    float64            `bson:"calories,omitempty" json:"calories,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"` 
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
}

type ReqUser struct {
    FullName string             `bson:"name,omitempty" json:"name,omitempty"`
    Phone    string             `bson:"phone,omitempty" json:"phone,omitempty"`
    Username string             `bson:"username,omitempty" json:"username,omitempty"`
    Password string             `bson:"password,omitempty" json:"password,omitempty"`
    Role     string             `bson:"role,omitempty" json:"role,omitempty"` // "admin" or "customer"
}