package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type environs struct {
	AddressOverride              *string `env:"RUN_ADDRESS"`
	DatabaseDsnOverride          *string `env:"DATABASE_URI"`
	AccuralSystemAddressOverride *string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	JWTSecretKey                 *string `env:"JWT_SECRET_KEY"`
}

func readEnv() (environs, error) {
	envs := environs{}

	err := env.Parse(&envs)
	if err != nil {
		return environs{}, fmt.Errorf("failed to parse env variables: %w", err)
	}

	return envs, nil
}

func (e environs) override(cfg *Config) {
	if e.AddressOverride != nil {
		cfg.overrideAddressIfNotDefault(*e.AddressOverride, "env")
	}

	if e.DatabaseDsnOverride != nil {
		cfg.overrideDatabaseDNSIfNotDefault(*e.DatabaseDsnOverride, "env")
	}

	if e.AccuralSystemAddressOverride != nil {
		cfg.overrideAccuralSystemEndpointIfNotDefault(*e.AccuralSystemAddressOverride, "env")
	}

	if e.JWTSecretKey != nil {
		cfg.overrideJWTKeyIfNotDefault(*e.JWTSecretKey, "env")
	}
}
