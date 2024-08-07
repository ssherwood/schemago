package schemago

const yugabyteDBYSQLDataTypesConfig = `
source: yugabytedb-ysql
dataTypes:
  - name: "BIGINT"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 50
    defaultExpression: "b'0'"
    valuesQuoted: false
    prependSchema: false
    distributionWeight: 5
  - name: "BIT"
    maxLengthRandomPow: 8
    nullablePercentage: 25
    defaultedPercentage: 50
    defaultExpression: "0"
    valuesQuoted: false
    prependSchema: false
    distributionWeight: 3
  - name: "BOOLEAN"
    maxLengthRandomPow: 0
    nullablePercentage: 0
    defaultedPercentage: 50
    defaultExpression: "false"
    valuesQuoted: false
    prependSchema: false
    distributionWeight: 3
  - name: "VARCHAR"
    maxLengthRandomPow: 16
    nullablePercentage: 25
    defaultedPercentage: 0
    defaultExpression: ""
    valuesQuoted: true
    prependSchema: false
    distributionWeight: 45
  - name: "DATE"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 80
    defaultExpression: "NOW()"
    valuesQuoted: true
    prependSchema: false
    distributionWeight: 20
  - name: "INTEGER"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 25
    defaultExpression: "0"
    valuesQuoted: false
    prependSchema: false
    distributionWeight: 15
  - name: "JSONB"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 50
    defaultExpression: "'{}'"
    valuesQuoted: true
    prependSchema: false
    distributionWeight: 10
  - name: "REAL"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 50
    defaultExpression: "0.0"
    valuesQuoted: false
    prependSchema: false
    distributionWeight: 3
  - name: "SMALLINT"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 25
    defaultExpression: "0"
    valuesQuoted: false
    prependSchema: false
    distributionWeight: 5
  - name: "TEXT"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 0
    defaultExpression: ""
    valuesQuoted: true
    prependSchema: false
    distributionWeight: 20
  - name: "TIMESTAMP"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 80
    defaultExpression: "NOW()"
    valuesQuoted: true
    prependSchema: false
    distributionWeight: 5
  - name: "TIMESTAMPTZ"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 80
    defaultExpression: "NOW()"
    valuesQuoted: true
    prependSchema: false
    distributionWeight: 15
  - name: "UUID"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 0
    defaultExpression: ""
    valuesQuoted: true
    prependSchema: false
    distributionWeight: 10
  - name: "[ENUM]"
    maxLengthRandomPow: 0
    nullablePercentage: 25
    defaultedPercentage: 90
    defaultExpression: ""
    valuesQuoted: true
    prependSchema: true
    distributionWeight: 10
`
