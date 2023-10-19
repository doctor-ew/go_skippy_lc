package main

import (
	"context"
	"errors"
	"github.com/tmc/langchaingo/llms"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tmc/langchaingo/schema"
)

type MockChat struct {
	mock.Mock
}

func (m *MockChat) Call(ctx context.Context, messages []schema.ChatMessage, options ...llms.CallOption) (*schema.AIChatMessage, error) {
	args := m.Called(ctx, messages, options)
	return args.Get(0).(*schema.AIChatMessage), args.Error(1)
}

func TestAskSkippy(t *testing.T) {
	// Create an instance of our mock Chat
	mockChat := new(MockChat)

	// Setup expectations
	mockChat.On("Call", mock.Anything, mock.Anything, mock.Anything).Return(&schema.AIChatMessage{Content: "Mocked response"}, nil)

	// Call the askSkippy function
	response, err := askSkippy(context.Background(), mockChat, "English")

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
	emptyResponse := &schema.AIChatMessage{}
	mockChat.On("Call", mock.Anything, mock.Anything, mock.Anything).Return(emptyResponse, errors.New("Mocked error"))

	// Call the askSkippy function
	_, err := askSkippy(context.Background(), mockChat, "English")

	// Check for errors
	if err == nil {
		t.Fatalf("Expected an error but got none")
	}
}
