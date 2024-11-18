package config

import (
	"github.com/aridae/gophermart-diploma/internal/logger"
	"sync"
)

const (
	addressDefaultVal          = "localhost:8080"
	databaseMaxOpenConnDefault = 5
	adminAddressDefaultVal     = "localhost:8004"
)

var (
	once         sync.Once
	globalConfig *Config
)

type Config struct {
	Address             string
	DatabaseDsn         string
	DatabaseMaxOpenConn int

	AccuralSystemAddress string

	JWTKey string

	AdminAddress string
}

func Obtain() *Config {
	once.Do(func() {
		globalConfig = &Config{}
		globalConfig.init()
	})

	return globalConfig
}

func (c *Config) init() {
	c.defaults()

	// перезатираем значениями, переданными через флаги
	parseFlags().override(c)

	// env, если есть, затирает флаги
	envValues, err := readEnv()
	if err != nil {
		logger.Errorf("error parsing environment, proceeding without env overrides: %v", err)
	} else {
		envValues.override(c)
	}
}

func (c *Config) defaults() {
	c.Address = addressDefaultVal
	c.DatabaseMaxOpenConn = databaseMaxOpenConnDefault
	c.AdminAddress = adminAddressDefaultVal
}

func (c *Config) overrideAddressIfNotDefault(address string, source string) {
	if address == addressDefaultVal {
		logger.Debugf("source %s provided default Address value, not overriding", source)
		return
	}

	logger.Infof("overriding Address from %s: (%s)-->(%s)", source, c.Address, address)
	c.Address = address
}

func (c *Config) overrideDatabaseDNSIfNotDefault(dns string, source string) {
	if dns == "" {
		logger.Debugf("source %s provided empty dns value, not overriding", source)
		return
	}

	logger.Infof("overriding dns from %s", source)
	c.DatabaseDsn = dns
}

func (c *Config) overrideAccuralSystemEndpointIfNotDefault(addr string, source string) {
	if addr == "" {
		logger.Debugf("source %s provided empty accural system address value, not overriding", source)
		return
	}

	logger.Infof("overriding accural system address from %s", source)
	c.AccuralSystemAddress = addr
}

func (c *Config) overrideJWTKeyIfNotDefault(key string, source string) {
	if key == "" {
		logger.Debugf("source %s provided empty JWT key, not overriding", source)
		return
	}

	logger.Infof("overriding JWT key from %s", source)
	c.JWTKey = key
}
