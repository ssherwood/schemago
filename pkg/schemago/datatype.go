package schemago

import (
	"math"
	"math/rand"
	"schemago/pkg/common"
	"strconv"
	"strings"
)

// use a weighted randomized map to determine relative probability of the data types to occur
var dataTypesRandomMap = common.WeightedRandomMap[string]{
	"BIGINT":      5,
	"BIT":         3,
	"BOOLEAN":     15,
	"VARCHAR":     50,
	"DATE":        10,
	"INTEGER":     15,
	"JSONB":       10,
	"REAL":        3,
	"SMALLINT":    5,
	"TEXT":        20,
	"TIMESTAMP":   5,
	"TIMESTAMPTZ": 10,
	"UUID":        10,
	"[ENUM]":      15,
}

func randomDataType() (string, int, string) {
	dataType := dataTypesRandomMap.RandomValue()
	maxLength := randomMaxLength(dataType)
	defaultExpression := randomDefaultExpression(dataType)

	return dataType, maxLength, defaultExpression
}

func randomPow2(maxExponent int) int {
	return int(math.Pow(2.0, float64(rand.Intn(maxExponent+1))))
}

func randomMaxLength(dataType string) int {
	maxLength := 0

	switch dataType {
	case "BIT":
		maxLength = randomPow2(8)
	case "VARCHAR":
		maxLength = randomPow2(16)
	}

	return maxLength
}

func randomDefaultExpression(dataType string) string {
	defaultExpression := ""

	switch dataType {
	case "DATE", "TIMESTAMP", "TIMESTAMPTZ":
		defaultExpression = "NOW()"
	case "BOOLEAN":
		defaultExpression = strings.ToUpper(strconv.FormatBool(rand.Intn(2) == 1))
	case "INTEGER":
		if rand.Intn(100) < 25 {
			defaultExpression = "0"
		}
	case "JSONB":
		if rand.Intn(100) < 50 {
			defaultExpression = "'{}'"
		}
	}

	return defaultExpression
}

func randomNullable(dataType string) bool {
	nullable := rand.Intn(100) < 25

	switch dataType {
	case "BOOLEAN":
		nullable = false
	}

	return nullable
}
