package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `json:"product_id,omitempty" bson:"_id,omitempty"`
	ProductName string             `json:"product_name" bson:"product_name"`
	Stock       int64              `json:"stock" bson:"stock"`
}
