package daemons

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"ports_api/apps/client_api/types"
	"ports_api/internal/pb"

	"github.com/lancer-kit/uwe/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Daemon struct {
	log              *logrus.Entry
	portsServiceAddr string
	dataFile         string
}

func NewDaemon(log *logrus.Entry, dataFile, portsServiceAddr string) *Daemon {
	return &Daemon{
		log:              log.WithField("app_layer", "daemons.Daemon"),
		dataFile:         dataFile,
		portsServiceAddr: portsServiceAddr,
	}
}

func (h *Daemon) Init() error { return nil }

func (h *Daemon) Run(ctx uwe.Context) error {
	client, err := h.portClient(ctx)
	if err != nil {
		return err
	}

	outBus := make(chan portData)
	errBus := make(chan error)

	go func() {
		h.log.WithField("data_file", h.dataFile).Info("start reading datafile")
		var file *os.File
		file, err = os.Open(h.dataFile)
		if err != nil {
			const msg = "unable to open data file"
			h.log.WithError(err).Error(msg)
			errBus <- errors.Wrap(err, msg)
			return
		}
		defer file.Close()

		err = readPortsStream(file, outBus)
		if err != nil {
			const msg = "data stream reading failed "
			h.log.WithError(err).Error(msg)
			errBus <- errors.Wrap(err, msg)
			return
		}
		h.log.WithField("data_file", h.dataFile).Info("finish reading datafile")
	}()

	for {
		select {
		case <-ctx.Done():
			return err
		case e := <-errBus:
			return e
		case val := <-outBus:
			_, err := client.SavePort(ctx, &pb.NewPort{
				Unlocode: val.Unlocode,
				Port:     val.Port.ToProto(),
			})
			if err != nil {
				const msg = "SavePort request failed "
				h.log.WithError(err).Error(msg)
			}

			h.log.WithField("unlocode", val.Unlocode).Debug("save new port")
		}
	}
}

type portData struct {
	Unlocode string
	Port     types.Port
}

func readPortsStream(reader io.Reader, outBus chan<- portData) error {
	dec := json.NewDecoder(reader)

	// read open bracket
	if _, err := dec.Token(); err != nil {
		return errors.Wrap(err, "unable to read open bracket")
	}

	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return errors.Wrap(err, "unable to read port key")
		}
		unlocode, ok := key.(string)
		if !ok {
			continue
		}

		var portRaw json.RawMessage
		err = dec.Decode(&portRaw)
		if err != nil {
			return errors.Wrap(err, "unable to decode raw port obj")
		}

		var port types.Port
		err = port.UnmarshalJSON(portRaw)
		if err != nil {
			return errors.Wrap(err, "unable to unmarshal port obj")
		}

		outBus <- portData{
			Unlocode: unlocode,
			Port:     port,
		}
	}

	if _, err := dec.Token(); err != nil {
		return errors.Wrap(err, "unable to read closing bracket")
	}

	return nil
}

func (h *Daemon) portClient(ctx context.Context) (pb.PortDomainServiceClient, error) {
	conn, err := grpc.DialContext(ctx, h.portsServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewPortDomainServiceClient(conn), nil
}
