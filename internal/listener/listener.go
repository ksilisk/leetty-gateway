package listener

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/segmentio/kafka-go"
	"io"
	"leetty-gateway/internal/config"
	"leetty-gateway/internal/logger"
	"net/http"
)

func PrepareRouter(config *config.Config, router *chi.Mux, writer *kafka.Writer) {
	addMiddlewares(router)
	router.Route("/", func(r chi.Router) {
		for _, botMapping := range config.Mapping {
			logger.Logger.Info("add endpoint: /" + botMapping.Endpoint)
			r.Post("/"+botMapping.Endpoint, func(w http.ResponseWriter, r *http.Request) {
				body := r.Body
				bodyBytes, _ := io.ReadAll(body)
				kafkaMessage := kafka.Message{
					Topic:     botMapping.KafkaTopic,
					Partition: botMapping.Partition,
					Value:     bodyBytes,
				}
				err := writer.WriteMessages(r.Context(), kafkaMessage)
				if err != nil {
					logger.Logger.Error("error while sending message", kafkaMessage)
					return
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
