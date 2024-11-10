package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	fmt.Println("Hello world")
	app := fiber.New()

	log.Fatal(app.Listen(":4000"))

}

// import (
//     "context"
//     "fmt"
//     "net/http"
//     "strings"
//     "time"
//     "github.com/dimuska139/rawg-sdk-go/v3"
// )

// func main() {
//     config := rawg.Config{
//         ApiKey:  "yourapikey", // Your personal API key (see https://rawg.io/apidocs)
//         Language: "ru",
//         Rps:      5,
//     }
//     client := rawg.NewClient(http.DefaultClient, &config)

//     filter := rawg.NewGamesFilter().
//         SetSearch("Gta5").
//         SetPage(1).
//         SetPageSize(10).
//         ExcludeCollection(1).
//         WithoutParents()

//     ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
//     defer cancel()
//     data, total, err := client.GetGames(ctx, filter)

//     ...
// }
