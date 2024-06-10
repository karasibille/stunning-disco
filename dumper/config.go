package dumper

import (
	"gopkg.in/yaml.v3"
)

type DumperConfig struct {
	Tables map[TableName]TableConfig `yaml:"tables"`
}

type TableName string

type TableConfig struct {
	Query QueryConfig                `yaml:"query"`
	Rules map[ColumnName]RulesConfig `yaml:"rules"`
}

type ColumnName string

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

func NewDumperConfig(content []byte) (*DumperConfig, error) {
	cfg := &DumperConfig{}

	err := yaml.Unmarshal(content, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
