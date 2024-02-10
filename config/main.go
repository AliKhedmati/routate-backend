package config

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
	"sync"
)

var instance *Config

// Config represents the configuration structure.
type Config struct {
	data map[string]string
	mu   sync.RWMutex
}

// Init initializes the configuration using go dotenv.
func Init() error {
	instance := GetInstance()
	return instance.Load()
}

// GetInstance returns the singleton instance of Config.
func GetInstance() *Config {
	if instance == nil {
		instance = &Config{}
	}
	return instance
}

// Get gets the value for given key.
func (c *Config) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// Set sets config values.
func (c *Config) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Load loads the configuration from the specified file using go dotenv.
func (c *Config) Load() error {
	// Lock the instance.
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	// Loop through environment variables and populate the Config struct
	c.data = make(map[string]string)

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		key, value := pair[0], pair[1]
		c.data[key] = value
	}

	return nil
}
