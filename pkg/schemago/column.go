package schemago

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"math/rand"
)

type Column struct {
	Name     string
	Type     string
	Length   int
	Default  string
	Nullable bool
	Unique   bool
}

func randomColumns(maxColumns int) map[string]Column {
	columns := map[string]Column{}

	numColumns := rand.Intn(maxColumns) + 1
	for i := 0; i < numColumns; i++ {
		attrName := randomColumnName()
		attrType, attrLength, attrDefault := randomDataType()
		attrNullable := randomNullable(attrType)

		columns[attrName] = Column{
			Name:     attrName,
			Type:     attrType,
			Length:   attrLength,
			Default:  attrDefault,
			Nullable: attrNullable,
		}
	}

	return columns
}

func randomColumnName() string {
	r := rand.Intn(100)

	if r < 60 {
		return randomdata.Noun()
	} else if r < 90 {
		return fmt.Sprintf("%s_%s", randomdata.Adjective(), randomdata.Noun())
	} else {
		return fmt.Sprintf("%s_%s_%s", randomdata.Adjective(), randomdata.Noun(), randomdata.Adjective())
	}
}
