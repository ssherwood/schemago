package schemago

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

const commaSeparator string = ", "
const createTableFmt string = "\nCREATE TABLE IF NOT EXISTS %s.%s(\n%s%s%s\n);\n\n%s"
const createForeignKeyFmt string = "\nALTER TABLE %s.%s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s.%s(%s);"
const createIndexFmt string = "CREATE %sINDEX %s ON %s.%s(%s);\n"

func WriteSQLStatements(writer io.Writer, schema Schema) error {
	for _, table := range schema.Tables {
		_, err := fmt.Fprintf(writer, createTableFmt,
			schema.Name, table.Name,
			primaryKeys(table),
			generateColumns(table),
			primaryKeyConstraints(table),
			createIndexes(schema.Name, table))

		if err != nil {
			return err
		}
	}

	if len(schema.ForeignKeys) > 0 {
		if err := writeForeignKeyConstraints(writer, schema); err != nil {
			return err
		}
	}

	return nil
}

func primaryKeys(table Table) (sql string) {
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

func primaryKeyConstraints(table Table) (sql string) {
	// only generated if there are multiple PKs
	if len(table.PrimaryKeys) > 1 {
		var pkNames []string
		for _, pk := range table.PrimaryKeys {
			pkNames = append(pkNames, pk.Name)
		}

		sql += fmt.Sprintf(",\n\tPRIMARY KEY(%s)", strings.Join(pkNames, ","))
	}

	return
}

func createIndexes(schemaName string, table Table) (sql string) {
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

		sql += fmt.Sprintf(createIndexFmt, unique, index.Name, schemaName, index.TableName, strings.Join(columnNames, commaSeparator))
	}

	return
}

func writeForeignKeyConstraints(output io.Writer, schema Schema) error {
	// sort by child table to keep ALTER TABLE statements together
	foreignKeys := schema.ForeignKeys
	sort.Slice(foreignKeys, func(i, j int) bool {
		return foreignKeys[i].ChildTableName < foreignKeys[j].ChildTableName
	})

	for _, foreignKey := range foreignKeys {
		_, err := fmt.Fprintf(output, createForeignKeyFmt,
			schema.Name,
			foreignKey.ChildTableName,
			foreignKey.Name,
			strings.Join(foreignKey.ChildTableColumns, commaSeparator),
			schema.Name,
			foreignKey.ParentTableName,
			strings.Join(foreignKey.ParentTableColumns, commaSeparator))

		if err != nil {
			return err
		}
	}

	return nil
}
