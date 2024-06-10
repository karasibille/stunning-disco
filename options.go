package stunningdisco

import (
	"os"

	"github.com/karasibille/stunning-disco/dumper"
)

type ConfigLoader struct {
	ConfigPath string
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

func (cl *ConfigLoader) Load() (*dumper.DumperConfig, error) {
	content, err := os.ReadFile(cl.ConfigPath)
	if err != nil {
		return nil, err
	}

	return dumper.NewDumperConfig(content)
}

func (cl *ConfigLoader) MustLoad() *dumper.DumperConfig {
	cfg, err := cl.Load()

	if err != nil {
		panic(err)
	}

	return cfg
}
