package main

import (
	"fmt"
	"testing"

	"github.com/zakirkun/x/config/yaml"
)

func TestConfig(t *testing.T) {
	config := yaml.NewYamlConfig("./config.yml")

	t.Logf(fmt.Sprintf("Log Config : %v", config))
}
