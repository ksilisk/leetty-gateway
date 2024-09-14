package listener

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"io"
	"leetty-gateway/internal/config"
	"leetty-gateway/internal/kafka"
	"leetty-gateway/internal/logger"
	"net/http"
)

func PrepareRouter(config *config.Config, router *chi.Mux, dataPipe chan<- *kafka.UpdateRequest) {
	addMiddlewares(router)
	router.Route("/", func(r chi.Router) {
		for _, botMapping := range config.Mapping {
			logger.Logger.Info("add endpoint: /" + botMapping.Endpoint)
			r.Post("/"+botMapping.Endpoint, func(w http.ResponseWriter, r *http.Request) {
				body, _ := io.ReadAll(r.Body)
				dataPipe <- &kafka.UpdateRequest{
					MessageBody: body,
					Topic:       botMapping.KafkaTopic,
					Partition:   botMapping.Partition,
				}
				w.WriteHeader(http.StatusOK)
			})
		}
	})
}

func addMiddlewares(router *chi.Mux) {
	router.Use(middleware.Heartbeat("/health"))
	router.Use(httplog.RequestLogger(&httplog.Logger{Logger: logger.Logger, Options: httplog.Options{Concise: true}}))
}
