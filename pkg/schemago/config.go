package schemago

type Config struct {
	SQLType                string
	DefaultSchemaName      string
	NumberOfTables         int
	MaximumNumberOfColumns int
	EnumsEnabled           bool
	TabletSplitsEnabled    bool
}
