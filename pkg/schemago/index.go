package schemago

import "fmt"

type Index struct {
	Name        string
	TableName   string
	ColumnNames []string
	Ordering    []string
	Unique      bool
}

func randomIndexes(tableName string, attributes map[string]Column) []Index {
	var indexes []Index

	for _, v := range attributes {
		// TODO make this randomized
		if v.Type == "VARCHAR" {
			indexes = append(indexes, Index{
				Name:        fmt.Sprintf("%s_%s_idx", tableName, v.Name),
				TableName:   tableName,
				ColumnNames: []string{v.Name},
				Ordering:    []string{""},
			})
		}
	}

	return indexes
}
