package store

import (
	"context"
	"errors"
	"fmt"
	"time"

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

func Disconnect() error {
	if _client == nil {
		return errors.New("cannot disconnect unconnected server")
	}

	return _client.Disconnect(context.TODO())
}

func GetStore() DataStore {
	return _store
}
