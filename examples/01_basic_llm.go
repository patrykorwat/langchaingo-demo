package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

// RunBasicLLM demonstrates basic LLM text completion with various options
func RunBasicLLM() {
	fmt.Println("🤖 Basic LLM Example")
	fmt.Println("Demonstrates simple text completion with Claude\n")

	// Initialize the LLM
	llm, err := anthropic.New(
		anthropic.WithModel("claude-sonnet-4-5-20250929"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Example 1: Simple completion
	fmt.Println("Example 1: Simple Completion")
	fmt.Println("Prompt: 'Explain what LangChain is in one sentence'")
	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		"Explain what LangChain is in one sentence",
		llms.WithTemperature(0.7),
	)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Response: %s\n\n", completion)

	// Example 2: Completion with temperature control
	fmt.Println("Example 2: Temperature Control (Creative vs Deterministic)")
	fmt.Println("High temperature (0.9) - More creative:")
	completion, err = llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		"Write a creative tagline for a coffee shop",
		llms.WithTemperature(0.9),
	)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%s\n\n", completion)

	fmt.Println("Low temperature (0.1) - More deterministic:")
	completion, err = llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		"Write a creative tagline for a coffee shop",
		llms.WithTemperature(0.1),
	)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%s\n\n", completion)

	// Example 3: Max tokens control
	fmt.Println("Example 3: Max Tokens Control")
	completion, err = llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		"Explain quantum computing",
		llms.WithMaxTokens(50),
	)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Response (limited to ~50 tokens): %s\n", completion)
}
