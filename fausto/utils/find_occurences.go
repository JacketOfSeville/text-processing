package utils

type Location struct {
	Line int `json:"line"`
	Col  int `json:"column"`
}

type OccurenceInText struct {
	Line   int      `json:"line"`
	Column int      `json:"column"`
	End    Location `json:"end"`
	Word   string   `json:"word"`
}

func FindOccurencesOf(content string, callback func(string) bool) []OccurenceInText {
	lexer := NewLexer(content)
	occurences := []OccurenceInText{}
	var token Token
	for token.Type != TokEOF {
		token = lexer.NextToken()

		if token.Type == TokWord && callback(token.Value) {
			occurences = append(occurences, OccurenceInText{
				Line:   token.Line,
				Column: token.Column - len(token.Value),
				End: Location{
					Line: token.Line,
					Col:  token.Column,
				},
				Word: token.Value,
			})
		}
	}

	return occurences
}
