package schemago

import (
	"math/rand"
)

func CreateSchema(numTables int, maxColumns int) []Table {
	var schema []Table

	for i := 0; i < numTables; i++ {
		tableName := randomTableName()
		attributes := randomColumns(maxColumns)
		schema = append(schema, Table{
			Name:        tableName,
			PrimaryKeys: randomPrimaryKey(),
			Columns:     attributes,
			Indexes:     randomIndexes(tableName, attributes),
		})
	}

	return schema
}

func randomPrimaryKey() []Column {
	// TODO sometimes we should generate composite PKs

	var dataType string
	if rand.Intn(100) < 25 {
		dataType = "BIGSERIAL"
	} else {
		dataType = "UUID"
	}

	return []Column{
		{
			Name:     "id",
			Type:     dataType,
			Length:   0,
			Default:  "",
			Nullable: false,
		},
	}
}
