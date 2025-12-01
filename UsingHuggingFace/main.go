package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	godotenv.Load()
	// Load HF token
	token := os.Getenv("HF_TOKEN")
	if token == "" {
		log.Fatal("ERROR: HF_TOKEN environment variable is not set")
	}

	// Create HF Router client (OpenAI-compatible)
	config := openai.DefaultConfig(token)

    // HF inference router URL
    config.BaseURL = "https://router.huggingface.co/v1"

    client := openai.NewClientWithConfig(config)

	// Send a test prompt
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "google/gemma-2-2b-it:nebius",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: "What is the capital of France?",
				},
			},
		},
	)
	if err != nil {
		log.Fatal("Request failed:", err)
	}

	// Print model response
	fmt.Println("Model reply:")
	fmt.Println(resp.Choices[0].Message.Content)
}
