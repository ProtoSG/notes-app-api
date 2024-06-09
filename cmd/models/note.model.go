package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title      string             `bson:"title" json:"title"`
	Content    string             `bson:"content" json:"content"`
	Categories []string           `bson:"categories" json:"categories"`
	SharedWith []string           `bson:"sharedWith" json:"sharedWith"`
	CreatedAt  primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt  primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
