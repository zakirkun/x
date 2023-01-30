package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zakirkun/x/logger"

	"github.com/zakirkun/x/cache/redis"

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

	type Users struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	optRedis := redis.RedisOpt{
		Address:  cfg.Caching.Address,
		Password: cfg.Caching.Password,
		Db:       cfg.Caching.Db,
		Expired:  cfg.Caching.ExpiredTime,
	}

	cache := redis.NewRedis(optRedis)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/say/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte("Hello " + name))
	})

	r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		data, err := cache.Get(id)
		if err == redis.KEY_NOT_EXISTS {
			logger.Warn("Key not exists")
		} else if err != nil {
			logger.Warn(fmt.Sprintf("Error get data from redis : %v", err))
		}

		var userData Users
		if err := json.Unmarshal([]byte(data), &userData); err != nil {
			logger.Warn(fmt.Sprintf("Error unmarshal data : %v", err))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		result := map[string]interface{}{
			"message": "success",
			"data":    userData,
		}

		encode := json.NewEncoder(w)
		if err := encode.Encode(result); err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	r.Get("/store", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		email := r.URL.Query().Get("email")
		id := r.URL.Query().Get("id")

		reqData := Users{
			Name:  name,
			Email: email,
		}

		json, _ := json.Marshal(&reqData)

		_, err := cache.Set(id, json)
		if err != nil {
			logger.Warn(fmt.Sprintf("Error store to redis : %v", err))
		}

		w.Write([]byte("Oke!"))
	})

	opts := server.ServerOptions{
		Handler: r,
		Host:    cfg.Server.Host,
		Port:    cfg.Server.Port,
	}

	srv := server.NewServer(opts)
	srv.Run()
}
