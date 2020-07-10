package main

import (
	"flag"

	"ports_api/apps/port_domain_service/api"
	"ports_api/apps/port_domain_service/migrations"
	"ports_api/internal/config"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
)

type Config struct {
	config.Cfg    `yaml:",inline"`
	DB            config.DBConfig `yaml:"db"`
	ServerAddress string          `yaml:"server_address"`
}

func (cfg Config) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Cfg, validation.Required),
		validation.Field(&cfg.DB, validation.Required),
		validation.Field(&cfg.ServerAddress, validation.Required),
	)
}

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg := new(Config)
	config.ReadConfig(*cfgPath, cfg)
	log := cfg.LogEntry()
	ensureMigrations(log, cfg.DB.ConnectionString())

	chief := uwe.NewChief()
	chief.AddWorker("grpc_server",
		api.NewServerWorker(log, cfg.ServerAddress, cfg.DB.Config()))

	chief.UseDefaultRecover()
	chief.EnableServiceSocket(config.AppInfo())
	chief.SetEventHandler(uwe.LogrusEventHandler(log))
	log.Info("start chief")
	chief.Run()
}

func ensureMigrations(entry *logrus.Entry, dbConn string) {
	count, err := migrations.Migrate(dbConn, db.MigrateUp)
	if err != nil {
		entry.WithError(err).Fatal("unable to apply migrations")
		return
	}

	if count > 0 {
		entry.WithField("count", count).Info("up migrations applied")
	}
}
