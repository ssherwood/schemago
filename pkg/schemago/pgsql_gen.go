package schemago

import (
	"fmt"
	"io"
	"strings"
)

const commaSeparator string = ", "
const createEnumFmt string = "CREATE TYPE %s.%s as ENUM(%s);\n"
const createTableFmt string = "CREATE TABLE IF NOT EXISTS %s.%s(\n%s%s%s\n);\n\n"
const createIndexFmt string = "CREATE %sINDEX %s ON %s.%s(%s);\n"
const createForeignKeyFmt string = "ALTER TABLE %s.%s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s.%s(%s);\n"

// WriteSQLStatements prints all the DDL needed to create the schema in the target database (today that is just
// PostgresSQL).
func WriteSQLStatements(writer io.Writer, schema Schema) error {
	if err := writeCreateTypeEnums(writer, schema.Name, schema.Enums); err != nil {
		return err
	}

	if err := writeCreateTables(writer, schema.Name, schema.Tables); err != nil {
		return err
	}

	if err := writeAlterTableAddConstraints(writer, schema.Name, schema.ForeignKeys); err != nil {
		return err
	}

	return nil
}

func writeCreateTypeEnums(output io.Writer, schemaName string, enums []Enum) error {
	sortEnumsByName(enums)
	for _, e := range enums {
		_, err := fmt.Fprintf(output, createEnumFmt, schemaName, e.Name, "'"+strings.Join(e.Values, "', '")+"'")
		if err != nil {
			return err
		}
	}

	if len(enums) > 0 {
		if _, err := fmt.Fprintln(output); err != nil {
			return err
		}
	}

	return nil
}

func writeCreateTables(writer io.Writer, schemaName string, tables []Table) error {
	// TODO should we sort these?
	for _, table := range tables {
		_, err := fmt.Fprintf(writer, createTableFmt, schemaName, table.Name, primaryKeys(table), columns(table), primaryKeyConstraints(table))
		if err != nil {
			return err
		}

		if err = writeCreateIndexes(writer, schemaName, table.Indexes); err != nil {
			return err
		}
	}

	return nil
}

func writeCreateIndexes(writer io.Writer, schemaName string, indexes []Index) error {
	// TODO sort index names?
	for _, index := range indexes {
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

		_, err := fmt.Fprintf(writer, createIndexFmt, unique, index.Name, schemaName, index.TableName, strings.Join(columnNames, commaSeparator))
		if err != nil {
			return err
		}
	}

	if len(indexes) > 0 {
		if _, err := fmt.Fprintln(writer); err != nil {
			return err
		}
	}

	return nil
}

func writeAlterTableAddConstraints(output io.Writer, schemaName string, foreignKeys []ForeignKey) error {
	sortForeignKeysByName(foreignKeys)
	for _, foreignKey := range foreignKeys {
		_, err := fmt.Fprintf(output, createForeignKeyFmt,
			schemaName,
			foreignKey.ChildTableName,
			foreignKey.Name,
			strings.Join(foreignKey.ChildTableColumns, commaSeparator),
			schemaName,
			foreignKey.ParentTableName,
			strings.Join(foreignKey.ParentTableColumns, commaSeparator))

		if err != nil {
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

func columns(table Table) (sql string) {
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
