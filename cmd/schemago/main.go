package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	_ "runtime/pprof"
	"schemago/pkg/schemago"
)

func main() {
	flag.Parse()

	if *cpuProfileFile != "" {
		cpuFile, err := os.Create(*cpuProfileFile)
		if err != nil {
			log.Panic(err)
		}
		defer func() {
			if err = cpuFile.Close(); err != nil {
				log.Panic(err)
			}
		}()

		if err = pprof.StartCPUProfile(cpuFile); err != nil {
			log.Panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	schema := schemago.CreateSchema(flagsAsConfig())
	err := schemago.WriteSQLStatements(os.Stdout, schema)
	if err != nil {
		log.Fatal("Schemago failed:", err)
	}
}

var defaultSqlType = flag.String("sqltype", "yugabytedb-ysql", "Type of SQL to generate")
var defaultSchemaName = flag.String("schema", "random", "Schema name to use instead of random")
var numberOfTables = flag.Int("tables", 10, "Number of tables to generate")
var maximumNumberOfColumns = flag.Int("columns", 10, "Maximum number of columns to generate per table")
var enumsEnabled = flag.Bool("enums", true, "Enable or disable generation of native enum types")
var tabletSplitsEnabled = flag.Bool("splits", true, "Enable or disable tablet splitting in table definitions")
var cpuProfileFile = flag.String("cpuprofile", "", "Write CPU profile to the specified file")

func flagsAsConfig() schemago.Config {
	return schemago.Config{
		SQLType:                *defaultSqlType,
		DefaultSchemaName:      *defaultSchemaName,
		NumberOfTables:         *numberOfTables,
		MaximumNumberOfColumns: *maximumNumberOfColumns,
		EnumsEnabled:           *enumsEnabled,
		TabletSplitsEnabled:    *tabletSplitsEnabled,
	}
}
