package server

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	Concurrency int      `json:"concurrency"`
	Proxies     []string `json:"proxies"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config

	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Validate() error {
	if c.Concurrency <= 0 {
		return fmt.Errorf("concurrency must be greater than 0")
	}

	if len(c.Proxies) == 0 {
		return fmt.Errorf("proxies must be provided")
	}

	// Make sure that all proxies are valid
	for _, proxy := range c.Proxies {
		if _, err := url.Parse(proxy); err != nil {
			return fmt.Errorf("invalid proxy: %w", err)
		}
	}
	return nil
}
