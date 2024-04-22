package plugins_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Gustrb/text-processing/fausto/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
func TestExecute(t *testing.T) {
	content := "This is a test!\tHowever\nIt may not work"
	objID := primitive.NewObjectID()

	count := countWord(content)

	if count != 9 {
		t.Errorf("Expected 9 words, got %d", count)
	}

	rawJson := CountReturnDTO{
		Id:        objID,
		WordCount: count,
	}

	jsonResult, err := json.Marshal(rawJson)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return
	}
	result := string(jsonResult)
	fmt.Println(result)
}
