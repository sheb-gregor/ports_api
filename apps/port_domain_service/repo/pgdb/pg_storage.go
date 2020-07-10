package pgdb

import (
	"ports_api/apps/port_domain_service/repo"

	"github.com/lancer-kit/armory/db"
	"github.com/sirupsen/logrus"
)

// repo implementation of the `Repo` interface.
type PGStorage struct {
	*db.SQLConn
	stored *db.Config
}

func (pgRepo *PGStorage) Clone() repo.Storage {
	return &PGStorage{
		SQLConn: pgRepo.SQLConn.Clone(),
		stored:  pgRepo.stored,
	}
}

// NewRepo returns initialized instance of the `Repo`.
func NewRepo(config db.Config, logger *logrus.Entry) (repo.Storage, error) {
	sqlConn, err := db.NewConnector(config, logger.WithField("app_layer", "repo.PGStorage"))
	if err != nil {
		return nil, err
	}

	return &PGStorage{SQLConn: sqlConn, stored: &config}, nil
}
