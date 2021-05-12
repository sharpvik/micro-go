package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sharpvik/log-go/v2"

	"github.com/sharpvik/micro-go/configs"
	"github.com/sharpvik/micro-go/database"
)

// Server is a wrapper around http.Server that allows us to define our own
// convenience methods.
type Server struct {
	*http.Server
}

// MustInit attempts to initialise all the connections to real services (e.g.
// databases, message queues etc.) and returns new Server instance.
// Use this function for testing!
func MustInit(config *configs.Config) (s *Server) {
	_ = database.MustInit(config.Database)
	return NewServer(config.Server)
}

// NewServer accepts all the neccessary repositories (workers attached to the
// running services or mocks of such) and returns a new Server instance.
func NewServer(server *configs.Server) *Server {
	return &Server{
		&http.Server{
			Addr:         server.Address,
			Handler:      newMainHandler(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

// Grace allows Server to shutdown gracefully based on an external done chan.
func (s *Server) Grace(done chan bool) {
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Debug("stopping server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.SetKeepAlivesEnabled(false)

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("graceful server shutdown failed")
	}

	close(done)
}

// Serve starts the server.
func (s *Server) Serve() (err error) {
	log.Debugf("serving at %s ...", s.Addr)
	err = s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Errorf("server shut with error: %s", err)
	}
	return
}

// ServeWithGrace spawns the greceful shutfown monitor thread and then calls
// ListenAndServe on the server.
func (s *Server) ServeWithGrace(done chan bool) {
	go s.Grace(done)
	go s.Serve()
}
