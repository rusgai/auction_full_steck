package config

import (
	"os"
)

type Config interface {
	Get(key string) string
}
type config struct{}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}
