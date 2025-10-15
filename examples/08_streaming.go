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

	// Example 2: Streaming with Processing
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Example 2: Streaming with Character Count")
	fmt.Println("Count characters as they arrive\n")
	fmt.Println("Prompt: 'Explain what an API is in simple terms'\n")
	fmt.Println("Response:")

	charCount := 0
	wordCount := 0
	lastWord := ""

	_, err = llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "Explain what an API is in simple terms (2-3 sentences)"),
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		text := string(chunk)
		fmt.Print(text)

		charCount += len(text)

		// Simple word counting (splits on spaces)
		for _, char := range text {
			if char == ' ' || char == '\n' {
				if lastWord != "" {
					wordCount++
					lastWord = ""
				}
			} else {
				lastWord += string(char)
			}
		}

		return nil
	}))

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	if lastWord != "" {
		wordCount++
	}

	fmt.Printf("\n\nStatistics:")
	fmt.Printf("\n  Characters streamed: %d", charCount)
	fmt.Printf("\n  Approximate words: %d\n\n", wordCount)

	// Example 3: Streaming with Timing
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Example 3: Streaming with Timing Metrics")
	fmt.Println("Measure time to first token and total time\n")
	fmt.Println("Prompt: 'List 5 programming languages'\n")
	fmt.Println("Response:")

	startTime := time.Now()
	var firstTokenTime time.Time
	firstToken := true

	_, err = llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "List 5 popular programming languages with one sentence about each"),
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
	fmt.Printf("\n  Total response time: %v\n\n", totalTime)

	// Example 4: Simulated Progress Bar
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Example 4: Non-Streaming vs Streaming Comparison")
	fmt.Println("Notice the difference in perceived speed\n")

	// Non-streaming
	fmt.Println("Non-streaming (wait for complete response):")
	fmt.Print("Waiting...")
	nonStreamStart := time.Now()

	response, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		"Write a 2-sentence description of machine learning",
		llms.WithTemperature(0.7),
	)

	nonStreamTime := time.Since(nonStreamStart)
	fmt.Printf("\n%s", response)
	fmt.Printf("\nTime: %v\n\n", nonStreamTime)

	// Streaming
	fmt.Println("Streaming (see response as it generates):")
	streamStart := time.Now()

	_, err = llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "Write a 2-sentence description of machine learning"),
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))

	streamTime := time.Since(streamStart)
	fmt.Printf("\nTime: %v\n\n", streamTime)

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	// Example 5: Streaming with Custom Handler
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Example 5: Streaming with Custom Processing")
	fmt.Println("Highlight specific words as they stream\n")
	fmt.Println("Prompt: 'Describe cloud computing'\n")
	fmt.Println("Response (highlighting 'cloud' and 'data'):")

	keywords := []string{"cloud", "data", "computing"}
	buffer := ""

	_, err = llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "Describe cloud computing in 2 sentences"),
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		text := string(chunk)
		buffer += text

		// Check if buffer contains any keywords
		for _, keyword := range keywords {
			if len(buffer) >= len(keyword) {
				// Simple keyword detection
				checkText := buffer[len(buffer)-len(keyword):]
				if checkText == keyword {
					// Print with emphasis
					fmt.Printf("\033[1m%s\033[0m", text) // Bold text
					return nil
				}
			}
		}

		fmt.Print(text)
		return nil
	}))

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\n")

	// Summary
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Summary:")
	fmt.Println("- Streaming provides better user experience")
	fmt.Println("- Allows processing chunks as they arrive")
	fmt.Println("- Useful for real-time applications")
	fmt.Println("- Can track metrics and progress")
	fmt.Println("- Lower perceived latency than batch responses")
}
