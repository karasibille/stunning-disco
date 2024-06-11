package main

import (
	"os"

	stunningdisco "github.com/karasibille/stunning-disco"
	"github.com/karasibille/stunning-disco/dumper"
)

func main() {
	loader := stunningdisco.ConfigLoader{
		ConfigPath: "examples/basic.yml",
	}

	config := loader.MustLoad()
	// TODO Configure writer so that we can dump the result in a file.
	dumper := dumper.Dumper{Config: config, Writer: os.Stdout}

	dumper.Dump()
}
