# Schemago

Schemago is a simple CLI for generating randomized database schemas.

It is highly unlikely that you will ever need this tool.

## To Build

Requires golang 1.20+

```shell
$ go mod download
$ go build -v ./cmd/schemago
```

## Usage

```shell
$ schemago [-tables n] [-columns m]
```
### Options

- `-tables` will generate the specific number of tables
- `-columns` will randomly generate up to this number of columns per table

## Details

Use `schemago` to generate output that can be used as a starting schema for a randomly sized database.  This tool is
primarily designed for testing databases for certain load and/or scale requirements.

Column attributes are randomly generated with specific data types having higher frequency of being generated than
others.  Appropriate lengths and defaults are also randomly generated with certain frequency that attempts to model a
somewhat realistic database schema.

Indexes and Foreign Keys are also generated with a randomized frequency on appropriate data types.

## Future Enhancements

- Additional configuration options:
  - Probability configurations 
  - Data types and frequency weights
- Additional randomized objects:
  - Views
  - Constraints
  - Functions
  - Enums