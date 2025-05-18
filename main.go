package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/karan/practicelapbs/database"
	"github.com/karan/practicelapbs/handlers"
	"github.com/karan/practicelapbs/models"
	routes "github.com/karan/practicelapbs/routes"
	"github.com/karan/practicelapbs/utils"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	app := fiber.New()
	// if you receive Bad request log out the user from FE...
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	if err := utils.LoadProblemsFromFile("data/problems.json"); err != nil {
		log.Fatal("Failed to load problems:", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("You are connected to the Practice Labs Server")
	})
	app.Get("/api/problems", handlers.GetAllProblems)
	app.Get("/api/problems/:slug", handlers.GetProblemBySlug)
	routes.AuthRoutes(app)

	println("🚀 Server running at http://localhost:8000")
	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}

}
