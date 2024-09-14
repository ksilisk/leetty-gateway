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
	KafkaWriter *kafka.Writer
	router      *chi.Mux
	dataChan    chan http.Request
}

func NewApp(config *config.Config) *App {
	return &App{
		Config:      config,
		router:      chi.NewRouter(),
		KafkaWriter: kafka2.CreateKafkaWriter(config)}
}

func (a *App) Start() {
	dataChan := make(chan *kafka2.UpdateRequest, a.Config.App.QueueSize)
	var port = strconv.Itoa(a.Config.Server.Port)
	logger.Logger.Info("starting server on port: " + port + ", queueSize: " + strconv.Itoa(a.Config.App.QueueSize))
	kafka2.CreateKafkaTopics(a.Config)
	listener.PrepareRouter(a.Config, a.router, dataChan)
	go kafka2.UpdatesSending(dataChan, a.KafkaWriter)
	_ = http.ListenAndServe(":"+port, a.router)
}

func (a *App) Close() {
	logger.Logger.Info("shutting down server")
	logger.Logger.Info("close kafka-sender channel")
	close(a.dataChan)
	logger.Logger.Debug("close kafka writer resource")
	err := a.KafkaWriter.Close()
	if err != nil {
		logger.Logger.Debug("failed to close kafka writer resource", err)
		return
	}
}
