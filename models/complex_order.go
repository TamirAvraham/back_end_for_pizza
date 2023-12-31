package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ComplexOrderData struct {
	Pizzas     map[primitive.ObjectID]int `json:"pizzas,omitempty" validate:"required"`
	Address    string                     `json:"address,omitempty" validate:"required"`
	Phone      string                     `json:"phone,omitempty" validate:"required"`
	TotalPrice float64                    `json:"total_price,omitempty" validate:"required"`
}
type ComplexOrder struct {
	Id   primitive.ObjectID `json:"id,omitempty"`
	Data ComplexOrderData   `json:"data,omitempty" validate:"required"`
}
