package schemago

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type ForeignKey struct {
	Name               string
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
		parentTable := tables[rand.Intn(numTables)]
		childTable := tables[rand.Intn(numTables)]

		for childTable.Name == parentTable.Name {
			// don't connect to ourselves
			childTable = tables[rand.Intn(numTables)]
		}

		var childFKColumns []Column
		for _, col := range parentTable.PrimaryKeys {
			var columnName string
			if col.Name == "id" {
				columnName = fmt.Sprintf("%s_id", parentTable.Name)
			} else {
				columnName = col.Name
			}

			childTableColumn := Column{
				Name:     columnName,
				Type:     col.Type,
				Length:   col.Length,
				Default:  col.Default,
				Nullable: false,
				Unique:   false,
			}

			childTable.Columns[columnName] = childTableColumn // add it to the child table
			childFKColumns = append(childFKColumns, childTableColumn)
		}

		foreignKeys = append(foreignKeys, ForeignKey{
			Name:               fmt.Sprintf("%s_%s_fk", parentTable.Name, strings.Join(primaryKeyNames(parentTable.PrimaryKeys, false), "_")),
			ChildTableName:     childTable.Name,
			ChildTableColumns:  columnsByName(childFKColumns, true),
			ParentTableName:    parentTable.Name,
			ParentTableColumns: primaryKeyNames(parentTable.PrimaryKeys, false),
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

func sortForeignKeysByName(foreignKeys []ForeignKey) {
	sort.Slice(foreignKeys, func(i, j int) bool {
		return foreignKeys[i].ChildTableName < foreignKeys[j].ChildTableName
	})
}
