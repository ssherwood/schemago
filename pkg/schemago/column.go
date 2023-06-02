package schemago

import (
	"math/rand"
	"sort"
)

type Column struct {
	Name     string
	Type     string
	Length   int
	Default  string
	Nullable bool
	Unique   bool
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

func randomColumns(maxColumns int) map[string]Column {
	columns := map[string]Column{}
	numColumns := rand.Intn(maxColumns) + 1

	for i := 0; i < numColumns; i++ {
		attrName := randomColumnName()
		attrType, attrLength, attrDefault := randomDataType()

		columns[attrName] = Column{
			Name:     attrName,
			Type:     attrType,
			Length:   attrLength,
			Default:  attrDefault,
			Nullable: randomNullable(attrType),
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
