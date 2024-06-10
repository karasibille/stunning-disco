package dumper_test

import (
	"testing"

	"github.com/karasibille/stunning-disco/dumper"
)

func Test_Returns_DumperConfig_When_Yaml_Content_Is_Valid(t *testing.T) {

	var content = `
tables:
  users:
    query:
      order_by: id
      per_page: 3
    rules:
      id:
        strategy: faker
        unique: true
        transformer: numberBetween
        arguments: [ 100, 200 ]
      name:
        strategy: faker
        unique: false
        locale: en_US
        transformer: firstName
`

	cfg, err := dumper.NewDumperConfig([]byte(content))

	if err != nil {
		t.Fatalf("dumper.NewDumperConfig: got error %s.", err)
	}

	if len(cfg.Tables) != 1 {
		t.Fatalf("dumper.NewDumperConfig: got %d tables, want 1.", len(cfg.Tables))
	}

	if _, found := cfg.Tables["users"]; !found {
		t.Fatalf("dumper.NewDumperConfig: missing cfg.Tables[%q].", "users")
	}

	query := cfg.Tables["users"].Query
	if query.OrderBy != "id" {
		msg := "dumper.NewDumperConfig: missing cfg.Tables[%q].Query.OrderBy to be %q, got %q"
		t.Fatalf(msg, "users", "id", query.OrderBy)
	}

	if query.PerPage != 3 {
		msg := "dumper.NewDumperConfig: missing cfg.Tables[%q].Query.PerPage to be %d, got %d"
		t.Fatalf(msg, "users", 3, query.PerPage)
	}

	if _, found := cfg.Tables["users"].Rules["id"]; !found {
		msg := "dumper.NewDumperConfig: missing cfg.Tables[%q].Rules[%q]"
		t.Fatalf(msg, "users", "id")
	}

	if _, found := cfg.Tables["users"].Rules["name"]; !found {
		msg := "dumper.NewDumperConfig: missing cfg.Tables[%q].Rules[%q]"
		t.Fatalf(msg, "users", "name")
	}

	transformer := cfg.Tables["users"].Rules["name"].Transformer
	if transformer != "firstName" {
		msg := "dumper.NewDumperConfig: want cfg.Tables[%q].Rules[%q] to be %q, got %q"
		t.Fatalf(msg, "users", "name", "firstName", transformer)
	}
}

func Test_Returns_Error_When_Yaml_Content_Is_Badly_Typed(t *testing.T) {
	var content = `
tables:
  users:
    query:
      order_by: 3
      per_page: id
    rules:
      id:
        name:
          strategy: faker
          unique: false
          locale: en_US
          transformer: firstName
`

	cfg, err := dumper.NewDumperConfig([]byte(content))

	if err == nil {
		t.Fatalf("dumper.NewDumperConfig: want error, got cfg %v", cfg)
	}
}
