package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	sloggin "github.com/samber/slog-gin"
)

type FiberConfig struct {
	HttpPort     string
	AllowOrigins string
	AllowMethods string
	AllowHeaders string
}

type FiberServer struct {
	App  *fiber.App
	Port string
}

func (s *FiberServer) Start() error {
	return s.App.Listen(":" + s.Port)
}

func NewFiberServer(conf *FiberConfig) Server {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(sloggin.New(slog.Default()))
	app.Use(gin.Logger())
	app.Use(cors.New(cors.Config{
		AllowOrigins: conf.AllowOrigins,
		AllowMethods: conf.AllowMethods,
		AllowHeaders: conf.AllowHeaders,
	}))

	return &FiberServer{
		App:  app,
		Port: conf.HttpPort,
	}
}
