package mongodb

import (
	"context"
	"errors"

	"github.com/cpprian/stucoin-backend/tasks/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskModel struct {
	C *mongo.Collection
}

func (m *TaskModel) All() ([]models.Task, error) {
	ctx := context.TODO()
	var tasks []models.Task

	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// FindByID finds a task by id
func (m *TaskModel) FindById(id string) (*models.Task, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	var task models.Task

	if err := m.C.FindOne(ctx, bson.M{"_id": p}).Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no task found")
		}
		return nil, err
	}

	return &task, nil
}

// FindByTaskname finds a task by taskname
func (m *TaskModel) FindByTitle(title string) (*models.Task, error) {
	ctx := context.TODO()
	var task models.Task

	if err := m.C.FindOne(ctx, bson.M{"title": title}).Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no task found")
		}
		return nil, err
	}

	return &task, nil
}

// InsertTask inserts a new task to the database
func (m *TaskModel) InsertTask(task *models.Task) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()

	res, err := m.C.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateTask updates the task with the given id when posts, comments or subscribes are added/removed/updated
func (m *TaskModel) UpdateTask(task *models.Task) (*mongo.UpdateResult, error) {
	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": task.ID}, bson.M{"$set": task})
	if err != nil {
		return nil, err
	}

	return res, nil
}
