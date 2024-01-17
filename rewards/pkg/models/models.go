package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reward struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Points      int                `bson:"points"`
	OrgName     string             `bson:"orgName"`
	OrgLocation string             `bson:"orgLocation"`
	Image       string             `bson:"image"`
}
