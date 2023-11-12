package server

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ishanmadhav/codefusion/internal/api"
)

type CodePayload struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Input    string `json:"input"`
}

type Resp struct {
	Message string `json:"message"`
}

//Code Controllers

// Gets the code execution results by ID, results like output
func (s *Server) getCodeExecutionResultsByID(c *fiber.Ctx) error {
	return c.SendString("Hello Code")
}

// Submits code for execution
func (s *Server) submitCode(c *fiber.Ctx) error {
	var codePayload CodePayload
	if err := c.BodyParser(&codePayload); err != nil {
		return c.Status(400).JSON(err)
	}
	code := api.Code{
		Code:     codePayload.Code,
		Language: codePayload.Language,
		Input:    codePayload.Code,
		Output:   "",
	}
	fmt.Print(code)
	topic := "executions"
	uniqueID := uuid.New()
	codeJson, err := json.Marshal(code)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	s.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(uniqueID.String()),
		Value:          codeJson,
	}, nil)

	resp := Resp{
		Message: "Code submitted successfully",
	}
	return c.JSON(resp)
}

// list of all the executed code programs
func (s *Server) getAllCodes(c *fiber.Ctx) error {
	return c.SendString("All codes are fetched here")
}

// delete code by ID, like a utility function for deletion
func (s *Server) deleteCodeByID(c *fiber.Ctx) error {
	return c.SendString("Code is deleted here")
}

// Room Controllers

// Gets the room by ID, room is the collaborative editor
func (s *Server) getRoomByID(c *fiber.Ctx) error {
	return c.SendString("Room is fetched here")
}

// Creates a new room, room is the collaborative editor
func (s *Server) createRoom(c *fiber.Ctx) error {
	return c.SendString("Room is created here")
}

// Gets the codes by room ID, room is the collaborative editor
func (s *Server) GetCodesByRoomID(c *fiber.Ctx) error {
	return c.SendString("Codes are fetched by room ID here")
}
