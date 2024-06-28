package main

import (
	"fmt"
	"os"

	"example.com/goyacc/calc"
	"example.com/goyacc/sql"
)

func testSQL(sql_str string) {
	results, err := sql.NewSqlParser(sql_str)
	if err == nil {
		for _, result := range results {
			fmt.Printf("Parsed Command: Type: %d, Command: %s, Params: %v\n", result.Type, result.Command(), result.Params)
		}
	} else {
		fmt.Printf("%v", err)
	}
}

func testCalc(expr string) {
	r := calc.NewCalcParser(expr)
	if r.Error == nil {
		fmt.Println("Result:", r.Num)
	} else {
		fmt.Print(r.Error.Error())
		os.Exit(1)
	}
}

func main() {
	testCalc("2*(1.6+(2.7-0.2))")
	testSQL("dbscale request node info;dbscale request cluster id;dbscale request cluster info;")
}
