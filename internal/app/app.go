package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/segmentio/kafka-go"
	"leetty-gateway/internal/config"
	kafka2 "leetty-gateway/internal/kafka"
	"leetty-gateway/internal/listener"
	"leetty-gateway/internal/logger"
	"net/http"
	"strconv"
)

type App struct {
	Config      *config.Config
	Router      *chi.Mux
	KafkaWriter *kafka.Writer
}

func NewApp(config *config.Config) *App {
	return &App{
		Config:      config,
		Router:      chi.NewRouter(),
		KafkaWriter: kafka2.CreateKafkaWriter(config)}
}

func (a *App) Start() {
	var port = strconv.Itoa(a.Config.Server.Port)
	logger.Logger.Info("starting server on port: " + port)
	kafka2.CreateKafkaTopics(a.Config)
	listener.PrepareRouter(a.Config, a.Router, a.KafkaWriter)
	_ = http.ListenAndServe(":"+port, a.Router)
}

func (a *App) Close() {
	logger.Logger.Info("shutting down server")
	logger.Logger.Debug("close kafka writer resource")
	err := a.KafkaWriter.Close()
	if err != nil {
		logger.Logger.Debug("failed to close kafka writer resource", err)
		return
	}
}
