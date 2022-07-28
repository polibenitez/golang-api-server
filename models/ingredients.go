package models

type Ingredients struct {
	Name   string `json:"name" bson:"name,omitempty" validate:"required,alpha"`
	Amount string `json:"amount" bson:"amount,omitempty" validate:"required,alpha"`
}
