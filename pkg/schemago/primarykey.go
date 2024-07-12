package schemago

import (
	"fmt"
	"math/rand"
	"sort"
)

type PrimaryKey struct {
	Name    string
	Type    string
	Length  int
	Default string
}

func randomPrimaryKey() []PrimaryKey {
	var primaryKeys []PrimaryKey
	for i := 0; i < randomNumberOfPrimaryKeys(); i++ {
		dataType := randomPrimaryKeyDataType()
		primaryKeys = append(primaryKeys, PrimaryKey{
			Name:    randomPrimaryKeyName(i),
			Type:    dataType,
			Length:  randomPrimaryKeyLength(dataType),
			Default: randomPrimaryKeyDefaultExpression(dataType),
		})
	}
	return primaryKeys
}

func primaryKeyNames(primaryKeys []PrimaryKey, sorted bool) []string {
	var names []string
	for _, v := range primaryKeys {
		names = append(names, v.Name)
	}

	if sorted {
		sort.Strings(names)
	}

	return names
}

func randomNumberOfPrimaryKeys() int {
	switch randomPercent := rand.Intn(100); {
	case randomPercent < 85:
		return 1
	case randomPercent < 95:
		return 2
	default:
		return 3
	}
}

func randomPrimaryKeyName(i int) (name string) {
	if i == 0 {
		name = "id"
	} else {
		name = fmt.Sprintf("%s_id", randomDescriptor(1, ""))
	}
	return
}

func randomPrimaryKeyDataType() (dataType string) {
	if percentChance(25) {
		dataType = "BIGSERIAL"
	} else {
		dataType = "UUID"
	}
	return
}

func randomPrimaryKeyLength(_dataType string) int {
	return 0
}

func randomPrimaryKeyDefaultExpression(dataType string) (defaultExpression string) {
	switch dataType {
	case "DATE", "TIMESTAMP", "TIMESTAMPTZ":
		defaultExpression = "NOW()"
	case "INTEGER":
		if percentChance(25) {
			defaultExpression = "0"
		}
	case "UUID":
		// TODO ensure uuid-ossp extension?
	}
	return
}
