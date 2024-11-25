package config

import (
	"flag"
	"time"
)

type flags struct {
	AddressOverride              string
	DatabaseDsnOverride          string
	AccrualSystemAddressOverride string
	AccrualSyncWorkersPoolSize   int
	AccrualSyncIntervalSeconds   int
}

func parseFlags() flags {
	flgs := flags{}

	flag.StringVar(&flgs.AddressOverride, "a", addressDefaultVal, "Address of server")

	flag.StringVar(&flgs.DatabaseDsnOverride, "d", "", "Database DSN")

	flag.StringVar(&flgs.AccrualSystemAddressOverride, "r", "", "Address of accrual system")

	flag.IntVar(&flgs.AccrualSyncWorkersPoolSize, "p", accrualSyncWorkersPoolSizeDefaultVal, "Accrual sync workers pool size")

	flag.IntVar(&flgs.AccrualSyncIntervalSeconds, "i", int(accrualSyncIntervalDefaultVal.Seconds()), "Accrual sync interval")

	flag.Parse()

	return flgs
}

func (f flags) override(cfg *Config) {
	cfg.overrideAddressIfNotDefault(f.AddressOverride, "flags")
	cfg.overrideDatabaseDNSIfNotDefault(f.DatabaseDsnOverride, "flags")
	cfg.overrideAccrualSystemEndpointIfNotDefault(f.AccrualSystemAddressOverride, "flags")
	cfg.overrideAccrualSyncWorkersPoolSizeIfNotDefault(f.AccrualSyncWorkersPoolSize, "flags")
	cfg.overrideAccrualSyncIntervalIfNotDefault(time.Duration(f.AccrualSyncIntervalSeconds)*time.Second, "flags")
}
