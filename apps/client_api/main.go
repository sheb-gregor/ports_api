package main

import (
	"flag"

	api2 "ports_api/apps/client_api/api"
	"ports_api/apps/client_api/daemons"
	"ports_api/internal/config"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lancer-kit/uwe/v2"
	"github.com/lancer-kit/uwe/v2/presets/api"
)

type Config struct {
	config.Cfg   `yaml:",inline"`
	DataFile     string     `yaml:"data_file"`
	API          api.Config `yaml:"api"`
	PortsService string     `yaml:"ports_service"`
}

func (cfg Config) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Cfg, validation.Required),
		validation.Field(&cfg.API, validation.Required),
		validation.Field(&cfg.DataFile, validation.Required),
	)
}

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()
	println("hhh")
	cfg := new(Config)
	config.ReadConfig(*cfgPath, cfg)

	log := cfg.LogEntry()
	chief := uwe.NewChief()
	chief.AddWorker("http_server", api2.NewServer(log, cfg.API, cfg.PortsService))
	chief.AddWorker("data_reader", daemons.NewDaemon(log, cfg.DataFile, cfg.PortsService))

	chief.UseDefaultRecover()
	chief.SetEventHandler(uwe.LogrusEventHandler(log))
	log.Info("start chief")
	chief.Run()
}
