package calc

import (
	"fmt"
	"strconv"
	"unicode"
)

type Result struct {
	Num   float64
	Error error
	State int
}

type lexer struct {
	text   string
	pos    int
	Result Result
}

func newLexer(text string) *lexer {
	return &lexer{
		text: text,
	}
}

func (l *lexer) Lex(lval *yySymType) int {
	if l.Result.Error != nil {
		return 0
	}
	for l.pos < len(l.text) {
		ch := l.text[l.pos]
		l.pos++

		switch {
		case unicode.IsDigit(rune(ch)):
			start := l.pos - 1
			for l.pos < len(l.text) && (unicode.IsDigit(rune(l.text[l.pos])) || l.text[l.pos] == '.') {
				l.pos++
			}
			numStr := l.text[start:l.pos]
			if num, err := strconv.ParseFloat(numStr, 64); err != nil {
				l.Error(fmt.Sprintf("invalid number: %s", numStr))
				return 0
			} else {
				lval.num = num
			}
			return NUMBER
		case unicode.IsSpace(rune(ch)):
			continue
		default:
			return int(ch)
		}
	}

	return 0
}

func (l *lexer) Error(s string) {
	if s != "" && l.Result.Error == nil {
		l.Result.Error = fmt.Errorf(s)
	}
}
