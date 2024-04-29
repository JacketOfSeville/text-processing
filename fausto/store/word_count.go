package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WordCountStoreImpl struct {
	database *mongo.Database
}

func (wc *WordCountStoreImpl) CreateWordCountMetadata(dto *CreateWordCountDTO) error {
	dto.Id = primitive.NewObjectID()

	_, err := wc.database.Collection("word_count_metadata").InsertOne(context.TODO(), dto)
	return err
}
