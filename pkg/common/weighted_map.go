package common

import "math/rand"

// WeightedRandomMap is a simplistic implementation of a map for use in random
// generation from a set of values.  Each entry is assigned a numeric totalWeight (or
// frequency of its chance to be emitted by random selection).
// TODO this isn't particularly efficient implementation but is pretty simple
type WeightedRandomMap[T comparable] map[T]int

// TotalWeight sums the individual weights to determine the total weight of the
// underlying map
func (m *WeightedRandomMap[T]) TotalWeight() int {
	totalWeight := 0
	for _, weight := range *m {
		totalWeight += weight
	}
	return totalWeight
}

// AddValue is an alternate method for adding values to the map
func (m *WeightedRandomMap[T]) AddValue(value T, weight int) *WeightedRandomMap[T] {
	(*m)[value] = weight
	return m
}

// RandomValue selects a random entry from the map based on the weighted factor of each
// entry in the map, the map's total weight and a randomized value within that range.
func (m *WeightedRandomMap[T]) RandomValue() T {
	cumulativeWeight := 0
	targetWeight := rand.Intn(m.TotalWeight())

	for key, weight := range *m {
		cumulativeWeight += weight
		if cumulativeWeight > targetWeight {
			return key
		}
	}

	var defaultValue T
	return defaultValue
}
