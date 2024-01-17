package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MicroCompetence struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Owner       string             `bson:"owner"`
	Supervisor  string             `bson:"supervisor"`
	StartDate   time.Time          `bson:"startDate"`
	EndDate     time.Time          `bson:"endDate"`
	Subject     string             `bson:"subject"`
	Tags        []string           `bson:"tags"`
	Resources   []string           `bson:"resources"`
	Files       []string           `bson:"files"`
}
