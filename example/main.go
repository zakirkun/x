package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zakirkun/x/config"
	"github.com/zakirkun/x/config/yaml"
	"github.com/zakirkun/x/server"
)

var cfg *config.Config

func init() {
	cfg = yaml.NewYamlConfig("./config.yml")
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/say/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte("Hello " + name))
	})

	opts := server.ServerOptions{
		Handler: r,
		Host:    cfg.Server.Host,
		Port:    cfg.Server.Port,
	}

	srv := server.NewServer(opts)
	srv.Run()
}
