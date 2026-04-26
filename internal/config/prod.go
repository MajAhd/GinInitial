package config

import (
	"log"

	goconfigtree "github.com/MajAhd/go-config-tree"
)

func Prod() *Schema {
	cfg := &Schema{}
	if err := goconfigtree.Load(cfg); err != nil {
		log.Fatalf("CRITICAL: Failed to load prod config: %v\n", err)
	}

	return cfg
}
