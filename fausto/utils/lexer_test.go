package utils_test

import (
	"testing"

	"github.com/Gustrb/text-processing/fausto/utils"
)

func TestLexerOnEmptyString(t *testing.T) {
	lexer := utils.NewLexer("")
	token := lexer.NextToken()

	if token.Type != utils.TokEOF {
		t.Errorf("Expected EOF token, got %d", token.Type)
	}
}

func TestLexerOnSingleWord(t *testing.T) {
	lexer := utils.NewLexer("hello")
	token := lexer.NextToken()

	if token.Type != utils.TokWord {
		t.Errorf("Expected Word token, got %d", token.Type)
	}

	if token.Value != "hello" {
		t.Errorf("Expected 'hello', got %s", token.Value)
	}
}

func TestLexerAPhrase(t *testing.T) {
	lexer := utils.NewLexer("Hello, world!")
	tokens := []utils.Token{}

	for {
		token := lexer.NextToken()
		tokens = append(tokens, token)

		if token.Type == utils.TokEOF {
			break
		}
	}

	if len(tokens) != 5 {
		t.Errorf("Expected 4 tokens, got %d", len(tokens))
	}

	if tokens[0].Type != utils.TokWord || tokens[0].Value != "Hello" {
		t.Errorf("Expected 'Hello', got %s", tokens[0].Value)
	}

	if tokens[1].Type != utils.TokPunctuations || tokens[1].Value != "," {
		t.Errorf("Expected ',', got %s", tokens[1].Value)
	}

	if tokens[2].Type != utils.TokWord || tokens[2].Value != "world" {
		t.Errorf("Expected 'world', got %s", tokens[2].Value)
	}

	if tokens[3].Type != utils.TokPunctuations || tokens[3].Value != "!" {
		t.Errorf("Expected '!', got %s", tokens[3].Value)
	}
}

func TestItCanCountTheWordsOfASentence(t *testing.T) {
	sentence := "Hello, this is a big sentence, and there are many words in it!"
	lexer := utils.NewLexer(sentence)

	words := 0
	var token utils.Token
	for token.Type != utils.TokEOF {
		token = lexer.NextToken()
		if token.Type == utils.TokWord {
			words++
		}
	}

	if words != 13 {
		t.Errorf("Expected 13 words, got %d", words)
	}
}

func TestItCanCountLinesAndColumns(t *testing.T) {
	sentence := "hello\n\tworld\n\t\t!"
	lexer := utils.NewLexer(sentence)

	tokens := []utils.Token{}
	for {
		token := lexer.NextToken()
		tokens = append(tokens, token)

		if token.Type == utils.TokEOF {
			break
		}
	}

	if tokens[0].Line != 0 || tokens[0].Column != 5 {
		t.Errorf("Expected line 0, column 0, got %d, %d", tokens[0].Line, tokens[0].Column)
	}

	if tokens[1].Line != 1 || tokens[1].Column != 10 {
		t.Errorf("Expected line 1, column 0, got %d, %d", tokens[1].Line, tokens[1].Column)
	}

	if tokens[2].Line != 2 || tokens[2].Column != 9 {
		t.Errorf("Expected line 2, column 1, got %d, %d", tokens[2].Line, tokens[2].Column)
	}
}
