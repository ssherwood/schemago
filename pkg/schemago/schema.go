package schemago

type Schema struct {
	Name        string
	Tables      []Table
	ForeignKeys []ForeignKey
}

// CreateSchema generates a certain number of tables with randomized columns and data types.
// provide the numTables to generate and the max possible number of columns on any one table.
func CreateSchema(numTables int, maxColumns int) Schema {
	tables := randomTables(numTables, maxColumns)

	return Schema{
		Name:        "public",
		Tables:      tables,
		ForeignKeys: generateForeignKeys(tables, 30),
	}
}
