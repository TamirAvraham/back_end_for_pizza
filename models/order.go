package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Pizza Pizza              `json:"pizza,omitempty" validate:"required"`
}
