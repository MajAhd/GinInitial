package config

import (
	"log"

	goconfigtree "github.com/MajAhd/go-config-tree"
)

func Local() *Schema {
	cfg := Default()
	if err := goconfigtree.Load(cfg); err != nil {
		log.Printf("Warning: Failed to load local config overrides: %v\n", err)
	}

	return cfg
}
