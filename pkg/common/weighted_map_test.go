package common

import (
	"testing"
)

func TestWeightedRandomMap_AddValue(t *testing.T) {
	tests := []testCaseWeightedRandomMap[string, int]{
		{"Empty", WeightedRandomMap[string]{}, []addValueArg[string]{}, 0},
		{"Single", WeightedRandomMap[string]{}, []addValueArg[string]{{"one", 1}}, 1},
		{"Multiple", WeightedRandomMap[string]{}, []addValueArg[string]{{"one", 1}, {"two", 2}, {"three", 3}}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tt.weightedMap.AddValue(arg.value, arg.weight)
			}
			if got := tt.weightedMap.totalWeight(); got != tt.wanted {
				t.Errorf("after AddValue() totalWeight() = %v, wanted %v", got, tt.wanted)
			}
		})
	}
}

func TestWeightedRandomMap_totalWeight(t *testing.T) {
	tests := []testCaseWeightedRandomMap[string, int]{
		{"Empty", WeightedRandomMap[string]{}, nil, 0},
		{"Single", WeightedRandomMap[string]{"one": 1}, nil, 1},
		{"Multiple", WeightedRandomMap[string]{"one": 1, "two": 2, "three": 3}, nil, 6},
		// the simplicity of the interface doesn't prevent non-positive entries that could break the map
		{"Dysfunctional", WeightedRandomMap[string]{"negative_one": -1, "negative_two": -2, "negative_three": -3}, nil, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.weightedMap.totalWeight(); got != tt.wanted {
				t.Errorf("totalWeight() = %v, wanted %v", got, tt.wanted)
			}
		})
	}
}

//
//func TestWeightedRandomMap_RandomValue(t *testing.T) {
//	type testCaseWeightedRandomMap[T comparable] struct {
//		name string
//		weightedMap    WeightedRandomMap[T]
//		wanted T
//	}
//	tests := []testCaseWeightedRandomMap[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.weightedMap.RandomValue(); !reflect.DeepEqual(got, tt.wanted) {
//				t.Errorf("RandomValue() = %v, wanted %v", got, tt.wanted)
//			}
//		})
//	}
//}

type addValueArg[T comparable] struct {
	value  T
	weight int
}

type testCaseWeightedRandomMap[T comparable, W any] struct {
	name        string
	weightedMap WeightedRandomMap[T]
	args        []addValueArg[T]
	wanted      W
}
