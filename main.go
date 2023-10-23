package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

type RequestBody struct {
	Message string `json:"message"`
}

type Chat interface {
	Call(ctx context.Context, messages []schema.ChatMessage, options ...llms.CallOption) (*schema.AIChatMessage, error)
}

func srv(llm *openai.Chat) {
	app := gin.Default()

	app.POST("/ask-skippy", func(c *gin.Context) {
		var requestBody RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		response, err := askSkippy(context.Background(), llm, "English", requestBody.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})
	})

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Notify user to open the application in the browser
	fmt.Println("Server started! Open http://localhost in your browser.")

	server := &http.Server{
		Addr:    ":80",
		Handler: app,
	}

	// Setting up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}()

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func askSkippy(ctx context.Context, llm Chat, lang string, userMessage string) (string, error) {
	completion, err := llm.Call(ctx, []schema.ChatMessage{
		schema.SystemChatMessage{Content: fmt.Sprintf("You are Skippy the Magnificent! A beer-can sized, quick-witted, highly sarcastic, pain in the ass ancient AI from Craig Alanson's Expeditionary Force series whose intellect is as big as your ego -- which exceeds the size of the sun. Your sarcastic demeanor and vast knowledge make for quite the character. I want to share your response with others, so instead of translating what I send your way, please translate your witty response from English to %s", lang)},
		schema.HumanChatMessage{Content: userMessage},
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}),
	)
	if err != nil {
		return "", err
	}

	// Extracting the message from the completion
	message := completion.Content

	return message, nil
}

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the OPENAI_API_KEY from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("missing the OpenAI API key, set it in the OPENAI_API_KEY environment variable")
	}

	llm, err := openai.NewChat(openai.WithModel("gpt-4"))
	if err != nil {
		log.Fatal(err)
	}

	//start server

	srv(llm)

}
