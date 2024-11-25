package config

import (
	"sync"
	"time"

	"github.com/aridae/gophermart-diploma/internal/logger"
)

const (
	addressDefaultVal                    = "localhost:8080"
	databaseMaxOpenConnDefault           = 5
	adminAddressDefaultVal               = "localhost:8004"
	accrualSyncWorkersPoolSizeDefaultVal = 100
	accrualSyncIntervalDefaultVal        = time.Second
)

var (
	once         sync.Once
	globalConfig *Config
)

type Config struct {
	Address             string
	DatabaseDsn         string
	DatabaseMaxOpenConn int

	AccrualSystemAddress string

	JWTKey string

	AdminAddress string

	AccrualSyncWorkersPoolSize int
	AccrualSyncInterval        time.Duration
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
	c.AccrualSyncWorkersPoolSize = accrualSyncWorkersPoolSizeDefaultVal
	c.AccrualSyncInterval = accrualSyncIntervalDefaultVal
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

func (c *Config) overrideAccrualSystemEndpointIfNotDefault(addr string, source string) {
	if addr == "" {
		logger.Debugf("source %s provided empty accrual system address value, not overriding", source)
		return
	}

	logger.Infof("overriding accrual system address from %s", source)
	c.AccrualSystemAddress = addr
}

func (c *Config) overrideJWTKeyIfNotDefault(key string, source string) {
	if key == "" {
		logger.Debugf("source %s provided empty JWT key, not overriding", source)
		return
	}

	logger.Infof("overriding JWT key from %s", source)
	c.JWTKey = key
}

func (c *Config) overrideAccrualSyncWorkersPoolSizeIfNotDefault(val int, source string) {
	if val == 0 {
		logger.Debugf("source %s provided zero workers pool size, not overriding", source)
		return
	}

	if val == accrualSyncWorkersPoolSizeDefaultVal {
		logger.Debugf("source %s provided default workers pool size, not overriding", source)
		return
	}

	logger.Infof("overriding workers pool size from %s: (%d)-->(%d)", source, c.AccrualSyncWorkersPoolSize, val)
	c.AccrualSyncWorkersPoolSize = val
}

func (c *Config) overrideAccrualSyncIntervalIfNotDefault(val time.Duration, source string) {
	if val == time.Duration(0) {
		logger.Debugf("source %s provided zero accrual sync interval, not overriding", source)
		return
	}

	if val == accrualSyncIntervalDefaultVal {
		logger.Debugf("source %s provided default accrual sync interval, not overriding", source)
		return
	}

	logger.Infof("overriding accrual sync interval from %s: (%s)-->(%s)", source, c.AccrualSyncInterval, val)
	c.AccrualSyncInterval = val
}
