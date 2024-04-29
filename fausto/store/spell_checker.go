package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type SpellCheckerStoreImpl struct {
	database *mongo.Database
}

func (s *SpellCheckerStoreImpl) CreateSpellCheckerMetadata(dto *CreateSpellCheckerMetaDTO) error {
	_, err := s.database.Collection("spell_checker").InsertOne(context.TODO(), dto)
	return err
}
