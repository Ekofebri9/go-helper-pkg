package config

import "github.com/spf13/viper"

type Option func(*Loader)

func WithConfigFile(path string) Option {
	return func(l *Loader) {
		l.v.SetConfigFile(path)
	}
}

func WithConfigName(name string) Option {
	return func(l *Loader) {
		l.v.SetConfigName(name)
	}
}

func WithConfigPath(path string) Option {
	return func(l *Loader) {
		l.v.AddConfigPath(path)
	}
}

func WithDefaults(values map[string]any) Option {
	return func(l *Loader) {
		for k, v := range values {
			l.v.SetDefault(k, v)
		}
	}
}

func WithViper(fn func(*viper.Viper)) Option {
	return func(l *Loader) {
		fn(l.v)
	}
}
