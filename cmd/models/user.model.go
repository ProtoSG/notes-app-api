package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName string             `bson:"userName" json:"userName"`
	Password string             `bson:"password" json:"password"`
}
