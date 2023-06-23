package schemago

import (
	"schemago/pkg/common"
	"strconv"
	"strings"
)

// TODO Plan to rework this and make datatypes a struct with more settings and functions -
// this will help clean up the ugliness that enums have caused

// use a weighted randomized map to determine relative probability of the data types to occur
var dataTypesRandomMap = common.WeightedRandomMap[string]{
	"BIGINT":      5,
	"BIT":         3,
	"BOOLEAN":     15,
	"VARCHAR":     45,
	"DATE":        10,
	"INTEGER":     15,
	"JSONB":       10,
	"REAL":        3,
	"SMALLINT":    5,
	"TEXT":        20,
	"TIMESTAMP":   5,
	"TIMESTAMPTZ": 15,
	"UUID":        10,
}

// randomDataType returns a weighted selection from the dataTypesRandomMap, this includes the
// name of the data type itself, its maximum length (if it has one) and a possible default
// expression.
func randomDataType() (string, int, string) {
	dataType := dataTypesRandomMap.RandomValue()
	maxLength := randomMaxLength(dataType)
	defaultExpression := randomDefaultExpression(dataType)
	return dataType, maxLength, defaultExpression
}

func randomMaxLength(dataType string) (maxLength int) {
	switch dataType {
	case "BIT":
		maxLength = randomPow2(8)
	case "VARCHAR":
		maxLength = randomPow2(16)
	}
	return
}

func randomDefaultExpression(dataType string) (defaultExpression string) {
	switch dataType {
	case "DATE", "TIMESTAMP", "TIMESTAMPTZ":
		defaultExpression = "NOW()"
	case "BOOLEAN":
		defaultExpression = strings.ToUpper(strconv.FormatBool(percentChance(50)))
	case "INTEGER":
		if percentChance(25) {
			defaultExpression = "0"
		}
	case "JSONB":
		if percentChance(50) {
			defaultExpression = "'{}'"
		}
	}
	return
}

func randomNullable(dataType string) bool {
	nullable := percentChance(25)
	switch dataType {
	case "BOOLEAN":
		nullable = false
	}
	return nullable
}
