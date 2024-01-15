package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func MustLoadConfig(filename string) {
	if err := godotenv.Load(filename); err != nil {
		panic(fmt.Errorf("error loading config file: %w", err))
	}
}
