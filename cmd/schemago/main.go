package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	_ "runtime/pprof"
	"schemago/pkg/schemago"
)

var numTables = flag.Int("tables", 10, "Number of tables to generate")
var maxColumns = flag.Int("columns", 10, "Maximum number of columns to generate per table")
var schemaName = flag.String("schema", "random", "Schema name to use instead of random")
var cpuProfileFile = flag.String("cpuprofile", "", "Write CPU profile to the specified file")

func main() {
	flag.Parse()

	if *cpuProfileFile != "" {
		cpuFile, err := os.Create(*cpuProfileFile)
		if err != nil {
			log.Fatal(err)
		}
		defer cpuFile.Close()

		err = pprof.StartCPUProfile(cpuFile)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	schema := schemago.CreateSchema(*schemaName, *numTables, *maxColumns)
	err := schemago.WriteSQLStatements(os.Stdout, schema)
	if err != nil {
		log.Fatal("Schemago failed:", err)
	}
}
