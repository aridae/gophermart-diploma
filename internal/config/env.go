package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type environs struct {
	AddressOverride              *string `env:"RUN_ADDRESS"`
	DatabaseDsnOverride          *string `env:"DATABASE_URI"`
	AccrualSystemAddressOverride *string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	JWTSecretKey                 *string `env:"JWT_SECRET_KEY"`
	AccrualSyncWorkersPoolSize   *int    `env:"ACCRUAL_SYNC_WORKERS_COUNT"`
	AccrualSyncIntervalSeconds   *int    `env:"ACCRUAL_SYNC_INTERVAL_SEC"`
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

	if e.AccrualSystemAddressOverride != nil {
		cfg.overrideAccrualSystemEndpointIfNotDefault(*e.AccrualSystemAddressOverride, "env")
	}

	if e.JWTSecretKey != nil {
		cfg.overrideJWTKeyIfNotDefault(*e.JWTSecretKey, "env")
	}

	if e.AccrualSyncWorkersPoolSize != nil {
		cfg.overrideAccrualSyncWorkersPoolSizeIfNotDefault(*e.AccrualSyncWorkersPoolSize, "env")
	}

	if e.AccrualSyncIntervalSeconds != nil {
		cfg.overrideAccrualSyncIntervalIfNotDefault(time.Second*time.Duration(*e.AccrualSyncIntervalSeconds), "env")
	}
}
