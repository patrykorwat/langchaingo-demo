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

	// Simple completion with temperature control
	fmt.Println("Simple Completion with Temperature Control")
	fmt.Println("Prompt: 'Explain what LangChain is in one sentence'\n")

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
	fmt.Printf("Response: %s\n", completion)
}
