package plugins

import (
	"log"

	"github.com/Gustrb/text-processing/fausto/store"
	"github.com/Gustrb/text-processing/fausto/utils"
)

type WordCountPlugin struct{}

func WordCountNew() WordCountPlugin {
	return WordCountPlugin{}
}

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

func (w WordCountPlugin) Init() {
	log.Println("WordCount plugin initialized")
}

func (w WordCountPlugin) Execute(input PluginInputData) {
	log.Printf("WordCount plugin begining execution on content: %s", input.Content)
	count := countWord(input.Content)

	wordCountMetadata := store.CreateWordCountDTO{Id: input.Id}
	wordCountMetadata.WordCount = count

	err := store.GetStore().WordCountStore().CreateWordCountMetadata(&wordCountMetadata)
	if err != nil {
		log.Printf("Error while saving word counting metadata: %v", err)
	}
}
