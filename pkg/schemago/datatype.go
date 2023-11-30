package schemago

import (
	"gopkg.in/yaml.v3"
	"log"
	"math/rand"
)

func LoadDataTypes(dbId string) DataTypes {
	var dataTypes DataTypes
	switch dbId {
	case "yugabytedb-ysql":
		if err := yaml.Unmarshal([]byte(yugabyteDBYSQLDataTypesConfig), &dataTypes); err != nil {
			log.Fatalf("error: %v", err)
			return dataTypes
		}
	}

	// calculate the distribution totalWeight
	totalWeight := 0
	for _, v := range dataTypes.DataTypes {
		totalWeight += v.DistributionWeight
	}
	dataTypes.TotalWeight = totalWeight

	return dataTypes
}

type DataTypes struct {
	Source      string     `yaml:"database"`
	DataTypes   []DataType `yaml:"dataTypes"`
	TotalWeight int
}

func (dts DataTypes) randomDataType() DataType {
	cumulativeWeight := 0
	targetWeight := rand.Intn(dts.TotalWeight)

	for i, dt := range dts.DataTypes {
		cumulativeWeight += dt.DistributionWeight
		if cumulativeWeight > targetWeight {
			return dts.DataTypes[i]
		}
	}

	// else?
	return dts.DataTypes[0]
}

type DataType struct {
	Name                string `yaml:"name"`
	MaxLengthRandomPow  int    `yaml:"maxLengthRandomPow"`
	NullablePercentage  int    `yaml:"nullablePercentage"`
	DefaultedPercentage int    `yaml:"defaultedPercentage"`
	DefaultExpression   string `yaml:"defaultExpression"`
	ValuesQuoted        bool   `yaml:"valuesQuoted"`
	DistributionWeight  int    `yaml:"distributionWeight"`
	PrependSchema       bool   `yaml:"prependSchema"`
}

func (dt DataType) randomMaxLength() int {
	if dt.MaxLengthRandomPow > 0 {
		return randomPow2(dt.MaxLengthRandomPow)
	}

	return 0
}

func (dt DataType) randomNullable() bool {
	return percentChance(dt.NullablePercentage)
}

func (dt DataType) randomDefaultExpression() string {
	if dt.DefaultedPercentage > 0 {
		return dt.DefaultExpression
	}

	return ""
}
