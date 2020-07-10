package types

import "ports_api/internal/pb"

//go:generate easyjson port.go

// easyjson:json
type Port struct {
	Name        string     `json:"name"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
	Alias       []string   `json:"alias"`
	Regions     []string   `json:"regions"`
	Coordinates [2]float32 `json:"coordinates"`
	Province    string     `json:"province"`
	Timezone    string     `json:"timezone"`
	Unlocs      []string   `json:"unlocs"`
	Code        string     `json:"code"`
}

func (v Port) ToProto() *pb.Port {
	return &pb.Port{
		Name:        v.Name,
		City:        v.City,
		Country:     v.Country,
		Alias:       v.Alias,
		Regions:     v.Regions,
		Coordinates: v.Coordinates[:],
		Province:    v.Province,
		Timezone:    v.Timezone,
		Unlocs:      v.Unlocs,
		Code:        v.Code,
	}

}
