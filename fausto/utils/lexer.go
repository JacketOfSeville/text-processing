package utils

import "unicode"

const (
	TokWord         = iota
	TokPunctuations = iota
	TokEOF          = iota
)

type Token struct {
	Type  int
	Value string
}

type Lexer struct {
	text string
	pos  int
}

func NewLexer(text string) *Lexer {
	return &Lexer{text: text, pos: 0}
}

func (l *Lexer) NextToken() Token {
	// Skip whitespaces
	for l.pos < len(l.text) && unicode.IsSpace(rune(l.text[l.pos])) {
		l.pos++
	}

	if l.pos >= len(l.text) {
		return Token{Type: TokEOF, Value: ""}
	}

	r := rune(l.text[l.pos])
	if unicode.IsLetter(r) {
		return l.consumeWord()
	}

	l.pos++

	return Token{Type: TokPunctuations, Value: string(r)}
}

func (l *Lexer) consumeWord() Token {
	start := l.pos
	for l.pos < len(l.text) && unicode.IsLetter(rune(l.text[l.pos])) {
		l.pos++
	}

	return Token{Type: TokWord, Value: l.text[start:l.pos]}
}
