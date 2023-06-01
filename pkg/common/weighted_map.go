package common

import "math/rand"

type WeightedRandomMap[T comparable] map[T]int

func (m *WeightedRandomMap[T]) AddValue(value T, weight int) {
	(*m)[value] = weight
}

func (m *WeightedRandomMap[T]) totalWeight() int {
	totalWeight := 0
	for _, weight := range *m {
		totalWeight += weight
	}
	return totalWeight
}

func (m *WeightedRandomMap[T]) RandomValue() T {
	totalWeight := m.totalWeight()
	randomNumber := rand.Intn(totalWeight)

	cumulativeWeight := 0
	for key, weight := range *m {
		cumulativeWeight += weight
		if cumulativeWeight > randomNumber {
			return key
		}
	}

	var defaultValue T
	return defaultValue
}
