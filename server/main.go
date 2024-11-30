package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dimuska139/rawg-sdk-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type Todo struct {
	ID        int    `json:"id" bson: "_id`
	Completed string `json:"completed"`
	body      string `json:"body"`
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	fmt.Println("hello world")
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// MongoDB setup
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Verify MongoDB connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MONGODB ATLAS")

	// Initialize Fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	// Retrieve API key from .env
	apiKey := goDotEnvVariable("RAWG_API_KEY")
	if apiKey == "" {
		log.Fatal("RAWG_API_KEY is not set in the environment")
	}

	rawgConfig := rawg.Config{
		ApiKey:   apiKey, // Your personal API key (see https://rawg.io/apidocs)
		Language: "ru",
		Rps:      5,
	}
	rawgClient := rawg.NewClient(http.DefaultClient, &rawgConfig)

	app.Get("/games", func(c *fiber.Ctx) error {
		filter := rawg.NewGamesFilter().
			SetSearch("Gta5").
			SetPage(1).
			SetPageSize(10).
			ExcludeCollection(1).
			WithoutParents()

		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		data, total, err := rawgClient.GetGames(ctx, filter)
		if err != nil {
			log.Printf("Error fetching games: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch games"})
		}

		return c.Status(200).JSON(fiber.Map{
			"total": total,
			"games": data,
		})
	})

	// Start the server
	log.Fatal(app.Listen(":4000"))
}
