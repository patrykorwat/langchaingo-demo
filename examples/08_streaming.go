package examples

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

// RunStreamingExample demonstrates streaming responses from LLMs
func RunStreamingExample() {
	fmt.Println("🌊 Streaming Example")
	fmt.Println("Demonstrates real-time streaming responses\n")

	// Initialize the LLM
	llm, err := anthropic.New(
		anthropic.WithModel("claude-sonnet-4-5-20250929"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Example 1: Basic Streaming
	fmt.Println("Example 1: Basic Streaming Response")
	fmt.Println("Watch the response appear in real-time\n")
	fmt.Println("Prompt: 'Write a haiku about programming'\n")
	fmt.Println("Response:")

	_, err = llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "Write a haiku about programming"),
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\n")

	// Example 2: Streaming with Timing
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Example 2: Streaming with Timing Metrics")
	fmt.Println("Measure time to first token and total time\n")
	fmt.Println("Prompt: 'List 3 programming languages'\n")
	fmt.Println("Response:")

	startTime := time.Now()
	var firstTokenTime time.Time
	firstToken := true

	_, err = llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "List 3 popular programming languages with one sentence about each"),
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		if firstToken {
			firstTokenTime = time.Now()
			firstToken = false
		}

		fmt.Print(string(chunk))
		return nil
	}))

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	totalTime := time.Since(startTime)
	timeToFirstToken := firstTokenTime.Sub(startTime)

	fmt.Printf("\n\nTiming Metrics:")
	fmt.Printf("\n  Time to first token: %v", timeToFirstToken)
	fmt.Printf("\n  Total response time: %v\n", totalTime)
}
