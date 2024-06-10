package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigLoader struct {
	Path   string
	Config DumperConfg
}

type TableName string

type DumperConfg struct {
	Tables map[TableName]TableConfig `yaml:"tables"`
}

type ColumnName string

type TableConfig struct {
	// Name  string
	Query QueryConfig                `yaml:"query"`
	Rules map[ColumnName]RulesConfig `yaml:"rules"`
}

type QueryConfig struct {
	OrderBy string `yaml:"order_by"`
	PerPage int    `yaml:"per_page"`
}

type RulesConfig struct {
	Strategy    string `yaml:"strategy"`
	Unique      bool   `yaml:"unique"`
	Locale      string `yaml:"locale"`
	Transformer string `yaml:"transformer"`
	Arguments   []any  `yaml:"arguments"`
}

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	t := DumperConfg{}

	path := "examples/basic.yml"
	data := load(path)

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for tableName, config := range t.Tables {
		fmt.Printf("--- table: %s", tableName)
		for colName, rule := range config.Rules {
			fmt.Printf("--- --- col: %s, strat: %v\n", colName, rule.Transformer)
		}
	}
}

func load(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
