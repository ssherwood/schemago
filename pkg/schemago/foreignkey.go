package schemago

import (
	"errors"
	"math/rand"
)

// ALTER TABLE child_table
//  ADD CONSTRAINT constraint_name
//    FOREIGN KEY (fk_columns)
//      REFERENCES parent_table (parent_key_columns);

type ForeignKey struct {
	ChildTableName     string
	ChildTableColumns  []string
	ParentTableName    string
	ParentTableColumns []string
}

func generateForeignKeys(tables []Table, percentage int) []ForeignKey {
	var foreignKeys []ForeignKey

	numTables := len(tables)
	numForeignKeys := (numTables * percentage) / 100

	for i := 0; i < numForeignKeys; i++ {
		childTable := tables[rand.Intn(numTables)]
		childColumn, ok := randomColumn(childTable.Columns)
		for ok != nil {
			for j := 0; j < 100; j++ {
				childColumn, ok = randomColumn(childTable.Columns)
			}

			if ok != nil {
				// escape, we cant generate proper keys
				return foreignKeys
			}
		}

		foreignKeys = append(foreignKeys, ForeignKey{
			ChildTableName:     childTable.Name,
			ChildTableColumns:  []string{childColumn.Name},
			ParentTableName:    "", // TODO
			ParentTableColumns: nil,
		})
	}

	return foreignKeys
}

func randomColumn(columns map[string]Column) (Column, error) {
	var arr []string
	for k, v := range columns {
		if v.Type == "VARCHAR" || v.Type == "BIGINT" || v.Type == "UUID" {
			arr = append(arr, k)
		}
	}

	if len(arr) == 0 {
		return Column{}, errors.New("no suitable columns")
	}

	return columns[arr[rand.Intn(len(arr))]], nil
}
