%{
package sql

%}

%union {
	stmt *Stmt
    ident string
    item interface{}
}

%token <ident>
    DBSCALE 
    SHOW
    SLOW 
    SQL 
    TOP 
    INTNUM
    AUDIT 
    USER 
    LIST
    REQUEST
    CLUSTER 
    ID
    NODE 
    INFO


%type <item>
    StatementList

%type <stmt>
    Statement
    DbscaleShowStmt
    DbscaleRequestStmt
    DbscaleShowParam
    DbscaleRequestParam


%start	Start

%%

Start:
	StatementList

StatementList:
	Statement
	{
		if $1 != nil {
            yylex.(*lexer).Results = append(yylex.(*lexer).Results, $1)
		}
	}
|	StatementList ';'  {}
|	StatementList ';' Statement
	{
		if $3 != nil {
            yylex.(*lexer).Results = append(yylex.(*lexer).Results, $3)
		}
	}

Statement:
    DbscaleShowStmt {$$=$1}
    | 
    DbscaleRequestStmt {$$=$1}
    ;

DbscaleShowStmt:
    DBSCALE SHOW DbscaleShowParam { $$ = $3}

DbscaleShowParam:
    SLOW SQL TOP INTNUM 
    {
        $$ = &Stmt{
            Type: STMT_DBSCALE_REQUEST_SLOW_SQL_TOP_N,
            Params: []string{$4},
        }
    }
    | AUDIT USER LIST 
    {
        $$ = &Stmt{
            Type: STMT_DBSCALE_SHOW_AUDIT_USER_LIST,
        }
    }
    ;

DbscaleRequestStmt: 
    DBSCALE REQUEST DbscaleRequestParam { $$ = $3}

DbscaleRequestParam:
    CLUSTER ID
    {
        $$ = &Stmt{
            Type: STMT_DBSCALE_REQUEST_CLUSTER_ID,
        }
    }
    |
    NODE INFO
    {
        $$ = &Stmt{
            Type: STMT_DBSCALE_REQUEST_NODE_INFO,
        }
    }
    |
    CLUSTER INFO
    {
        $$ = &Stmt{
            Type: STMT_DBSCALE_REQUEST_CLUSTER_INFO,
        }
    }
%%
