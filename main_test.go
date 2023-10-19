package main

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tmc/langchaingo/schema"
)

type MockChat struct {
	mock.Mock
}

func TestAskSkippy(t *testing.T) {
	// Create an instance of our mock Chat
	mockChat := new(MockChat)

	// Setup expectations
	mockChat.On("Call", mock.Anything, mock.Anything, mock.Anything).Return(&schema.AIChatMessage{Content: "Mocked response"}, nil)

	// Call the askSkippy function with a custom message
	response, err := askSkippy(context.Background(), mockChat, "English", "Test message for Skippy")

	// Check for errors
	if err != nil {
		t.Fatalf("askSkippy failed with error: %v", err)
	}

	// Check if the response matches our expectation
	if response != "Mocked response" {
		t.Fatalf("Expected 'Mocked response' but got '%s'", response)
	}
}

func TestAskSkippy_Error(t *testing.T) {
	// Create an instance of our mock Chat
	mockChat := new(MockChat)

	// Setup expectations for an error scenario
	mockChat.On("Call", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("Mocked error"))

	// Call the askSkippy function with a custom message
	_, err := askSkippy(context.Background(), mockChat, "English", "Test error message for Skippy")

	// Check for errors
	if err == nil {
		t.Fatalf("Expected an error but got none")
	}
}
