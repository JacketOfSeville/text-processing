package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Gustrb/text-processing/fausto/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _store DataStore
var _client *mongo.Client

const (
	ClientDefaultTimeout = 30 * time.Second
)

type DataStore interface {
	FileStore() FileStore
	ProfanityStore() ProfanityStore
	SpellCheckerStore() SpellCheckerStore
	WordCountStore() WordCountStore
}

type DataStoreImpl struct {
	database *mongo.Database
}

type CreateFileDTO struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Content string             `json:"content" bson:"content"`
	Name    string             `json:"name" bson:"name"`
}

type FileStore interface {
	CreateFile(*CreateFileDTO) error
}

type Location struct {
	Line int `json:"line"`
	Col  int `json:"column"`
}

type ProfanityInText struct {
	Line   int      `json:"line"`
	Column int      `json:"column"`
	End    Location `json:"end"`
	Word   string   `json:"word"`
}

type CreateProfanityDTO struct {
	TextID      primitive.ObjectID      `json:"text_id"`
	Profanities []utils.OccurenceInText `json:"profanities"`
}

type ProfanityStore interface {
	CreateProfanity(*CreateProfanityDTO) error
}

type CreateSpellCheckerMetaDTO struct {
	TextID         primitive.ObjectID      `json:"text_id"`
	SpellingErrors []utils.OccurenceInText `json:"spelling_errors"`
}

type SpellCheckerStore interface {
	CreateSpellCheckerMetadata(*CreateSpellCheckerMetaDTO) error
}

type CreateWordCountDTO struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	WordCount int                `json:"wordCount" bson:"WordCount"`
}

type WordCountStore interface {
	CreateWordCountMetadata(*CreateWordCountDTO) error
}

func Initialize(uri string) error {
	ctx := context.Background()

	clientOpts := options.Client()
	clientOpts.ApplyURI(uri)
	clientOpts.SetConnectTimeout(ClientDefaultTimeout)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		return fmt.Errorf("mongodrv: could not connect: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	_client = client

	_store = &DataStoreImpl{database: _client.Database("fausto")}

	return nil
}

func (d *DataStoreImpl) FileStore() FileStore {
	return &FileStoreImpl{database: d.database}
}

func (d *DataStoreImpl) WordCountStore() WordCountStore {
	return &WordCountStoreImpl{database: d.database}
}

func (d *DataStoreImpl) ProfanityStore() ProfanityStore {
	return &ProfanityStoreImpl{database: d.database}
}

func (d *DataStoreImpl) SpellCheckerStore() SpellCheckerStore {
	return &SpellCheckerStoreImpl{database: d.database}
}

func Disconnect() error {
	if _client == nil {
		return errors.New("cannot disconnect unconnected server")
	}

	return _client.Disconnect(context.TODO())
}

func GetStore() DataStore {
	return _store
}
