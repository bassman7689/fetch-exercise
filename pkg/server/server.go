package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/bassman7689/fetch-exercise/pkg/config"
	"github.com/bassman7689/fetch-exercise/pkg/controllers"
	"github.com/bassman7689/fetch-exercise/pkg/store"
)

type Server struct {
	conf *config.Config
	store store.Store
}

func New(conf *config.Config) (*Server, error) {
	store := store.NewMemoryStore()
	return &Server{conf: conf, store: store}, nil
}

func (s *Server) Run() error {
	r := mux.NewRouter()
	controllers.Register(r, s.store)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Handler:      r,
		Addr:         s.conf.ListenAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go srv.ListenAndServe()
	log.Printf("Listening on %s", s.conf.ListenAddr)

	<- c
	return nil
}
