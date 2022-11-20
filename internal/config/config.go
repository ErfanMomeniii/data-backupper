package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"strings"
)

var (
	Default = "config.yaml"
	C       *Config
)

type Config struct {
	Logger  Logger  `yaml:"logger"`
	Tracing Tracing `yaml:"tracing"`
}

// Tracing config struct
type Tracing struct {
	Enabled      bool    `yaml:"enabled"`
	AgentHost    string  `yaml:"agent_host"`
	AgentPort    string  `yaml:"agent_port"`
	SamplerRatio float64 `yaml:"sampler_ratio"`
}

// Logger config values
type Logger struct {
	Level string `yaml:"level"`
}

func Init() *Config {
	c := new(Config)
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	v.SetConfigName(Default)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error loading configs: %s", err.Error()).(any))
	}

	err := v.Unmarshal(c, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		config.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		)
	})
	if err != nil {
		panic(fmt.Errorf("failed on config `%s` unmarshal: %w", Default, err).(any))
	}

	C = c

	return c
}
