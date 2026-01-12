// Package main : a simple web app
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/rand"
)

/*
GetBytesBody returns the body of an HTTP request as a []byte.
  - It takes a pointer to an http.Request as a parameter.
  - It returns a []byte.

func GetBytesBody(request *http.Request) []byte {
	body := make([]byte, request.ContentLength)
	request.Body.Read(body)
	return body
}
*/

type Restaurant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
	Tags    string `json:"tags"`
}

func main() {
	ctx := context.Background()

	var redisServer = os.Getenv("REDIS_URL")

	client := redis.NewClient(&redis.Options{
		Addr:     redisServer, // the name is defined in the compose.yml file
		Password: "",          // no password set
		DB:       0,           // use default DB
	})

	appName := generateFunnyName()

	var httpPort = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	var message = os.Getenv("MESSAGE")
	if message == "" {
		message = "this is a message..."
	}

	var mainTitle = os.Getenv("TITLE")
	if mainTitle == "" {
		mainTitle = "this is a title"
	}

	log.Println("🚀 starting web server on port: " + httpPort)
	log.Println("📝 message: " + message)

	mux := http.NewServeMux()

	fileServerHtml := http.FileServer(http.Dir("public"))
	mux.Handle("/", fileServerHtml)

	mux.HandleFunc("/api/info", func(response http.ResponseWriter, request *http.Request) {
		// read the content of the file info.txt
		// and return it as a response
		content, err := os.ReadFile("./public/info.txt")
		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		}
		response.WriteHeader(http.StatusOK)
		response.Write(content)
	})

	mux.HandleFunc("/api/variables", func(response http.ResponseWriter, request *http.Request) {

		variables := map[string]interface{}{
			"message":   message,
			"appName":   appName,
			"mainTitle": mainTitle,
		}

		jsonString, err := json.Marshal(variables)

		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		}

		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.WriteHeader(http.StatusOK)
		response.Write(jsonString)
	})

	mux.HandleFunc("/api/restaurants", func(response http.ResponseWriter, request *http.Request) {
		restaurants, err := getAllRestaurants(ctx, client)
		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		}

		jsonData, err := json.MarshalIndent(restaurants, "", "  ")
		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		}
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.WriteHeader(http.StatusOK)
		response.Write(jsonData)

	})

	var errListening error
	log.Println("🌍 http server is listening on: " + httpPort)
	errListening = http.ListenAndServe(":"+httpPort, mux)

	log.Fatal(errListening)
}

func generateFunnyName() string {
	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixMilli()))

	adjectives := []string{"Funny", "Silly", "Crazy", "Goofy", "Wacky", "Loony", "Zany"}
	nouns := []string{"Chicken", "Goose", "Giraffe", "Koala", "Lemur", "Quokka", "Axolotl"}

	// Get random indices for both arrays
	index1 := rand.Intn(len(adjectives))
	index2 := rand.Intn(len(nouns))

	// Combine the selected elements to form the name
	randomName := adjectives[index1] + " " + nouns[index2]

	return randomName
}

func getRestaurant(ctx context.Context, client *redis.Client, id string) (Restaurant, error) {
	result, err := client.HGetAll(ctx, "restaurant:"+id).Result()
	if err != nil {
		return Restaurant{}, err
	}

	// Check if the restaurant exists
	if len(result) == 0 {
		return Restaurant{}, fmt.Errorf("restaurant not found")
	}

	return Restaurant{
		ID:      id,
		Name:    result["name"],
		Address: result["address"],
		Website: result["website"],
		Phone:   result["phone"],
		Tags:    result["tags"],
	}, nil
}

func getAllRestaurants(ctx context.Context, client *redis.Client) ([]Restaurant, error) {
	// Get all restaurant IDs from the set
	restaurantIDs, err := client.SMembers(ctx, "restaurants").Result()
	if err != nil {
		return nil, err
	}

	restaurants := make([]Restaurant, 0, len(restaurantIDs))

	for _, id := range restaurantIDs {
		restaurant, err := getRestaurant(ctx, client, id)
		if err != nil {
			return nil, fmt.Errorf("error retrieving restaurant %s: %v", id, err)
		}
		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}
