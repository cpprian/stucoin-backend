package mongodb

import (
	"context"
	"errors"

	"github.com/cpprian/stucoin-backend/rewards/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RewardModel struct {
	C *mongo.Collection
}

func (m *RewardModel) All() ([]models.Reward, error) {
	ctx := context.TODO()
	var rewards []models.Reward

	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &rewards); err != nil {
		return nil, err
	}

	return rewards, nil
}

// FindByID finds a reward by id
func (m *RewardModel) FindById(id string) (*models.Reward, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	var reward models.Reward

	if err := m.C.FindOne(ctx, bson.M{"_id": p}).Decode(&reward); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no reward found")
		}
		return nil, err
	}

	return &reward, nil
}

// FindByRewardname finds a reward by rewardname
func (m *RewardModel) FindByName(name string) (*models.Reward, error) {
	ctx := context.TODO()
	var reward models.Reward

	if err := m.C.FindOne(ctx, bson.M{"name": name}).Decode(&reward); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no reward found")
		}
		return nil, err
	}

	return &reward, nil
}

// InsertReward inserts a new reward to the database
func (m *RewardModel) InsertReward(reward *models.Reward) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()

	res, err := m.C.InsertOne(ctx, reward)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateReward updates the reward with the given id when posts, comments or subscribes are added/removed/updated
func (m *RewardModel) UpdateReward(reward *models.Reward) (*mongo.UpdateResult, error) {
	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": reward.ID}, bson.M{"$set": reward})
	if err != nil {
		return nil, err
	}

	return res, nil
}
