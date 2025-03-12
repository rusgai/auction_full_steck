package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}
type config struct{}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}
func NewConfig(fileName ...string) Config {
	err := godotenv.Load(fileName...)
	if err != nil {
		panic(err)
	}
	return &config{}
}
