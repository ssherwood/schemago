package schemago

import (
	"fmt"
	"math/rand"
	"sort"
)

type Column struct {
	Name         string
	Type         string
	Length       int
	Default      string
	Nullable     bool
	Unique       bool
	SchemaNeeded bool
}

func columnsByName(columns []Column, sorted bool) []string {
	var columnNames []string
	for _, col := range columns {
		columnNames = append(columnNames, col.Name)
	}

	if sorted {
		sort.Strings(columnNames)
	}

	return columnNames
}

func randomColumns(maxColumns int, dataTypes DataTypes, enums []Enum) map[string]Column {
	columns := map[string]Column{}

	numColumns := rand.Intn(maxColumns) + 1
	for i := 0; i < numColumns; i++ {
		attrName := randomColumnName()
		dataType := dataTypes.randomDataType()

		if dataType.Name == "[ENUM]" {
			if len(enums) == 0 {
				continue // this will skip enums
			}
			randomEnum := enums[rand.Intn(len(enums))]
			dataType.Name = randomEnum.Name
			dataType.DefaultExpression = fmt.Sprintf("'%s'", randomEnumDefault(randomEnum))
		}

		columns[attrName] = Column{
			Name:         attrName,
			Type:         dataType.Name,
			Length:       dataType.randomMaxLength(),
			Default:      dataType.randomDefaultExpression(),
			Nullable:     dataType.randomNullable(),
			SchemaNeeded: dataType.PrependSchema,
		}
	}

	return columns
}

func randomColumnName() (columnName string) {
	switch r := rand.Intn(100); {
	case r < 60:
		columnName = randomDescriptor(1, "")
	case r < 90:
		columnName = randomDescriptor(2, "_")
	default:
		columnName = randomDescriptor(3, "_")
	}
	return
}
