package yaml

import (
	"fmt"
	"os"

	"github.com/zakirkun/x/config"
	"github.com/zakirkun/x/logger"
	"gopkg.in/yaml.v2"
)

// init Yaml config
func NewYamlConfig(filename string) *config.Config {
	// Create config structure
	config := &config.Config{}

	file, err := os.Open(filename)
	if err != nil {
		logger.Warn(fmt.Sprintf("Failed load config : %v", err))
		return nil
	}
	defer file.Close()

	// Init new YAML decode
	decode := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := decode.Decode(&config); err != nil {
		logger.Warn(fmt.Sprintf("Failed parse config : %v", err))
		return nil
	}

	return config
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
