package api

import (
	"context"
	"net/http"

	"ports_api/internal/config"
	"ports_api/internal/pb"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const unlocodeParam = "unlocode"

func NewServer(entry *logrus.Entry, cfg api.Config, portsServiceAddr string) *api.Server {
	h := newHandler(entry, portsServiceAddr)
	mux := chi.NewMux()
	mux.With(
		middleware.Recoverer,
		middleware.RealIP,
		middleware.RequestID,
		log.NewRequestLogger(entry.Logger),
	)

	mux.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		render.Success(w, config.AppInfo())
	})

	mux.Route("/ports", func(r chi.Router) {
		r.Get("/", h.getPortList)
		r.Get("/{"+unlocodeParam+"}", h.getPort)
	})

	return api.NewServer(cfg, mux)
}

type handler struct {
	log              *logrus.Entry
	portsServiceAddr string
}

func (h handler) portClient(ctx context.Context) (pb.PortDomainServiceClient, error) {
	conn, err := grpc.DialContext(ctx, h.portsServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewPortDomainServiceClient(conn), nil
}

func newHandler(log *logrus.Entry, portsServiceAddr string) *handler {
	return &handler{
		log:              log.WithField("app_layer", "api.Handler"),
		portsServiceAddr: portsServiceAddr,
	}
}

func (h handler) getPortList(w http.ResponseWriter, r *http.Request) {
	page, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		render.BadRequest(w, "invalid page params")
		return
	}
	client, err := h.portClient(r.Context())
	if err != nil {
		log.IncludeRequest(h.log, r).WithError(err).Error("can't initialize port client")
		render.ServerError(w)
		return
	}

	ports, err := client.GetPortsList(r.Context(), &pb.PageParams{
		Order:    page.Order,
		Page:     page.Page,
		PageSize: page.PageSize,
		OrderBy:  page.OrderBy,
	})
	if err != nil {
		log.IncludeRequest(h.log, r).WithError(err).Error("can't get port records")
		render.ServerError(w)
		return
	}

	render.Success(w, ports)
}

func (h handler) getPort(w http.ResponseWriter, r *http.Request) {
	unlocode := chi.URLParam(r, unlocodeParam)
	if unlocode == "" {
		render.BadRequest(w, unlocodeParam+": must not be empty")
		return
	}

	client, err := h.portClient(r.Context())
	if err != nil {
		log.IncludeRequest(h.log, r).WithError(err).Error("can't initialize port client")
		render.ServerError(w)
		return
	}

	port, err := client.GetPort(r.Context(), &pb.GetPortRequest{Unlocode: unlocode})
	if err != nil {
		log.IncludeRequest(h.log, r).WithError(err).Error("can't get port record")
		render.ServerError(w)
		return
	}

	render.Success(w, port)
}
