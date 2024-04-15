package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileStoreImpl struct {
	database *mongo.Database
}

func (f *FileStoreImpl) CreateFile(dto *CreateFileDTO) error {
	dto.Id = primitive.NewObjectID()

	_, err := f.database.Collection("files").InsertOne(context.TODO(), dto)
	return err
}
