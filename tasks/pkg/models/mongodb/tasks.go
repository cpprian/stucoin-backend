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

func (m *TaskModel) AllTeacherTasks(owner string) ([]models.Task, error) {
	ctx := context.TODO()
	var tasks []models.Task

	cursor, err := m.C.Find(ctx, bson.M{"owner": owner})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

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

func (m *TaskModel) InsertTask(task *models.Task) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()

	res, err := m.C.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) UpdateTask(task *models.Task) (*mongo.UpdateResult, error) {
	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": task.ID}, bson.M{"$set": task})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) DeleteTask(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.DeleteOne(ctx, bson.M{"_id": p})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) UpdateCoverImageById(id string, coverImage models.CoverImage) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": bson.M{"cover_image": coverImage.CoverImage}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) UpdateContentById(id string, content models.Content) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": bson.M{"description": content.Content}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) UpdateTitleById(id string, title models.Title) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": bson.M{"title": title.Title}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) SaveFilesById(id string, file models.File) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$push": bson.M{"files": file}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) DeleteFileById(id string, file models.File) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$pull": bson.M{"files": file}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) AssignTaskById(id string, inCharge models.InCharge) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": bson.M{"in_charge": inCharge.InCharge, "completed": "INCOMPLETED"}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) CompleteTaskById(id string) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": bson.M{"completed": "COMPLETED"}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) AcceptTaskById(id string) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": bson.M{"completed": "ACCEPTED"}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) RejectTaskById(id string) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": bson.M{"completed": "ABORTED"}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TaskModel) UpdatePointsById(id string, points models.Points) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	
	res, err := m.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": bson.M{"points": points.Points}})
	if err != nil {
		return nil, err
	}

	return res, nil
}