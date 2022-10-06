package models

type Toping struct {
	Name     string  `json:"name,omitempty" validate:"required"`
	Amount   uint8   `json:"amount,omitempty" validate:"required"`
	Location float32 `json:"location,omitempty" validate:"required"`
	Price    float32 `json:"price,omitempty" validate:"required"`
}
