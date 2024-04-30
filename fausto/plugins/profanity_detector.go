package plugins

import (
	"bufio"
	"log"
	"os"

	"github.com/Gustrb/text-processing/fausto/store"
	"github.com/Gustrb/text-processing/fausto/utils"
)

const (
	ProfanityListPath = "fixtures/bad-words.txt"
)

type ProfanityDetector struct {
	badWords map[string]bool
}

func ProfanityDetectorNew() Plugin {
	return ProfanityDetector{badWords: make(map[string]bool)}
}

func (p ProfanityDetector) Init() {
	file, err := os.Open(ProfanityListPath)

	if err != nil {
		log.Printf("Error loading ProfanityDetctor plugin: %v", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// Using a map to store the bad words for O(1) lookup
		// otherwise we would have to iterate over the list of bad words
		// for each word in the text (sigh)
		p.badWords[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning ProfanityDetector plugin: %v", err)
		return
	}

	log.Printf("ProfanityDetector plugin initialized with %d bad words", len(p.badWords))
}

func (p ProfanityDetector) Execute(data PluginInputData) {
	log.Printf("ProfanityDetector plugin executing on file: %s\n", data.Id.Hex())

	badWordsMetadata := store.CreateProfanityDTO{TextID: data.Id}
	badWordsMetadata.Profanities = utils.FindOccurencesOf(data.Content, func(word string) bool {
		return p.badWords[word]
	})

	if len(badWordsMetadata.Profanities) > 0 {
		err := store.GetStore().ProfanityStore().CreateProfanity(&badWordsMetadata)
		if err != nil {
			log.Printf("Error saving profanity metadata: %v", err)
		}
	}
}
