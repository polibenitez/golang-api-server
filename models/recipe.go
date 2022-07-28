package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty" validate:"required,alpha"`
	Description string             `json:"description" bson:"description,omitempty" validate:"required,alpha"`
	Time        string             `json:"time" bson:"time,omitempty" validate:"required,alpha"`
	Complexity  string             `json:"complexity" bson:"complexity,omitempty" validate:"required,alpha"`
	Ingredients []Ingredients      `json:"ingredients" bson:"ingredients,omitempty" validate:"required,alpha"`
	Steps       []string           `json:"steps" bson:"steps,omitempty" validate:"required,alpha"`
	Cousine     string             `json:"cousine" bson:"cousine,omitempty" validate:"required,alpha"`
	Healthy     string             `json:"healthy" bson:"healthy,omitempty" validate:"required,alpha"`
}
