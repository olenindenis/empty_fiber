package application

import (
	"encoding/json"
	_ "envs/docs"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const (
	DateTimeLayout = "15:04:05 02-01-2006"
)

type HttpServer struct {
	client *Client
}

func NewHttpServer(options ...Option) *HttpServer {
	client := &Client{
		host: defaultHttpServerHost,
		port: defaultHttpServerPort,
	}
	for _, option := range options {
		option(client)
	}

	return &HttpServer{
		client: client,
	}
}

func (s *HttpServer) GetServer() *fiber.App {
	return fiber.New(fiber.Config{
		DisableStartupMessage: false,
		Prefork:               false,
		CaseSensitive:         false,
		StrictRouting:         true,
		IdleTimeout:           idleTimeout,
		ServerHeader:          "CustomServer",
		JSONEncoder:           json.Marshal,
	})
}

func (s *HttpServer) GetDSN() string {
	return fmt.Sprintf("%s:%s", s.client.host, s.client.port)
}
