package store

import (
	"context"

	"go-mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectStore struct {
	db *mongo.Collection
}

func NewProjectStore(db *mongo.Collection) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

func (us *ProjectStore) Create(u *model.Project) error {
	_, err := us.db.InsertOne(context.TODO(), u)
	return err
}

func (us *ProjectStore) Remove(field, value string) error {
	_, err := us.db.DeleteOne(context.TODO(), bson.M{field: value})
	return err
}

func (us *ProjectStore) GetByName(name string) (*model.Project, error) {
	var u model.Project
	err := us.db.FindOne(context.TODO(), bson.M{"name": name}).Decode(&u)
	return &u, err
}
