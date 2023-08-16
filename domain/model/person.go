package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Address   *Address           `json:"address" bson:"address"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
	Telephone string             `json:"telephone,omitempty" bson:"telephone"`
	Cellphone string             `json:"cellphone,omitempty" bson:"cellphone"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password"`
	Access    []AccessModule     `json:"access" bson:"access"`
	IsMaster  bool               `json:"is_master" bson:"is_master"`
}

type PersonJWT struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	IsMaster bool               `json:"is_master" bson:"is_master"`
}
