 package main

 import (
"fmt"
"net/http"
"os"

log "github.com/sirupsen/logrus"
"github.com/gorilla/mux"

)

type Server struct {
	logger *log.Logger
	router *mux.Router
}

func (s *Server) Init() error {
	s.logger.Info("Initializing the server")
	return nil
}

func (s *Server) Run() error {
	s.logger.Info("Starting the server")
	s.router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Index is working fine")
	})
	http.ListenAndServe(":8080", s.router)
	return nil
}

func main() {
	logger:= log.New()
	logger.Trace("Logger initialized")
	server := Server{
		logger: logger,
		router: mux.NewRouter(),
	}
	err := server.Init()
	if err != nil {
		logger.Error("Couldn't initialize the server", err)
		os.Exit(-1)
	}
	err = server.Run()
	if err != nil {
		logger.Error("Couldn't start the server", err)
		os.Exit(-1)
	}
}

