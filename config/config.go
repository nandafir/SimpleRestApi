package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var (
	once sync.Once
	conf *Config
)

// Config ...
type Config struct {
	Port       int  `envconfig:"PORT" default:"3001"`
	RetryCount int  `envconfig:"RETRY_COUNT" default:"3"`
	DebugMode  bool `envconfig:"DEBUG_MODE" default:"true"`
}

// Init ...
func Init() {
	conf = new(Config)
	envconfig.MustProcess("", conf)
}

// Get ...
func Get() *Config {
	return conf
}
