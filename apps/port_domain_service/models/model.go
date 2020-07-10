package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"ports_api/internal/pb"
)

type Port struct {
	Unlocode string    `db:"unlocode"`
	Name     string    `db:"name"`
	City     string    `db:"city"`
	Country  string    `db:"country"`
	Timezone string    `db:"timezone"`
	Code     string    `db:"code"`
	Extra    PortExtra `db:"extra"`

	CreatedAt int64 `db:"created_at"`
	UpdatedAt int64 `db:"updated_at"`
}

func (port *Port) ToDBMap() map[string]interface{} {
	return map[string]interface{}{
		"unlocode":   port.Unlocode,
		"name":       port.Name,
		"city":       port.City,
		"country":    port.Country,
		"timezone":   port.Timezone,
		"code":       port.Code,
		"extra":      port.Extra,
		"updated_at": time.Now().UTC().Unix(),
	}
}

func (port *Port) ToProto() *pb.Port {
	return &pb.Port{
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Extra.Alias,
		Regions:     port.Extra.Regions,
		Coordinates: port.Extra.Coordinates,
		Province:    port.Extra.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Extra.Unlocs,
		Code:        port.Code,
	}
}

func (Port) FromProto(unlocode string, port *pb.Port) Port {
	return Port{
		Unlocode: unlocode,
		Name:     port.Name,
		City:     port.City,
		Country:  port.Country,
		Timezone: port.Timezone,
		Code:     port.Code,
		Extra: PortExtra{
			Alias:       port.Alias,
			Regions:     port.Regions,
			Coordinates: port.Coordinates,
			Province:    port.Province,
			Unlocs:      port.Unlocs,
		},
	}
}

type PortExtra struct {
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float32 `json:"coordinates"`
	Province    string    `json:"province"`
	Unlocs      []string  `json:"unlocs"`
}

func (p *PortExtra) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p)
}

func (p PortExtra) Value() (driver.Value, error) {
	return json.Marshal(p)
}
