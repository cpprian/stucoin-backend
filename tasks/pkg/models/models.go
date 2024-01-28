package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EnumCompleted string

const (
	Open 	  	EnumCompleted = "OPEN"
	Completed   EnumCompleted = "COMLETED"
	Incompleted EnumCompleted = "INCOMPLETED"
	Aborted     EnumCompleted = "ABORTED"
	Accepted    EnumCompleted = "ACCEPTED"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CoverImage  string             `bson:"cover_image"`
	Points      int                `bson:"points"`
	Completed   EnumCompleted      `bson:"completed"`
	Owner       string             `bson:"owner"`
	InCharge    string             `bson:"in_charge"`
	Files       []File             `bson:"files"`
}

type CoverImage struct {
	CoverImage string `bson:"cover_image"`
}

type Content struct {
	Content string `bson:"content"`
}

type Title struct {
	Title string `bson:"title"`
}

type File struct {
	Name string `bson:"file"`
	Path string `bson:"path"`
	Size int64  `bson:"size"`
}

type Files []File

type InCharge struct {
	InCharge string `bson:"in_charge"`
}

type Points struct {
	Points int `bson:"points"`
}