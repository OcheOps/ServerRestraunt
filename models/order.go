package models 


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order is a struct that represents a single order.
type Order struct {

	// ID is the unique identifier for the order.
		ID         	primitive.ObjectID 	`bson:"_id"`
		Dish       	*string            	`json:"dish"`
		Price   	*float64           	`json:"price""`
		Server		*string			`json:"server"`
		Table		*string			`json:"table"`
	}