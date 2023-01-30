package config

import "time"

type Config struct {
	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port string `yaml:"port"`

		Timeout struct {
			// Server is the general server timeout to use
			// for graceful shutdowns
			Server time.Duration `yaml:"server"`

			// Write is the amount of time to wait until an HTTP server
			// write opperation is cancelled
			Write time.Duration `yaml:"write"`

			// Read is the amount of time to wait until an HTTP server
			// read operation is cancelled
			Read time.Duration `yaml:"read"`

			// Read is the amount of time to wait
			// until an IDLE HTTP session is closed
			Idle time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`

	// database
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DbName   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`

	Caching struct {
		Address     string        `yaml:"address"`
		ExpiredTime time.Duration `yaml:"expiredTime"`
	} `yaml:"caching"`
}
