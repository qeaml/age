package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django"
)

var routesGET = map[string]fiber.Handler{
	"/":     routeHome,
	"/time": routeTime,
}

func routeHome(c *fiber.Ctx) error {
	return render(c, "home", fiber.Map{})
}

func routeTime(c *fiber.Ctx) error {
	t := time.Now()
	return render(c, "time", fiber.Map{
		"time": t.Format(time.UnixDate),
	})
}

//go:embed templates
var templates embed.FS

func main() {
	templateFS, _ := fs.Sub(templates, "templates")
	views := django.NewFileSystem(http.FS(templateFS), ".jinja")

	app := fiber.New(fiber.Config{
		AppName:      "age",
		ServerHeader: "age",
		Views:        views,
		ViewsLayout:  "layout",
	})

	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("BodyOnly", false)
		return c.Next()
	})
	addRoutes(app)
	body := app.Group("/body")
	body.Use(func(c *fiber.Ctx) error {
		c.Locals("BodyOnly", true)
		return c.Next()
	})
	addRoutes(body)

	app.Static("/static", "static")

	log.Println("Starting app.")
	go app.Listen(":1234")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)

	log.Println("Press Ctrl+C to stop.")
	<-sc

	log.Println("Shutting down gracefully.")
	app.Shutdown()
	log.Println("Shut down complete.")
}

func addRoutes(c fiber.Router) {
	for route, handler := range routesGET {
		c.Get(route, handler)
	}
}

func render(c *fiber.Ctx, template string, bind any) error {
	if c.Locals("BodyOnly").(bool) {
		return c.Render(template, bind, "")
	} else {
		return c.Render(template, bind)
	}
}
