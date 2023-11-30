package schemago

type Schema struct {
	Name        string
	Enums       []Enum
	Tables      []Table
	ForeignKeys []ForeignKey
}

// CreateSchema generates a certain number of tables with randomized columnsSQL and data types.
// provide the numTables to generate and the max possible number of columnsSQL on any one table.
func CreateSchema(config Config) Schema {

	dataTypes := LoadDataTypes(config.SQLType)

	schemaName := config.DefaultSchemaName
	if schemaName == "random" {
		schemaName = randomSchemaName()
	}

	var enums []Enum
	if config.EnumsEnabled {
		enums = randomEnums(10)
	}

	tables := randomTables(config, dataTypes, enums)

	return Schema{
		Name:        schemaName,
		Enums:       enums,
		Tables:      tables,
		ForeignKeys: generateForeignKeys(tables, 30),
	}
}

func randomSchemaName() string {
	return randomDescriptor(1, "")
}
