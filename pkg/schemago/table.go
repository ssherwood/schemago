package schemago

import "math/rand"

type Table struct {
	Name         string
	PrimaryKeys  []PrimaryKey
	Columns      map[string]Column
	Indexes      []Index
	TabletSplits int
}

func randomTables(config Config, dataTypes DataTypes, enums []Enum) []Table {
	var tables []Table
	for i := 0; i < config.NumberOfTables; i++ {
		tableName := randomTableName()
		attributes := randomColumns(config.MaximumNumberOfColumns, dataTypes, enums)
		tabletSplits := 0
		if config.TabletSplitsEnabled {
			tabletSplits = randomTabletSplits()
		}

		tables = append(tables, Table{
			Name:         tableName,
			PrimaryKeys:  randomPrimaryKey(),
			Columns:      attributes,
			Indexes:      randomIndexes(tableName, attributes),
			TabletSplits: tabletSplits,
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

// use for tablet splits if applicable
func randomTabletSplits() (tabletSplits int) {
	switch r := rand.Intn(100); {
	case r < 10:
		tabletSplits = 1
	case r < 55:
		tabletSplits = 40
	default:
		tabletSplits = 100
	}
	return
}
