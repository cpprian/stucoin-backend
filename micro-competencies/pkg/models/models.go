package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MicroCompetence struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Owner       string             `bson:"owner"`
	Supervisor  string             `bson:"supervisor"`
	Subject     string             `bson:"subject"`
	Files       []string           `bson:"files"`
}
