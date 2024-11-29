package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dimuska139/rawg-sdk-go/v3"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	fmt.Println("Hello world")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	// Retrieve API key from .env
	apiKey := goDotEnvVariable("RAWG_API_KEY")
	if apiKey == "" {
		log.Fatal("RAWG_API_KEY is not set in the environment")
	}

	config := rawg.Config{
		ApiKey:   apiKey, // Your personal API key (see https://rawg.io/apidocs)
		Language: "ru",
		Rps:      5,
	}
	client := rawg.NewClient(http.DefaultClient, &config)

	filter := rawg.NewGamesFilter().
		SetSearch("Gta5").
		SetPage(1).
		SetPageSize(10).
		ExcludeCollection(1).
		WithoutParents()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
	defer cancel()
	data, total, err := client.GetGames(ctx, filter)
	if err != nil {
		log.Fatalf("Error fetching games: %v", err)
	}

	app.Get("/games", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"total": total,
			"games": data,
		})
	})

	log.Fatal(app.Listen(":4000"))
}
