package plugins

import (
	"encoding/json"
	"fmt"
)

type WordCountPlugin struct{}

// Return:  textId, wordCount
type CountReturnDTO struct {
	Id        string `json:"textId" bson:"textId"`
	WordCount int    `json:"wordCount" bson:"wordCount"`
}

// Word counting function
func countWord(data string) int {
	var tot int

	for _, char := range data {
		switch char {
		case ' ', '\t', '\n':
			tot++
		}
	}

	return tot
}

// Recieve: id, text, createdAt
func (*WordCountPlugin) Execute(input PluginInputData) {
	rawData := input.Content

	// Create a map to store JSON data
	var data map[string]interface{}

	// Decode JSON into a map
	err := json.Unmarshal([]byte(rawData), &data)
	if err != nil {
		fmt.Println("Erro decoding JSON: ", err)
		return
	}

	// Assign map key values to varaibles
	id := data["_id"].(string)
	text := data["text"].(string)

	// Count the text's words
	count := countWord(text)

	// Struct the response
	rawJson := CountReturnDTO{
		Id:        id,
		WordCount: count,
	}

	// Marshal the struct into JSON
	jsonResult, err := json.Marshal(rawJson)
	if err != nil {
		fmt.Println("Erro ao criar JSON:", err)
		return
	}

	// Conver JSON data to a string
	result := string(jsonResult)

	// TODO: Do something with the generated JSON
	fmt.Println(result)
}
