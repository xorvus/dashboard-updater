package server

import (
	_ "dashboard-updater/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"os"
)

// Health godoc
// @Summary      Check Health
// @Description  Check Health
// @Accept       json
// @Produce      json
// @Tags         health
// @Router       /health [get]
func (s *FiberServer) RegisterFiberRoutes() {
	// Apply CORS middleware
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	basicAuthMiddleware := basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("DOCS_USER"): os.Getenv("DOCS_PASS"), // ganti sesuai kebutuhan
		},
	})

	s.App.Get("/docs/*", basicAuthMiddleware, swagger.HandlerDefault)

	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/health", s.healthHandler)

}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
