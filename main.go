package main

import (
	"context"
	"softarchrest/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	ctx := context.Background()
	sv := server.NewServer(app, ctx)
	sv.Start(true) // true for local mode, false for deploy mode
}
