package main

import (
	"fmt"

	stunningdisco "github.com/karasibille/stunning-disco"
)

func main() {
	loader := stunningdisco.ConfigLoader{
		ConfigPath: "examples/basic.yml",
	}

	config := loader.MustLoad()

	for tableName, config := range config.Tables {
		fmt.Printf("--- table: %s\n", tableName)
		for colName, rule := range config.Rules {
			fmt.Printf("--- --- col: %s, strat: %v\n", colName, rule.Transformer)
		}
	}
}
