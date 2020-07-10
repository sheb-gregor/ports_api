package config

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/noble"
)

// DBConfig configuration with secrets support
type DBConfig struct {
	Driver      string               `yaml:"driver"`
	Name        string               `yaml:"name"`
	Host        string               `yaml:"host"`
	Port        int                  `yaml:"port" `
	User        noble.Secret         `yaml:"user"`
	Password    noble.Secret         `yaml:"password"`
	SSL         bool                 `yaml:"ssl"`
	InitTimeout int                  `yaml:"init_timeout"`
	AutoMigrate bool                 `yaml:"auto_migrate"`
	WaitForDB   bool                 `yaml:"wait_for_db"`
	Params      *db.ConnectionParams `yaml:"params"`
}

// ConnectionString returns Connection String for selected driver
func (d DBConfig) ConnectionString() string {
	port := ""
	if d.Port != 0 {
		port = fmt.Sprintf(":%d", d.Port)
	}

	if d.Driver == "postgres" {
		mode := ""
		if !d.SSL {
			mode = "?sslmode=disable"
		}
		DSN := `postgres://%s:%s@%s%s/%s%s`
		return fmt.Sprintf(DSN, d.User.Get(), d.Password.Get(), d.Host, port, d.Name, mode)
	}

	return ""
}

// Config returns lancer db Config
func (d DBConfig) Config() db.Config {
	return db.Config{
		ConnURL:     d.ConnectionString(),
		InitTimeout: d.InitTimeout,
		AutoMigrate: d.AutoMigrate,
		WaitForDB:   d.WaitForDB,
		Params:      d.Params,
	}
}

// Validate config
func (d DBConfig) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Driver, validation.Required),
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Host, validation.Required),
		validation.Field(&d.InitTimeout, validation.Required),
		validation.Field(&d.User, validation.Required, noble.RequiredSecret),
		validation.Field(&d.Password, validation.Required, noble.RequiredSecret),
	)
}
