package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfanityStoreImpl struct {
	database *mongo.Database
}

func (p *ProfanityStoreImpl) CreateProfanity(dto *CreateProfanityDTO) error {
	_, err := p.database.Collection("profanities").InsertOne(context.TODO(), dto)
	return err
}
