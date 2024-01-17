package mongodb

import (
	"context"
	"errors"

	"github.com/cpprian/stucoin-backend/micro-competencies/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MicroCompetenceModel struct {
	C *mongo.Collection
}

func (m *MicroCompetenceModel) All() ([]models.MicroCompetence, error) {
	ctx := context.TODO()
	var microCompetences []models.MicroCompetence

	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &microCompetences); err != nil {
		return nil, err
	}

	return microCompetences, nil
}

// FindByID finds a microCompetence by id
func (m *MicroCompetenceModel) FindById(id string) (*models.MicroCompetence, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	var microCompetence models.MicroCompetence

	if err := m.C.FindOne(ctx, bson.M{"_id": p}).Decode(&microCompetence); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no microCompetence found")
		}
		return nil, err
	}

	return &microCompetence, nil
}

// FindByMicroCompetencename finds a microCompetence by microCompetencename
func (m *MicroCompetenceModel) FindByName(name string) (*models.MicroCompetence, error) {
	ctx := context.TODO()
	var microCompetence models.MicroCompetence

	if err := m.C.FindOne(ctx, bson.M{"name": name}).Decode(&microCompetence); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no microCompetence found")
		}
		return nil, err
	}

	return &microCompetence, nil
}

// InsertMicroCompetence inserts a new microCompetence to the database
func (m *MicroCompetenceModel) InsertMicroCompetence(microCompetence *models.MicroCompetence) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()

	res, err := m.C.InsertOne(ctx, microCompetence)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateMicroCompetence updates the microCompetence with the given id when posts, comments or subscribes are added/removed/updated
func (m *MicroCompetenceModel) UpdateMicroCompetence(microCompetence *models.MicroCompetence) (*mongo.UpdateResult, error) {
	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": microCompetence.ID}, bson.M{"$set": microCompetence})
	if err != nil {
		return nil, err
	}

	return res, nil
}
