package config

import (
	"context"
	"log"
	"sync"
)

//Config is a map which contains string/string key value configuration pairs
type Config struct {
	configurationMap map[string]string
	mux              sync.Mutex
}

//ConfigSetting defines operations served on a configuration map
type ConfigSetting interface {
	GetConfigValue(ctx context.Context, key string) string
	SetConfigValue(ctx context.Context, key, value string)
	LoadConfigMap(ctx context.Context, valset map[string]string)
	DumpConfigMap(ctx context.Context)
}

//LoadConfigMap loads a map[string]string of config values
func (bc *Config) LoadConfigMap(ctx context.Context, valset map[string]string) {
	for item, val := range valset {
		i, v := item, val
		bc.SetConfigValue(ctx, i, v)
	}
}

//NewConfig returns a new base config
func NewConfig(ctx context.Context) *Config {
	cfm := make(map[string]string)
	bc := &Config{configurationMap: cfm}

	return bc
}

//GetConfigValue returns a config value from the map
func (bc *Config) GetConfigValue(ctx context.Context, key string) string {
	bc.mux.Lock()
	defer bc.mux.Unlock()
	return bc.configurationMap[key]
}

//SetConfigValue sets a config value in the map
func (bc *Config) SetConfigValue(ctx context.Context, key, value string) {
	bc.mux.Lock()
	bc.configurationMap[key] = value
	bc.mux.Unlock()
}

//DumpConfigMap writes out the configmap
func (bc *Config) DumpConfigMap(ctx context.Context) {
	log.Printf("%v", bc.configurationMap)
}
