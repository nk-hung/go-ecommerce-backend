package kafka

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var kafkaProducer *kafka.Writer

const (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "user_topic_0001"
)

// For Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// for Consumer
func getKafkaReader(kafkaUrl, topic, groupID string) *kafka.Reader {
	broker := strings.Split(kafkaUrl, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        broker,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10kb
		MaxBytes:       10e6, // 10Mb
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))

	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "Action successfully",
	})
}
