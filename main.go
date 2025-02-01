package main

import (
	"log"

	"github.com/Nidasakinaa/ws-kaloriku/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/Nidasakinaa/ws-kaloriku/url"
	_ "github.com/Nidasakinaa/ws-kaloriku/docs"

	"github.com/gofiber/fiber/v2"
)

// @title TES SWAGGER KALORIKU
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/Nidasakinaa
// @contact.email 714220040@std.ulbi.ac.id

// @host ws-kaloriku-4cf736febaf0.herokuapp.com
// @BasePath /
// @schemes https http
func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
