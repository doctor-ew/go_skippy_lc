package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tmc/langchaingo/schema"
)

type MockChat struct {
	mock.Mock
}

func (m *MockChat) Call(ctx context.Context, messages []schema.ChatMessage, options ...llms.CallOption) (*schema.AIChatMessage, error) {
	args := m.Called(ctx, messages, options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*schema.AIChatMessage), args.Error(1)
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

func TestRootEndpoint(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Register your endpoints
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record the response
	resp := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(resp, req)

	// Check the response status code
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status 200 but got %d", resp.Code)
	}

	// Check the response body
	expectedBody := `{"message":"Hello World!"}`
	if resp.Body.String() != expectedBody {
		t.Fatalf("Expected body '%s' but got '%s'", expectedBody, resp.Body.String())
	}
}
