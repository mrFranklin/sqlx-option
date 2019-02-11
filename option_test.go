package sqlxopt

import (
	"testing"
)

const (
	configFile = "./config_test.yaml"
	key        = "mysql.dev"
)

func TestLoadOptionFromConfigFile(t *testing.T) {
	option, err := LoadOptionFromConfigFile(configFile, "mysql.dev")
	if err != nil {
		t.Fatal(err)
	}

	opt := &Option{
		User: "root",
		Password: "",
		Host: "127.0.0.1",
		Port: 3306,
		Charset: "utf8mb4",
		DbName: "sqlx",
	}


	t.Logf("Load option success! %s", option.String())
}
