package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"softarchrest/database/gen"
	"softarchrest/handler"
	"softarchrest/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type server struct {
	fiberApp *fiber.App
	ctx      context.Context
}

func NewServer(fiberApp *fiber.App, ctx context.Context) *server {
	return &server{fiberApp: fiberApp, ctx: ctx}
}

func (s *server) Start(localMode bool) {
	// Load config
	cfg, err := loadConfig(localMode)
	if err != nil {
		log.Fatal("Cannot load environment:", err)
	}

	// Init databases
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), cfg.PostgresURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	q := gen.New(conn)

	// Init services
	ts := service.NewTaskService(q)

	// Init handlers
	th := handler.NewTaskHandler(ts)

	// Set routes
	s.setRoutes(th)

	// Start the server on port 8080
	fmt.Println("Starting server...")
	err = s.fiberApp.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
