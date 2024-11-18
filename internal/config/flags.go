package config

import (
	"flag"
)

type flags struct {
	AddressOverride              string
	DatabaseDsnOverride          string
	AccuralSystemAddressOverride string
}

func parseFlags() flags {
	flgs := flags{}

	flag.StringVar(&flgs.AddressOverride, "a", addressDefaultVal, "Address of server")

	flag.StringVar(&flgs.DatabaseDsnOverride, "d", "", "Database DSN")

	flag.StringVar(&flgs.AccuralSystemAddressOverride, "r", "", "Address of Accural system")

	flag.Parse()

	return flgs
}

func (f flags) override(cfg *Config) {
	cfg.overrideAddressIfNotDefault(f.AddressOverride, "flags")
	cfg.overrideDatabaseDNSIfNotDefault(f.DatabaseDsnOverride, "flags")
	cfg.overrideAccuralSystemEndpointIfNotDefault(f.AccuralSystemAddressOverride, "flags")
}
