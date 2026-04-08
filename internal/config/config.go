package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Actor represents a user session for testing BOLA and logic flaws
type Actor struct {
	Name     string            `yaml:"name"`
	Headers  map[string]string `yaml:"headers"` // Usually Authorization tokens
	EntityID string            `yaml:"entity_id"` // E.g., user_id or post_id belonging to this actor
}

// ConfigPath represents a single endpoint with its method and optional body
type EndpointConfig struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
	Body   string `yaml:"body"` // Optional JSON body for POST/PUT/PATCH
}

// Config is the main configuration file for Vex
type Config struct {
	ActorA    Actor            `yaml:"actor_a"`
	ActorB    Actor            `yaml:"actor_b"`
	Endpoints []string         `yaml:"endpoints"` // Simple string format "METHOD PATH"
	Complex   []EndpointConfig `yaml:"complex"`   // Detailed format with body
}

// LoadConfig reads and parses the YAML configuration file
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file '%s': %w", filename, err)
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse yaml in '%s': %w", filename, err)
	}

	return &conf, nil
}
