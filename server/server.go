package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zakirkun/x/logger"
)

func NewServer(opt ServerOptions) IServer {
	return ServerOptions{
		Handler:      opt.Handler,
		WriteTimeout: opt.WriteTimeout,
		ReadTimeout:  opt.ReadTimeout,
		IdleTimeout:  opt.IdleTimeout,
		Host:         opt.Host,
		Port:         opt.Port,
	}
}

func (s ServerOptions) Run() {
	// Set up a channel to listen to for interrupt signals
	var runningChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(context.Background(), s.WriteTimeout)

	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         s.Host + ":" + s.Port,
		Handler:      s.Handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runningChan, os.Interrupt, syscall.SIGTERM)

	// Alert the user that the server is starting
	logger.Info(fmt.Sprintf("Server is starting on %s\n", server.Addr))

	// Run the server on a new goroutine
	go func() {

		//  Running Server Without SSL
		// TODO Adding SSL
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// todo logger error
			} else {
				logger.Warn(fmt.Sprintf("Server failed to start due to err: %v", err))
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runningChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	logger.Error(fmt.Sprintf("Server is shutting down due to %+v\n", interrupt))

	// shutdown server
	if err := server.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server was unable to gracefully shutdown due to err: %+v", err))
	}
}
