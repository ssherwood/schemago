package main

import (
	"flag"
	"fmt"
	"schemago/pkg/schemago"
)

func main() {
	fmt.Println("Schema Go!")

	numTables := flag.Int("tables", 10, "Number of tables to generate")
	maxColumns := flag.Int("maxcols", 10, "Maximum number of columns to generate per table")
	flag.Parse()

	schema := schemago.CreateSchema(*numTables, *maxColumns)
	fmt.Println(schemago.GenerateSQLStatements(schema))
}
