package kafka

import (
	"github.com/segmentio/kafka-go"
	"leetty-gateway/internal/config"
	"leetty-gateway/internal/logger"
)

func CreateKafkaWriter(config *config.Config) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(config.KafkaBrokers...),
		AllowAutoTopicCreation: true,
	}
}

func CreateKafkaTopics(config *config.Config) {
	conn, err := kafka.Dial("tcp", config.KafkaBrokers[0])
	if err != nil {
		logger.Logger.Error("error while creating kafka connection", err)
		panic(err)
	}
	defer func(conn *kafka.Conn) {
		logger.Logger.Debug("closing kafka dial connection")
		err := conn.Close()
		if err != nil {
			logger.Logger.Debug("error while closing kafka dial connection", err)
		}
	}(conn)
	for _, value := range config.Mapping {
		logger.Logger.Info("create if not exists kafka topic: " + value.KafkaTopic)
		var topicConfig = kafka.TopicConfig{Topic: value.KafkaTopic, NumPartitions: -1, ReplicationFactor: -1}
		err := conn.CreateTopics(topicConfig)
		if err != nil {
			logger.Logger.Error("error while creating kafka topic", err)
			panic(err)
		}
	}
}
