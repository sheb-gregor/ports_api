package config

import (
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lancer-kit/armory/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Configuration interface {
	Init()
	LogEntry() *logrus.Entry
	Validate() error
}

type Cfg struct {
	Log LogConfig `yaml:"log"`

	entry *logrus.Entry
}

func (cfg *Cfg) LogEntry() *logrus.Entry {
	return cfg.entry
}

func (cfg *Cfg) Init() {
	entry, err := cfg.InitLog()
	if err != nil {
		logrus.New().
			WithError(err).
			Fatal("Unable to init log")
	}
	cfg.entry = entry
}

func (cfg Cfg) InitLog() (*logrus.Entry, error) {
	return log.Init(log.Config{
		AppName: cfg.Log.AppName,
		Level:   cfg.Log.Level.Get(),
		JSON:    cfg.Log.JSON,
	})
}

func (cfg Cfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Log, validation.Required),
	)
}

func ReadConfig(path string, config Configuration) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.New().
			WithError(err).
			WithField("path", path).
			Fatal("unable to read config file")
	}

	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		logrus.New().
			WithError(err).
			WithField("raw_config", string(rawConfig)).
			Fatal("unable to unmarshal config file")
	}

	err = config.Validate()
	if err != nil {
		logrus.New().
			WithError(err).
			Fatal("Invalid configuration")
	}

	config.Init()
}
