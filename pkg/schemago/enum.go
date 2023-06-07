package schemago

import (
	"math/rand"
	"sort"
	"strings"
)

type Enum struct {
	Name   string
	Values []string
}

func randomEnums(maxEnums int) (enums []Enum) {
	for i := 0; i < rand.Intn(maxEnums)+1; i++ {
		enums = append(enums, Enum{
			Name:   randomEnumName(),
			Values: randomEnumValues(10),
		})
	}
	return
}

func randomEnumName() (enumName string) {
	switch r := rand.Intn(100); {
	case r < 90:
		enumName = randomDescriptor(2, "_")
	default:
		enumName = randomDescriptor(3, "_")
	}
	return
}

func randomEnumValues(maxEnumValues int) (enumValues []string) {
	for i := 0; i < rand.Intn(maxEnumValues)+2; i++ {
		enumValues = append(enumValues, strings.ToUpper(randomDescriptor(2, "_")))
	}
	return
}

func sortEnumsByName(enums []Enum) {
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})
}
