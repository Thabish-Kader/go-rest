package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	Id primitive.ObjectID 	`json:"id"`
	Name string 				`json:"name"`
	Number int					`json:"number"`
}