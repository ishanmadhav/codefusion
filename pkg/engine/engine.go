package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ishanmadhav/codefusion/internal/api"
	"github.com/ishanmadhav/codefusion/internal/loader"
	"github.com/ishanmadhav/codefusion/pkg/engine/executor"
)

type CodePayload struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Input    string `json:"input"`
}

type Engine struct {
	KafkaConsumer *kafka.Consumer
}

func NewEngine() *Engine {
	kafkaConfigFile := "getting-started.properties"
	conf := loader.ReadConfig(kafkaConfigFile)
	conf["group.id"] = "codefusion_engine"
	conf["auto.offset.reset"] = "earliest"
	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		panic("Failed to create consumer")
	}

	return &Engine{
		KafkaConsumer: c,
	}
}

func (e *Engine) StartEngine() error {
	topic := "executions"
	err := e.KafkaConsumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := e.KafkaConsumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

			var code api.Code
			err = json.Unmarshal(ev.Value, &code)
			if err != nil {
				fmt.Print("Failed to unmarshal code payload")
				fmt.Print(err)
			}
			e.Execute(&code)
		}
	}

	e.KafkaConsumer.Close()
	return nil
}

func (e *Engine) Execute(code *api.Code) {
	err := executor.Execute(code)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(code.Output)
}
