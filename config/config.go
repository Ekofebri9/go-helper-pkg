package config

import (
	"strings"

	"github.com/spf13/viper"
)

type configType string

const (
	YAML configType = "yaml"
	JSON configType = "json"
	TOML configType = "toml"
)

type Loader struct {
	v *viper.Viper
}

func New(configType configType, opts ...Option) *Loader {
	v := viper.New()

	// default behavior
	if configType == "" {
		configType = YAML
	}
	v.SetConfigType(string(configType))
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	l := &Loader{v: v}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (l *Loader) Read() error {
	if l.v.ConfigFileUsed() != "" {
		return l.v.ReadInConfig()
	}
	return nil
}

func (l *Loader) Unmarshal(out any) error {
	return l.v.Unmarshal(out)
}

func (l *Loader) Viper() *viper.Viper {
	return l.v
}
