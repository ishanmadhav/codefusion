package server

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/ishanmadhav/codefusion/internal/loader"
)

type Server struct {
	App           *fiber.App
	KafkaProducer *kafka.Producer
}

func NewServer() *Server {
	app := fiber.New()
	kafkaConfigFile := "getting-started.properties"
	conf := loader.ReadConfig(kafkaConfigFile)
	p, err := kafka.NewProducer(&conf)
	if err != nil {
		panic("Failed to create Kafka producer")
	}

	return &Server{
		App:           app,
		KafkaProducer: p,
	}
}

func (s *Server) Start() error {

	go func() {
		for e := range s.KafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	s.setupRoutes()
	err := s.App.Listen(":3000")
	if err != nil {
		return err
	}
	return nil
}
