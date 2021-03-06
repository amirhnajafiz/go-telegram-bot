package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	srv     *http.ServeMux
	address string
}

// NewServer creates a new monitoring server.
func NewServer(cfg Config) Server {
	var srv *http.ServeMux

	srv = http.NewServeMux()
	srv.Handle("/metrics", promhttp.Handler())

	return Server{
		address: cfg.Host,
		srv:     srv,
	}
}

// Start creates and run a metric server for prometheus in new go routine.
func (s Server) Start() {
	go func() {
		if err := http.ListenAndServe(s.address, s.srv); err != nil {
			panic(err)
		}
	}()
}
