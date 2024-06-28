%{
package calc


func setResult(l yyLexer, v float64) {
  l.(*lexer).Result.Num = v
}
%}

%union {
    num float64
}

%token <num> NUMBER 
%type <num> exp
%left '+' '-'
%left '*' '/'
%right '('
%left ')'

%start exp

%%

exp:
    NUMBER           
    { 
        $$ = $1
        setResult(yylex, $$)
    }
    | exp '+' exp    { $$ = $1 + $3 }
    | exp '-' exp    { $$ = $1 - $3 }
    | exp '*' exp    { $$ = $1 * $3 }
    | exp '/' exp    { $$ = $1 / $3 }
    | '(' exp ')'    { $$ = $2 }
    ;
