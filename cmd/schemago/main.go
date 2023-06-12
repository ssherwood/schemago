package main

import (
	"flag"
	"log"
	_ "net/http/pprof"
	"os"
	_ "runtime/pprof"
	"schemago/pkg/schemago"
)

func main() {
	numTables := flag.Int("tables", 10, "Number of tables to generate")
	maxColumns := flag.Int("columns", 10, "Maximum number of columns to generate per table")
	schemaName := flag.String("schema", "random", "Schema name to use instead of random")
	flag.Parse()

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()
	//
	//// Trigger CPU profiling
	//cpuProfileFile, err := os.Create("cpu.prof")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer cpuProfileFile.Close()
	//
	//if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
	//	log.Fatal(err)
	//}
	//defer pprof.StopCPUProfile()

	schema := schemago.CreateSchema(*schemaName, *numTables, *maxColumns)

	err := schemago.WriteSQLStatements(os.Stdout, schema)
	if err != nil {
		log.Fatal("Error:", err)
	}
}
