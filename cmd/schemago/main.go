package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	_ "runtime/pprof"
	"schemago/pkg/schemago"
)

var numberOfTables = flag.Int("tables", 10, "Number of tables to generate")
var maximumNumberOfColumns = flag.Int("columns", 10, "Maximum number of columns to generate per table")
var defaultSchemaName = flag.String("schema", "random", "Schema name to use instead of random")
var enumsEnabled = flag.Bool("enums", true, "Enable or disable generation of native enum types")
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

	schema := schemago.CreateSchema(flagsAsConfig())
	err := schemago.WriteSQLStatements(os.Stdout, schema)
	if err != nil {
		log.Fatal("Schemago failed:", err)
	}
}

func flagsAsConfig() schemago.Config {
	return schemago.Config{
		DefaultSchemaName:      *defaultSchemaName,
		NumberOfTables:         *numberOfTables,
		MaximumNumberOfColumns: *maximumNumberOfColumns,
		EnumsEnabled:           *enumsEnabled,
	}
}
