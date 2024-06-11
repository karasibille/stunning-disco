package stunningdisco

import (
	"io"

	"github.com/karasibille/stunning-disco/dumper"
)

// ConfigLoader is a [DumperConfig] Factory.
type ConfigLoader struct {
	// Use the provided [io.Reader] to read yaml configuration.
	Reader io.Reader
	// todo
	// todo {--D|db-driver=mysql : Database driver to use}
	// todo {--H|db-host=localhost : Database host to connect to}
	// todo {--P|db-port=3306 : Database port used for the connection}
	// todo {--N|db-name= : Database name to dump (or path to the sqlite database)}
	// todo {--U|db-user=root : Database user to use for the connection}
	// todo {--p|ask-password : You will be prompted for a password}
	// todo {--C|db-charset=utf8mb4 : Database charset}
	// todo {--c|db-collation=utf8mb4_unicode_ci : Database charset}
	// todo {config : Use the provided YAML config file}
}

// Load creates a new DumperConfig by reading the content of [ConfigLoader.Reader]
func (cl *ConfigLoader) Load() (*dumper.DumperConfig, error) {
	content, err := io.ReadAll(cl.Reader)
	if err != nil {
		return nil, err
	}

	return dumper.NewDumperConfig(content)
}

// MustLoad creates a new DumperConfig by reading the content of
// [ConfigLoader.Reader] and panics if the something goes wrong.
func (cl *ConfigLoader) MustLoad() *dumper.DumperConfig {
	cfg, err := cl.Load()

	if err != nil {
		panic(err)
	}

	return cfg
}
