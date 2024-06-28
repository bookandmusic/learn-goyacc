package sql

func NewSqlParser(text string) ([]*Stmt, error) {
	l := newLexer(text)
	yyParse(l)
	return l.Results, l.Err
}
