package main

import (
	"context"
	"testing"

	"github.com/tmc/langchaingo/llms/openai"
)

func TestAskSkippy(t *testing.T) {
	// Mocking the llm.Call method for testing
	mockChat := &openai.Chat{}

	// Call the askSkippy function
	response, err := askSkippy(context.Background(), mockChat, "English")

	// Check for errors
	if err != nil {
		t.Fatalf("askSkippy failed with error: %v", err)
	}

	// Check if the response is not empty
	if response == "" {
		t.Fatalf("Expected a non-empty response from Skippy")
	}
}
