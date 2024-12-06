package main

import (
	"C"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

	// Initialize Fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	url := "https://api.rawg.io/api/games?key=< apiKey > "

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	// add in a return statement
	// return c.Status(200).JSON(fiber.Map{})

	// Start the server
	log.Fatal(app.Listen(":4000"))

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
}

// old api key stuff

// Retrieve API key from .env
// apiKey := goDotEnvVariable("RAWG_API_KEY")
// if apiKey == "" {
// 	log.Fatal("RAWG_API_KEY is not set in the environment")
// }
// config := rawg.Config{
// 	ApiKey:   "RAWG_API_KEY", // Your personal API key (see https://rawg.io/apidocs)
// 	Language: "ru",
// 	Rps:      5,
// }

// client = rawg.NewClient(http.DefaultClient, &config)

// app.Get("/games", func(c *fiber.Ctx) error {
// 	filter := rawg.NewGamesFilter().
// 		SetSearch("Gta5").
// 		SetPage(1).
// 		SetPageSize(10).
// 		ExcludeCollection(1).
// 		WithoutParents()

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
// 	defer cancel()
// 	data, total, err := client.GetGames(ctx, filter)

// ...
