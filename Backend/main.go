package main

import (
	"Backend/config"
	"Backend/router"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)
func init() {
	 err := godotenv.Load()
	 if err != nil {
		 log.Fatalf("Error loading .env file")
	 }
}

func main() {
	config.DB = config.MongoConnect(config.DBName)
	if config.DB == nil {
		log.Fatal("Failed to connect to MongoDB")
	}


	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(),","),
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
}))
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Route not found",
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
