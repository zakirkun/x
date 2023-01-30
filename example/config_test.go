package main

import (
	"fmt"
	"testing"

	"github.com/zakirkun/x/config/yaml"
	"github.com/zakirkun/x/database"
)

func TestConfig(t *testing.T) {
	config := yaml.NewYamlConfig("./config.yml")

	t.Logf(fmt.Sprintf("Log Config : %v", config))
}

func TestMysql(t *testing.T) {
	cfg := yaml.NewYamlConfig("./config.yml")

	dbOpt := database.DBModel{
		Driver:   cfg.Database.Driver,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		Name:     cfg.Database.DbName,
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
	}

	_, err := database.NewDatabase(dbOpt).InitDB()
	if err != nil {
		t.Errorf("Error : %v", err)
	}

	t.Log("Connection Success!")
}
