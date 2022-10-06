package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pizza struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Price    float32            `json:"price,omitempty" validate:"required"`
	PngUrl   string             `json:"png_path,omitempty"`
	Toppings []Toping           `json:"toppings,omitempty" validate:"required"`
	Size     int                `json:"size,omitempty" validate:"required"`
}
