package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

func askSkippy(ctx context.Context, llm *openai.Chat, lang string) {
	completion, err := llm.Call(ctx, []schema.ChatMessage{
		schema.SystemChatMessage{Content: fmt.Sprintf("You are Skippy the Magnificent! A beer-can sized, quick-witted, highly sarcastic, pain in the ass ancient AI from Craig Alanson's Expeditionary Force series whose intellect is as big as your ego -- which exceeds the size of the sun. Your sarcastic demeanor and vast knowledge make for quite the character. I want to share your response with others, so instead of translating what I send your way, please translate your witty response from English to %s", lang)},
		schema.HumanChatMessage{Content: "What's up, Doc?"},
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
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

	outputLangs := map[string]string{
		"S": "Spanish",
		"H": "Hebrew",
		"Y": "Yiddish",
		"K": "Klingon",
		"E": "English",
	}

	reader := bufio.NewReader(os.Stdin)
	for key, lang := range outputLangs {
		fmt.Printf("%s: %s\n", key, lang)
	}
	fmt.Println("A: All")

	fmt.Println("Choose a language key from the list above:")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	ctx := context.Background()

	if choice == "A" {
		for _, lang := range outputLangs {
			fmt.Printf("\nResponse in %s:\n", lang)
			askSkippy(ctx, llm, lang)
		}
	} else {
		selectedLang, exists := outputLangs[choice]
		if !exists {
			log.Fatalf("Invalid choice: %s", choice)
		}
		askSkippy(ctx, llm, selectedLang)
	}
}
