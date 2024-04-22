package plugins

import (
	// "context"
	"encoding/json"
	"fmt"

	"github.com/Gustrb/text-processing/fausto/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

type WordCountPlugin struct{}

// Return:  textId, wordCount
type CountReturnDTO struct {
	Id        primitive.ObjectID `json:"textId" bson:"textId"`
	WordCount int                `json:"wordCount" bson:"wordCount"`
}

/* Database and storage related functions
// Primary tests succeeded but did still need
// to find a simpler way to do this

type DataStore interface {
	WordStore() WordStore
}

type WordStore interface {
	CreateWord(*CountReturnDTO) error
}

type WordStoreImpl struct {
	database *mongo.Database
}

func (d *WordStoreImpl) WordStore() WordStore {
	return &WordStoreImpl{database: d.database}
}

func (f *WordStoreImpl) CreateWord(dto *CountReturnDTO) error {
	_, err := f.database.Collection("words").InsertOne(context.TODO(), dto)
	return err
}
*/

// Word counting function
func countWord(data string) int {
	words := 0
	lexer := utils.NewLexer(data)

	var token utils.Token
	for token.Type != utils.TokEOF {
		token = lexer.NextToken()
		if token.Type == utils.TokWord {
			words++
		}
	}

	return words
}

// Recieve: id, text, createdAt
func (*WordCountPlugin) Execute(input PluginInputData) {
	count := countWord(input.Content)
	rawResult := CountReturnDTO{
		Id:        input.Id, // ID equals to the original text's ID in the DB
		WordCount: count,
	}

	// Marshal the struct into JSON
	jsonResult, err := json.Marshal(rawResult)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return
	}
	result := string(jsonResult) // Convert JSON data to a string
	fmt.Println(result)          // TODO: Do something with the generated JSON

	// TODO: Store said result in the database
}
