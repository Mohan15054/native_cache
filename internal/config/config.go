package config

import (
	"github.com/spf13/viper"
)

type CacheConfig struct {
	MaxMemory      int
	EvictionPolicy string
	DefaultTTL     int
	MetricsEnabled bool
	ListenAddr     string
}

func Load() *CacheConfig {
	return &CacheConfig{
		MaxMemory:      viper.GetInt("cache.max_memory"),
		EvictionPolicy: viper.GetString("cache.eviction_policy"),
		DefaultTTL:     viper.GetInt("cache.default_ttl"),
		MetricsEnabled: viper.GetBool("cache.metrics_enabled"),
		ListenAddr:     viper.GetString("cache.listen_addr"),
	}
}
