package schemago

import "github.com/Pallinder/go-randomdata"

type Table struct {
	Name        string
	PrimaryKeys []Column
	Columns     map[string]Column
	Indexes     []Index
}

func randomTableName() string {
	return randomdata.Adjective() + "_" + randomdata.Noun()
}
