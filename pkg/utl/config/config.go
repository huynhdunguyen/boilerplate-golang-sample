package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Config)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Config main
type Config struct {
	Server  *Server   `yaml:"server,omitempty"`
	DB      *Database `yaml:"database,omitempty"`
	JWT     *JWT      `yaml:"jwt,omitempty"`
	Swagger *Swagger  `yamrl:"swagger,omitempty"`
	Log     *Log      `yaml:"log,omitempty"`
}

// Mssql configuration
type Pgql struct {
	DSN        string `yaml:"dsn,omitempty"`
	LogQueries bool   `yaml:"log_queries,omitempty"`
	Timeout    int    `yaml:"timeout_seconds,omitempty"`
}

// Couchbase configuration
type Mongo struct {
	DSN string `yaml:"dsn,omitempty"`
}

// Database configuration
type Database struct {
	Pgql  *Pgql  `yaml:"pgql,omitempty"`
	Mongo *Mongo `yaml:"mongo,omitempty"`
}

// Server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
	GinMode      string `yaml:"gin_mode,omitempty"`
}

// JWT configuration
type JWT struct {
	Secret           string `yaml:"secret,omitempty"`
	Duration         int    `yaml:"duration_minutes,omitempty"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes,omitempty"`
	MaxRefresh       int    `yaml:"max_refresh_minutes,omitempty"`
	SigningAlgorithm string `yaml:"signing_algorithm,omitempty"`
}

// Swagger configuration
type Swagger struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
}

// Log  configuration
type Log struct {
	EnableConsole     bool   `yaml:"enable_console,omitempty"`
	ConsoleLevel      string `yaml:"console_level,omitempty"`
	ConsoleJSONFormat bool   `yaml:"console_json_format,omitempty"`
	EnableFile        bool   `yaml:"enable_file,omitempty"`
	FileLevel         string `yaml:"file_level,omitempty"`
	FileJSONFormat    bool   `yaml:"file_json_format,omitempty"`
}
