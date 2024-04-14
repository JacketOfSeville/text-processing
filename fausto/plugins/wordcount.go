package plugins

import (
	"encoding/json"
	"fmt"

	"github.com/Gustrb/text-processing/fausto/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WordCountPlugin struct{}

// Return:  textId, wordCount
type CountReturnDTO struct {
	Id        primitive.ObjectID `json:"textId" bson:"textId"`
	WordCount int                `json:"wordCount" bson:"wordCount"`
}

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
	rawJson := CountReturnDTO{
		Id:        input.Id,
		WordCount: count,
	}

	// Marshal the struct into JSON
	jsonResult, err := json.Marshal(rawJson)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return
	}

	// Convert JSON data to a string
	result := string(jsonResult)

	// TODO: Do something with the generated JSON
	fmt.Println(result)
}
