package sql

import (
	"fmt"
	"strings"
)

var keywords = map[string]int{
	"DBSCALE": DBSCALE,
	"SHOW":    SHOW,
	"SLOW":    SLOW,
	"SQL":     SQL,
	"TOP":     TOP,
	"AUDIT":   AUDIT,
	"USER":    USER,
	"LIST":    LIST,
	"REQUEST": REQUEST,
	"CLUSTER": CLUSTER,
	"ID":      ID,
	"NODE":    NODE,
	"INFO":    INFO,
}

type lexer struct {
	input   string
	start   int
	pos     int
	Results []*Stmt
	Err     error
}

func newLexer(input string) *lexer {
	return &lexer{input: input}
}

func (l *lexer) Lex(lval *yySymType) int {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		l.pos++
		switch {
		case isSpace(ch):
			l.ignore()
		case ch == ';':
			l.ignore()
			return int(ch)
		case isAlpha(ch):
			v_start := l.pos - 1
			ret := l.scanIdentifier()
			if ret == -1 {
				return l.reverseScanKeyValue(v_start, lval)
			}
			return ret
		case isNumber(ch):
			v_start := l.pos - 1
			l.scanNumber()
			return l.reverseScanKeyValue(v_start, lval)

		default:
			return int(ch)
		}
	}
	return 0
}

func (l *lexer) Error(e string) {
	if e != "" {
		l.Err = fmt.Errorf(e)
	}
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) scanNumber() int {
	for isDigit(l.peek()) {
		l.pos++
	}
	return 0
}

func (l *lexer) scanIdentifier() int {
	for isAlpha(l.peek()) {
		l.pos++
	}
	word := l.input[l.start:l.pos]
	if token, ok := keywords[strings.ToUpper(word)]; ok {
		return token
	}
	return -1
}

func (l *lexer) reverseScanKeyValue(v_start int, lval *yySymType) int {
	v_end := l.pos
	l.pos = v_start - 1
	for l.pos >= 0 && (isEqual(l.peek()) || isSpace(l.peek())) {
		l.pos--
	}
	k_end := l.pos + 1
	for l.pos >= 0 && isAlpha(l.peek()) {
		l.pos--
	}
	key := l.input[l.pos+1 : k_end]
	l.pos = v_end
	switch strings.ToUpper(key) {
	case "TOP":
		lval.ident = l.input[v_start:v_end]
		return INTNUM
	}
	return -1
}

func (l *lexer) peek() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch == '_')
}

func isEqual(ch byte) bool {
	return ch == '='
}

func isNumber(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
