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
		Name:        "",
		Tables:      tables,
		ForeignKeys: generateForeignKeys(tables, 30),
	}
}

func randomTables(numTables int, maxColumns int) []Table {
	var tables []Table
	for i := 0; i < numTables; i++ {
		tableName := randomTableName()
		attributes := randomColumns(maxColumns)
		tables = append(tables, Table{
			Name:        tableName,
			PrimaryKeys: randomPrimaryKey(),
			Columns:     attributes,
			Indexes:     randomIndexes(tableName, attributes),
		})
	}

	return tables
}
