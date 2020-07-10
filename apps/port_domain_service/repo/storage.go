package repo

import (
	"ports_api/apps/port_domain_service/models"

	"github.com/lancer-kit/armory/db"
)

type Storage interface {
	db.Transactional
	Clone() Storage
	Ports() PortsTable
}

type PortsTable interface {
	InsertOrUpdate(unlocode string, port models.Port) error

	Get(unlocode string) (*models.Port, error)
	Select() ([]models.Port, error)
	SelectWithPage(page *db.PageQuery) ([]models.Port, int64, error)
}
