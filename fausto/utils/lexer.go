package utils

import (
	"unicode"
)

const (
	TokWord         = iota
	TokPunctuations = iota
	TokEOF          = iota
)

type Token struct {
	Type   int
	Value  string
	Line   int
	Column int
}

type Lexer struct {
	text string
	pos  int
	line int
	col  int
}

func NewLexer(text string) *Lexer {
	return &Lexer{text: text, pos: 0, line: 0, col: 0}
}

func (l *Lexer) advance() {
	if l.pos >= len(l.text) {
		return
	}

	if l.text[l.pos] == '\n' {
		l.line++
		l.col = 0
	}
	if l.text[l.pos] == '\t' {
		l.col += 4
	} else {
		l.col++
	}

	l.pos++
}

func (l *Lexer) NextToken() Token {
	// Skip whitespaces
	for l.pos < len(l.text) && unicode.IsSpace(rune(l.text[l.pos])) {
		l.advance()
	}

	if l.pos >= len(l.text) {
		return Token{Type: TokEOF, Value: "", Line: l.line, Column: l.col}
	}

	r := rune(l.text[l.pos])
	if unicode.IsLetter(r) {
		return l.consumeWord()
	}

	l.pos++

	return Token{Type: TokPunctuations, Value: string(r), Line: l.line, Column: l.col}
}

func (l *Lexer) consumeWord() Token {
	start := l.pos
	for l.pos < len(l.text) && unicode.IsLetter(rune(l.text[l.pos])) {
		l.advance()
	}

	return Token{Type: TokWord, Value: l.text[start:l.pos], Line: l.line, Column: l.col}
}
