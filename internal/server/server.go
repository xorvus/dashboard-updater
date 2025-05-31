package server

import (
	"github.com/gofiber/fiber/v2"

	"dashboard-updater/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "dashboard-updater",
			AppName:      "dashboard-updater",
		}),

		db: database.New(),
	}

	return server
}
