package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
    godotenv.Load()

	fmt.Println("Redis Connection")

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping)

	ctx := context.Background()

	person := Person{
		Name:  "uthfol",
		Email: "uthfol@gmail.com",
		Age:   24,
	}

	personJSON, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling person:", err.Error())
		return
	}

	err = client.Set(ctx, "user:1", personJSON, 0).Err()
	if err != nil {
		fmt.Println("Error setting user:", err.Error())
		return
	}

	fmt.Println("User set successfully:", person)

	val, err := client.Get(ctx, "user:1").Result()
	if err != nil {
		fmt.Println("Error getting user:", err.Error())
		return
	}

	var retrievedPerson Person
	err = json.Unmarshal([]byte(val), &retrievedPerson)
	if err != nil {
		fmt.Println("Error unmarshaling person:", err.Error())
		return
	}

	fmt.Println("User retrieved successfully:", retrievedPerson)
}
