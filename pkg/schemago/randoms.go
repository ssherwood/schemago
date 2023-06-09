package schemago

import (
	"github.com/Pallinder/go-randomdata"
	"math"
	"math/rand"
	"strings"
)

func randomDescriptor(words int, separator string) string {
	nameArray := []string{randomdata.Noun()}
	if words > 1 {
		for i := 2; i < words; i++ {
			nameArray = append(nameArray, randomdata.Noun())
		}
		// prepend an adjective
		nameArray = append([]string{randomdata.Adjective()}, nameArray...)
	}
	return strings.Join(nameArray, separator)
}

func percentChance(percent int) bool {
	return rand.Intn(100) < percent
}

func randomPow2(maxExponent int) int {
	return int(math.Pow(2.0, float64(rand.Intn(maxExponent+1))))
}
