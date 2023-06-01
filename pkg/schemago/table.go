package schemago

type Table struct {
	Name        string
	PrimaryKeys []PrimaryKey
	Columns     map[string]Column
	Indexes     []Index
}

func randomTableName() string {
	return randomDescriptor(2, "_")
}
