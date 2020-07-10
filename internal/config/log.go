package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lancer-kit/noble"
)

// Config is a options for the initialization
// of the default logrus.Entry.
type LogConfig struct {
	// AppName identifier of the app.
	AppName string `yaml:"app_name"`
	// Level is a string representation of the `lorgus.Level`.
	Level noble.Secret `yaml:"level"`
	// Sentry is a DSN string for sentry hook.
	Sentry string `yaml:"sentry"`
	// AddTrace enable adding of the filename field into log.
	AddTrace bool `yaml:"add_trace"`
	// JSON enable json formatted output.
	JSON bool `yaml:"json"`
}

func (cfg LogConfig) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.AppName, validation.Required),
		validation.Field(&cfg.Level, validation.Required, noble.RequiredSecret),
	)
}
