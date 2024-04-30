package plugins

import (
	"bufio"
	"log"
	"os"

	"github.com/Gustrb/text-processing/fausto/store"
	"github.com/Gustrb/text-processing/fausto/utils"
)

const (
	DictionaryPath = "fixtures/dictionary.txt"
)

type SpellChecker struct {
	dictionary map[string]bool
}

func SpellCheckerNew() SpellChecker {
	return SpellChecker{dictionary: make(map[string]bool)}
}

func (s SpellChecker) Init() {
	log.Println("Initializing SpellChecker plugin")

	file, err := os.Open(DictionaryPath)
	if err != nil {
		log.Printf("Error loading SpellChecker plugin: %v", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s.dictionary[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning SpellChecker plugin: %v", err)
		return
	}

	log.Println("SpellChecker plugin initialized")
}

func (s SpellChecker) Execute(data PluginInputData) {
	log.Printf("Running SpellChecker plugin on file: %s\n", data.Id.Hex())

	spellCheckingMetadata := store.CreateSpellCheckerMetaDTO{TextID: data.Id}
	spellCheckingMetadata.SpellingErrors = utils.FindOccurencesOf(data.Content, func(word string) bool {
		return !s.dictionary[word]
	})

	if len(spellCheckingMetadata.SpellingErrors) > 0 {
		if err := store.GetStore().SpellCheckerStore().CreateSpellCheckerMetadata(&spellCheckingMetadata); err != nil {
			log.Printf("Error saving SpellChecker metadata: %v", err)
		}
	}

	log.Println("SpellChecker plugin executed")
}
