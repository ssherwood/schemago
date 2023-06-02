package schemago

import (
	"fmt"
	"sort"
	"strings"
)

func GenerateSQLStatements(schema Schema) string {
	sql := ""

	for _, table := range schema.Tables {
		sql += fmt.Sprintf("\nCREATE TABLE%s%s.%s(\n%s%s%s\n);\n%s\n",
			" IF NOT EXISTS ",
			schema.Name,
			table.Name,
			generatePrimaryKeys(table),
			generateColumns(table),
			generatePrimaryKeyConstraints(table),
			generateCreateIndexes(schema.Name, table))
	}

	if len(schema.ForeignKeys) > 0 {
		sql += generateForeignKeyConstraints(schema)
	}

	return sql
}

func generatePrimaryKeys(table Table) (sql string) {
	if len(table.PrimaryKeys) == 1 {
		sql += fmt.Sprintf("\t%s %s PRIMARY KEY,\n", table.PrimaryKeys[0].Name, table.PrimaryKeys[0].Type)
	} else {
		for _, pk := range table.PrimaryKeys {
			sql += fmt.Sprintf("\t%s %s,\n", pk.Name, pk.Type)
		}
	}
	return
}

func generateColumns(table Table) (sql string) {
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
	return
}

func generatePrimaryKeyConstraints(table Table) (sql string) {
	// only done if there are multiple PKs
	if len(table.PrimaryKeys) > 1 {
		var pkNames []string
		for _, pk := range table.PrimaryKeys {
			pkNames = append(pkNames, pk.Name)
		}

		sql += fmt.Sprintf(",\n\tPRIMARY KEY(%s)", strings.Join(pkNames, ","))
	}
	return
}

func generateCreateIndexes(schemaName string, table Table) (sql string) {
	for _, index := range table.Indexes {
		var columnNames []string
		for i, columnName := range index.ColumnNames {
			if len(index.Ordering[i]) > 0 {
				columnNames = append(columnNames, fmt.Sprintf("%s %s", columnName, index.Ordering[i]))
			} else {
				columnNames = append(columnNames, columnName) // HASH is implied
			}
		}

		unique := ""
		if index.Unique {
			unique = "UNIQUE "
		}

		sql += fmt.Sprintf("\nCREATE %sINDEX %s ON %s.%s(%s);", unique, index.Name, schemaName, index.TableName, strings.Join(columnNames, ","))
	}
	return
}

func generateForeignKeyConstraints(schema Schema) (sql string) {

	// sort by child table to keep ALTER TABLE statements together
	foreignKeys := schema.ForeignKeys
	sort.Slice(foreignKeys, func(i, j int) bool {
		return foreignKeys[i].ChildTableName < foreignKeys[j].ChildTableName
	})

	for _, foreignKey := range foreignKeys {
		// TODO ON DELETE ?
		sql += fmt.Sprintf("\nALTER TABLE %s.%s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s.%s(%s);\n",
			schema.Name,
			foreignKey.ChildTableName,
			foreignKey.Name,
			strings.Join(foreignKey.ChildTableColumns, ", "),
			schema.Name,
			foreignKey.ParentTableName,
			strings.Join(foreignKey.ParentTableColumns, ", "))
	}
	return
}
