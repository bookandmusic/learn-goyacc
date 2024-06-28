package calc

func NewCalcParser(text string) Result {
	l := newLexer(text)
	yyParse(l)
	return l.Result
}
