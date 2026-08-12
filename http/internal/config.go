package internal

import (
	"fmt"
)

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (c *HTTPConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *HTTPConfig) ServiceName() string {
	return "http"
}
