package api

import (
	"context"
	"errors"
	"net"
	"time"

	"ports_api/apps/port_domain_service/repo"
	"ports_api/apps/port_domain_service/repo/pgdb"
	"ports_api/internal/pb"

	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const forceStopTimeout = 30 * time.Second

type ServerWorker struct {
	address string
	entry   *logrus.Entry
	db      repo.Storage
	dbCfg   db.Config
}

func NewServerWorker(entry *logrus.Entry, address string, dbCfg db.Config) *ServerWorker {
	return &ServerWorker{address: address, entry: entry, dbCfg: dbCfg}
}

func (s *ServerWorker) Init() error {
	var err error
	s.db, err = pgdb.NewRepo(s.dbCfg, s.entry)
	return err
}

func (s *ServerWorker) Run(ctx uwe.Context) error {

	gServer := grpc.NewServer()
	pb.RegisterPortDomainServiceServer(gServer,
		newServer(s.entry.WithField("app_layer", "api.gRPCServer"), s.db.Clone()),
	)

	done := make(chan struct{})
	fail := make(chan struct{})

	go func() {
		s.entry.Info("start gRPC server")

		lis, err := net.Listen("tcp", s.address)
		if err != nil {
			s.entry.WithError(err).Error("failed to listen")
			fail <- struct{}{}
			return
		}
		if err := gServer.Serve(lis); err != nil {
			s.entry.WithError(err).Error("failed to serve gRPC")
			fail <- struct{}{}
			return
		}
		done <- struct{}{}
		s.entry.Info("gRPC server stopped")
	}()

	select {
	case <-ctx.Done():
	case <-fail:
		return errors.New("gRPC PortDomainServiceServer failed")
	}

	gServer.GracefulStop()

	forceStop, cancel := context.WithTimeout(context.Background(), forceStopTimeout)
	select {
	case <-done:
		cancel()
		return nil
	case <-forceStop.Done():
		cancel()
		return errors.New("gRPC PortDomainServiceServer has not been gracefully stopped")
	}
}
