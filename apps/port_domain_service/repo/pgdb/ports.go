package pgdb

import (
	"strings"

	"ports_api/apps/port_domain_service/models"
	"ports_api/apps/port_domain_service/repo"

	"github.com/lancer-kit/armory/db"
)

const tablePorts = "ports"

type PortsPGQ struct {
	parent *PGStorage
	table  db.Table
}

func (pgRepo *PGStorage) Ports() repo.PortsTable {
	return &PortsPGQ{
		parent: pgRepo,
		table:  db.NewTable(tablePorts, "p", "*"),
	}
}

func (pgq *PortsPGQ) InsertOrUpdate(unlocode string, port models.Port) error {
	query := pgq.table.IQBuilder.SetMap(port.ToDBMap())

	err := pgq.parent.Exec(query)
	if err != nil && strings.Contains(err.Error(), "violates unique constraint") {
		query := pgq.table.UQBuilder.SetMap(port.ToDBMap()).Where("unlocode = ? ", unlocode)
		err = pgq.parent.Exec(query)
	}

	return err
}

func (pgq *PortsPGQ) Get(unlocode string) (*models.Port, error) {
	val := new(models.Port)
	err := pgq.parent.Get(pgq.table.QBuilder.Where("unlocode = ?", unlocode), val)
	return val, err
}

func (pgq *PortsPGQ) Select() ([]models.Port, error) {
	rows := make([]models.Port, 0)

	err := pgq.parent.Select(pgq.table.QBuilder, &rows)
	return rows, err
}

func (pgq *PortsPGQ) SelectWithPage(page *db.PageQuery) ([]models.Port, int64, error) {
	rows := make([]models.Port, 0)

	count, err := pgq.table.SelectWithCount(pgq.parent.SQLConn, &rows, page.OrderBy, page)
	return rows, count, err
}
