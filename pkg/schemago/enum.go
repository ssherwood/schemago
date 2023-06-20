package schemago

import (
	"math/rand"
	"schemago/pkg/common"
	"sort"
	"strings"
)

type empty = struct{}

// Enum - Enumerated types are data types that comprise a static, ordered set of values. They are equivalent
// to the enum types supported in a number of programming languages. An example of an enum type might be the
// days of the week, or a set of status values for a piece of data.
//
// Ex. CREATE TYPE mood AS ENUM ('sad', 'ok', 'happy');
type Enum struct {
	Name   string
	Values map[string]empty
}

// randomEnums returns a slice of enums up to the length of the maxEnums value.
// TODO duplicate enums can be produced by this method
func randomEnums(maxEnums int) (enums []Enum) {
	numberOfEnums := rand.Intn(maxEnums) + 1
	for i := 0; i < numberOfEnums; i++ {
		enums = append(enums, Enum{
			Name:   randomEnumName(),
			Values: randomEnumValues(10),
		})
	}
	return
}

// randomEnumName returns a randomized string to represent the identifier of the enumeration.  This
// implementation makes a randomized determination if the name should be one, two or sometimes three
// words concatenated by an underscore (snake_case).
func randomEnumName() (enumName string) {
	switch r := rand.Intn(100); {
	case r < 90:
		enumName = randomDescriptor(2, "_")
	default:
		enumName = randomDescriptor(3, "_")
	}
	return
}

// randomEnumValues returns a set of possible values for a given enum.  Each enum value is returned
// as a 2 word value, upper-cased and concatenated with an underscore.
func randomEnumValues(maxEnumValues int) map[string]empty {
	enumValues := map[string]empty{}
	for i := 0; i < rand.Intn(maxEnumValues)+2; i++ {
		value := strings.ToUpper(randomDescriptor(2, "_"))
		enumValues[value] = empty{}
	}
	return enumValues
}

// randomEnumDefault returns a random value entry
func randomEnumDefault(enums Enum) (attrDefault string) {
	values := common.KeysOfMap(enums.Values)
	return values[rand.Intn(len(values))]
}

// sortEnumsByName modifies the slice to order the entries by name alphabetically.
func sortEnumsByName(enums []Enum) {
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})
}
