package schemago

type Schema struct {
	Name        string
	Enums       []Enum
	Tables      []Table
	ForeignKeys []ForeignKey
}

// CreateSchema generates a certain number of tables with randomized columns and data types.
// provide the numTables to generate and the max possible number of columns on any one table.
func CreateSchema(numTables int, maxColumns int) Schema {
	enums := randomEnums(10)
	tables := randomTables(numTables, maxColumns, enums)

	return Schema{
		Name:        "public",
		Enums:       enums,
		Tables:      tables,
		ForeignKeys: generateForeignKeys(tables, 30),
	}
}
