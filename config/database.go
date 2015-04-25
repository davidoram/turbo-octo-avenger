package config

import (
	"path/filepath"
	"fmt"
	"os"
	"log"
	"github.com/kylelemons/go-gypsy/yaml"
)

// TODO not keen on returning a string here should be a struct of the connection params
func DataBaseURL(env string) (string, error) {

	cfgFile := filepath.Join("./db", "dbconf.yml")

	f, err := yaml.ReadFile(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return "error", err
	}
	open, err := f.Get(fmt.Sprintf("%s.open", env))
	if err != nil {
		return "error", err
	}
	open = os.ExpandEnv(open)
	return "user=root dbname=turbo_octo_avenger_development sslmode=disable host=localhost", nil
}
