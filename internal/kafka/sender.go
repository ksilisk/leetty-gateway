package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"leetty-gateway/internal/logger"
	"log/slog"
)

type UpdateRequest struct {
	Topic       string
	Partition   int
	MessageBody []byte
}

func UpdatesSending(dataPipe <-chan *UpdateRequest, writer *kafka.Writer) {
	logger.Logger.Info("starting kafka-sender goroutine")
	for message := range dataPipe {
		logger.Logger.Debug("read message from kafka-sender channel", slog.Attr{Key: "QueueSize", Value: slog.IntValue(len(dataPipe))})
		logger.Logger.Info("sending update to kafka", slog.Attr{Key: "message", Value: slog.StringValue(string(message.MessageBody))})
		kafkaMessage := kafka.Message{
			Topic:     message.Topic,
			Partition: message.Partition,
			Value:     message.MessageBody,
		}
		err := writer.WriteMessages(context.TODO(), kafkaMessage)
		if err != nil {
			logger.Logger.Error("unable to send update to kafka. skip update.",
				slog.Attr{Key: "Message", Value: slog.AnyValue(message)}, err)
		}
	}
	logger.Logger.Info("stopped kafka-sender goroutine")
}
