package sql

import (
	"strings"
	"testing"
)

func TestQLParse(t *testing.T) {
	tests := []struct {
		text    string
		expexts []Stmt
	}{
		{
			"dbscale show slow sql top 10",
			[]Stmt{
				{
					Text:   "dbscale show slow sql top 10",
					Type:   STMT_DBSCALE_REQUEST_SLOW_SQL_TOP_N,
					Params: []string{"10"},
				},
			},
		},
		{
			"dbscale show audit user list",
			[]Stmt{
				{
					Text:   "dbscale show audit user list",
					Type:   STMT_DBSCALE_SHOW_AUDIT_USER_LIST,
					Params: []string{},
				},
			},
		},
		{
			"dbscale request node info;dbscale request cluster id;dbscale request cluster info;",
			[]Stmt{
				{
					Text:   "dbscale request node info",
					Type:   STMT_DBSCALE_REQUEST_NODE_INFO,
					Params: []string{},
				},
				{
					Text:   "dbscale request cluster id",
					Type:   STMT_DBSCALE_REQUEST_CLUSTER_ID,
					Params: []string{},
				},
				{
					Text:   "dbscale request cluster info",
					Type:   STMT_DBSCALE_REQUEST_CLUSTER_INFO,
					Params: []string{},
				},
			},
		},
	}

	for _, test := range tests {
		results, _ := NewSqlParser(test.text)
		for i, result := range results {
			if result.Type != test.expexts[i].Type || strings.ToLower(result.Command()) != test.expexts[i].Text {
				t.Errorf("Parse SQL: %s,Expect SQL: %s, Parse Type: %d, Expect Type: %d", result.Command(),test.expexts[i].Text, result.Type, test.expexts[i].Type)
			}
		}
	}

}
