package server

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/ishanmadhav/codefusion/internal/api"
	"github.com/ishanmadhav/codefusion/internal/loader"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn = "host=localhost user=jamadmin password=jampass dbname=jamlydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

type Server struct {
	App           *fiber.App
	KafkaProducer *kafka.Producer
	db            *gorm.DB
}

func NewServer() *Server {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
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
		db:            db,
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

	err := s.db.AutoMigrate(&api.Code{})
	if err != nil {
		return err
	}

	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	s.setupRoutes()
	err = s.App.Listen(":3000")
	if err != nil {
		return err
	}

	return nil
}
