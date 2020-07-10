package api

import (
	"context"

	"ports_api/apps/port_domain_service/models"
	"ports_api/apps/port_domain_service/repo"
	"ports_api/internal/pb"

	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/db"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	entry *logrus.Entry
	db    repo.Storage
}

func newServer(entry *logrus.Entry, db repo.Storage) *server {
	return &server{entry: entry, db: db}
}

func (s *server) GetPortsList(_ context.Context, params *pb.PageParams) (*pb.PortList, error) {
	s.entry.WithField("method", "GetPortsList").Info("got new call")

	ports, count, err := s.db.Ports().SelectWithPage(&db.PageQuery{
		Order:    params.Order,
		Page:     params.Page,
		PageSize: params.PageSize,
		OrderBy:  params.OrderBy,
	})
	if err != nil {
		const msg = "unable to fetch ports list from database"
		s.entry.WithError(err).Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}

	page := render.Page{}
	page.SetTotal(uint64(count), params.PageSize)

	var records = make([]*pb.Port, len(ports))
	for i, port := range ports {
		records[i] = port.ToProto()
	}

	return &pb.PortList{
		Page:     params.Page,
		PageSize: params.PageSize,
		Order:    params.Order,
		OrderBy:  params.OrderBy,
		Total:    page.Total,
		Records:  records,
	}, nil
}

func (s *server) GetPort(_ context.Context, request *pb.GetPortRequest) (*pb.Port, error) {
	s.entry.WithField("method", "GetPort").Info("got new call")

	port, err := s.db.Ports().Get(request.Unlocode)
	if err != nil {
		const msg = "unable to fetch port from database"
		s.entry.WithField("unlocode", request.Unlocode).WithError(err).Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}

	return port.ToProto(), nil
}

func (s *server) SavePort(_ context.Context, request *pb.NewPort) (*pb.Result, error) {
	s.entry.WithField("method", "SavePort").Info("got new call")

	if request.Unlocode == "" || request.Port == nil {
		return &pb.Result{Ok: false}, status.Errorf(codes.InvalidArgument, "parameter NewPort has invalid value")
	}

	err := s.db.Ports().InsertOrUpdate(
		request.Unlocode, models.Port{}.FromProto(request.Unlocode, request.Port))
	if err != nil {
		const msg = "unable to store port to database"
		s.entry.WithField("unlocode", request.Unlocode).WithError(err).Error(msg)
		return &pb.Result{Ok: false}, status.Errorf(codes.Internal, msg)
	}

	return &pb.Result{Ok: true}, nil
}
