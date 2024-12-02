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
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// game detail information
type GameDetailed struct {
	ID           int    `json:"id"`
	Slug         string `json:"slug"`
	Name         string `json:"name"`
	NameOriginal string `json:"name_original"`
	Description  string `json:"description"`
	Metacritic   int    `json:"metacritic"`
	// MetacriticPlatforms       []*MetacriticPlatform `json:"metacritic_platforms"`
	// Released                  DateTime              `json:"released"`
	// Tba                       bool                  `json:"tba"`
	// Updated                   DateTime              `json:"updated"`
	ImageBackground           string `json:"background_image"`
	ImageBackgroundAdditional string `json:"background_image_additional"`
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// get game function
// func (api *Client) GetGame(id int) (*GameDetailed, error)

func main() {
	fmt.Println("hello world")
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// MongoDB setup

	// MONGODB_URI := os.Getenv("MONGODB_URI")
	// clientOptions := options.Client().ApplyURI(MONGODB_URI)
	// client, err := mongo.Connect(context.Background(), clientOptions)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Verify MongoDB connection

	// err = client.Ping(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connected to MONGODB ATLAS")

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
	config := rawg.Config{
		ApiKey:   "RAWG_API_KEY", // Your personal API key (see https://rawg.io/apidocs)
		Language: "ru",
		Rps:      5,
	}

	client = rawg.NewClient(http.DefaultClient, &config)

	app.Get("/games", func(c *fiber.Ctx) error {
		filter := rawg.NewGamesFilter().
			SetSearch("Gta5").
			SetPage(1).
			SetPageSize(10).
			ExcludeCollection(1).
			WithoutParents()

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
		defer cancel()
		data, total, err := client.GetGames(ctx, filter)

		// ...

		// add in a return statement
		return c.Status(200).JSON(fiber.Map{})
	})

	// Start the server
	log.Fatal(app.Listen(":4000"))
}

// rawgConfig := rawg.Config{
// 	ApiKey:   apiKey, // Your personal API key (see https://rawg.io/apidocs)
// 	Language: "ru",
// 	Rps:      5,
// }
// rawgClient := rawg.NewClient(http.DefaultClient, &rawgConfig)

// app.Get("/games", func(c *fiber.Ctx) error {
// 	filter := rawg.NewGamesFilter().
// 		SetSearch("Gta5").
// 		SetPage(1).
// 		SetPageSize(10).
// 		ExcludeCollection(1).
// 		WithoutParents()

// 	// ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
// 	// defer cancel()

// 	gamesData, err := rawgClient.GetGames(filter)
// 	if err != nil {
// 		log.Printf("Error fetching games: %v", err)
// 		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch games"})
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"total": gamesData.Count,
// 		"games": gamesData.Results,
// 	})
// })
