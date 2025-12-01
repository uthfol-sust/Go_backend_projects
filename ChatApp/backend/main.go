package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Load .env if present
	godotenv.Load()

	token := os.Getenv("HF_TOKEN")
	if token == "" {
		log.Fatal("HF_TOKEN not set. Copy .env.example to .env and set HF_TOKEN")
	}

	model := os.Getenv("HF_MODEL")
	if model == "" {
		model = "google/gemma-2-2b-it:nebius"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config := openai.DefaultConfig(token)
	config.BaseURL = "https://router.huggingface.co/v1"
	client := openai.NewClientWithConfig(config)

	r := gin.Default()

	// Simple CORS middleware to allow local frontend
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Health
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	type chatRequest struct {
		Message string `json:"message" binding:"required"`
	}

	type chatResponse struct {
		Reply string `json:"reply"`
	}

	r.POST("/chat", func(c *gin.Context) {
		var req chatRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "message field required"})
			return
		}

		// Build the chat completion request
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: model,
				Messages: []openai.ChatCompletionMessage{
					{Role: "user", Content: req.Message},
				},
			},
		)
		if err != nil {
			log.Printf("chat completion error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(resp.Choices) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no response from model"})
			return
		}

		c.JSON(http.StatusOK, chatResponse{Reply: resp.Choices[0].Message.Content})
	})

	log.Printf("starting server on :%s (model=%s)", port, model)
	r.Run(":" + port)

}
