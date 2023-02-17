package v1

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	logger *log.Logger
	router *mux.Router
}

func NewServer(logger *logrus.Logger) *Server {
	return &Server{
		logger: logger,
		router: mux.NewRouter(),
	}
}

func (s *Server) Init() error {
	s.logger.Info("Initializing the server")
	return nil
}

func (s *Server) Run(clientServerAddr string) error {
	s.router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Index is working fine")
	})

	s.logger.Info("Starting the server " + clientServerAddr)
	http.ListenAndServe(clientServerAddr, s.router)

	return nil
}
