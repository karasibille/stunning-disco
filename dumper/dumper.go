package dumper

import (
	"fmt"
	"io"
)

type Dumper struct {
	Config *DumperConfig
	Writer io.Writer
}

func (d *Dumper) Dump() {
	// TODO Dump schemas
	// TODO Dump data
	// TODO -> Need Database connexion + sql builder ?
	io.WriteString(d.Writer, "Dump !\n")

	for tableName, config := range d.Config.Tables {
		io.WriteString(d.Writer, fmt.Sprintf("--- table: %s\n", tableName))

		for colName, rule := range config.Rules {
			io.WriteString(d.Writer, fmt.Sprintf("--- --- col: %s, strat: %v\n", colName, rule.Transformer))
		}
	}
}
