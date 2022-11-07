package application

import (
	"bytes"
	_ "envs/docs"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HttpDSN struct {
	client *Client
}

func NewHttpDSN(options ...Option) HttpDSN {
	client := &Client{
		host: defaultHttpServerHost,
		port: defaultHttpServerPort,
	}
	for _, option := range options {
		option(client)
	}

	return HttpDSN{
		client: client,
	}
}

func (s *HttpDSN) DSN() string {
	buffer := bytes.NewBufferString("")
	buffer.WriteString(s.client.host)
	buffer.WriteString(":")
	buffer.WriteString(s.client.port)
	return buffer.String()
}

func NewServer() *fiber.App {
	server := fiber.New(fiber.Config{
		DisableStartupMessage: false,
		Prefork:               false,
		CaseSensitive:         false,
		StrictRouting:         true,
		ServerHeader:          "CustomServer",
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Info(err.Error())
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

			if code == http.StatusInternalServerError {
				return ctx.Status(code).SendString(http.StatusText(http.StatusInternalServerError))
			}

			return ctx.Status(code).SendString(err.Error())
		},
	})

	return server
}
