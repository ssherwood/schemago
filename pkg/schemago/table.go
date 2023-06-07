package schemago

import "math/rand"

type Table struct {
	Name        string
	PrimaryKeys []PrimaryKey
	Columns     map[string]Column
	Indexes     []Index
}

func randomTables(numTables int, maxColumns int, enums []Enum) []Table {
	var tables []Table
	for i := 0; i < numTables; i++ {
		tableName := randomTableName()
		attributes := randomColumns(maxColumns, enums)

		tables = append(tables, Table{
			Name:        tableName,
			PrimaryKeys: randomPrimaryKey(),
			Columns:     attributes,
			Indexes:     randomIndexes(tableName, attributes),
		})
	}
	return tables
}

func randomTableName() (tableName string) {
	switch r := rand.Intn(100); {
	case r < 25:
		tableName = randomDescriptor(1, "")
	case r < 95:
		tableName = randomDescriptor(2, "_")
	default:
		tableName = randomDescriptor(3, "_")
	}
	return
}
