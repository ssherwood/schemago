package schemago

import (
	"fmt"
	"strings"
)

func GenerateSQLStatements(schema []Table) string {
	sql := ""

	for _, table := range schema {
		sql += fmt.Sprintf("\nCREATE TABLE%s%s(\n%s%s%s\n);\n%s",
			" IF NOT EXISTS ",
			table.Name,
			generatePrimaryKeys(table),
			generateColumns(table),
			generatePrimaryKeyConstraints(table),
			generateCreateIndexes(table))
	}

	return sql
}

func generatePrimaryKeys(table Table) string {
	sql := ""

	if len(table.PrimaryKeys) == 1 {
		sql += fmt.Sprintf("\t%s %s PRIMARY KEY,\n", table.PrimaryKeys[0].Name, table.PrimaryKeys[0].Type)
	} else {
		for _, pk := range table.PrimaryKeys {
			sql += fmt.Sprintf("\t%s %s,\n", pk.Name, pk.Type)
		}
	}

	return sql
}

func generateColumns(table Table) string {
	sql := ""

	currentItem := 0
	for _, col := range table.Columns {
		sql += fmt.Sprintf("\t%s %s", col.Name, col.Type)

		if col.Length > 0 {
			sql += fmt.Sprintf("(%d)", col.Length)
		}

		// constraints
		if !col.Nullable {
			sql += " NOT NULL"
		}

		if len(col.Default) > 0 {
			sql += fmt.Sprintf(" DEFAULT %s", col.Default)
		}

		if col.Unique {
			sql += fmt.Sprintf(" UNIQUE")
		}

		currentItem += 1
		if currentItem != len(table.Columns) {
			sql += ",\n"
		}
	}

	return sql
}

func generatePrimaryKeyConstraints(table Table) string {
	sql := ""

	if len(table.PrimaryKeys) > 1 {
		var pkNames []string
		for _, pk := range table.PrimaryKeys {
			pkNames = append(pkNames, pk.Name)
		}

		sql += fmt.Sprintf(",\n\tPRIMARY KEY(%s)", strings.Join(pkNames, ","))
	}

	return sql
}

func generateCreateIndexes(table Table) string {
	sql := ""

	if len(table.Indexes) > 0 {
		sql += "\n"
		for _, v := range table.Indexes {

			var cols []string
			for i, c := range v.ColumnNames {
				if len(v.Ordering[i]) > 0 {
					cols = append(cols, c+" "+v.Ordering[i])
				} else {
					cols = append(cols, c) // HASH is implied
				}
			}

			unique := ""
			if v.Unique {
				unique = "UNIQUE "
			}

			sql += fmt.Sprintf("CREATE %sINDEX %s ON %s(%s);\n", unique, v.Name, v.TableName, strings.Join(cols, ","))
		}
	}

	return sql
}
