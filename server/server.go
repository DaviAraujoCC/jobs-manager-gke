package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hurbcom/jobs-manager-gke/config"
	"github.com/hurbcom/jobs-manager-gke/server/router"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type Server interface {
	Listen()
	Shutdown()
}

type server struct {
	router       *mux.Router
	httpServer   *http.Server
	shutdownChan chan struct{}
}

func NewServer(config *config.Config) Server {

	s := &server{
		router:       router.NewRouter(),
		shutdownChan: make(chan struct{}, 1),
	}

	setupServer(s, config)

	return s
}

func setupServer(s *server, config *config.Config) {

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false

	neg := negroni.New(
		recovery,
		negroni.NewLogger(),
	)

	s.router.Use(mux.CORSMethodMiddleware(s.router))
	neg.UseHandler(s.router)

	s.httpServer = &http.Server{
		Addr:    ":" + config.HTTPPort,
		Handler: neg,
	}

}

func (s server) Listen() {

	ctx := context.Background()

	go func() {
		<-s.shutdownChan
		logrus.Warn("shutting down server")

		logrus.Warn("exiting...")
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logrus.Fatalf("Error while shutting down server: %v", err.Error())
		}
	}()

	go func() {

		if err := s.httpServer.ListenAndServe(); err != nil {
			logrus.WithError(err).Warn("server shutdown")
			s.Shutdown()
		}

	}()

}

func (s *server) Shutdown() {
	s.shutdownChan <- struct{}{}
}
