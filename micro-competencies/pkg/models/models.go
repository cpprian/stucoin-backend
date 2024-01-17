package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EnumCompleted string

const (
	Completed   EnumCompleted = "COMLETED"
	Incompleted EnumCompleted = "INCOMPLETED"
	Aborted     EnumCompleted = "ABORTED"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Points      int                `bson:"points"`
	Completed   EnumCompleted      `bson:"completed"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Owner       string             `bson:"owner"`
	InCharge    string             `bson:"in_charge"`
	Files       []string           `bson:"files"`
	Images      []string           `bson:"images"`
	Tags        []string           `bson:"tags"`
}