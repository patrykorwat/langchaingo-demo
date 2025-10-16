package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
)

// RunMemory demonstrates conversation memory and context management
func RunMemory() {
	fmt.Println("🧠 Memory Example")
	fmt.Println("Demonstrates conversation history and context management\n")

	// Initialize the LLM
	llm, err := anthropic.New(
		anthropic.WithModel("claude-sonnet-4-5-20250929"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Conversation Buffer Memory
	fmt.Println("Conversation Buffer Memory")
	fmt.Println("Maintains full conversation history\n")

	// Create a simple buffer memory
	bufferMemory := memory.NewConversationBuffer()

	// Create a conversation prompt that uses history
	conversationPrompt := prompts.NewPromptTemplate(
		`The following is a conversation between a human and an AI assistant.

{{.history}}
Human: {{.input}}
AI:`,
		[]string{"history", "input"},
	)

	// Questions that build on previous context
	questions := []string{
		"My name is Alice and I love programming in Go.",
		"What programming language did I mention?",
		"What is my name?",
	}

	for i, question := range questions {
		fmt.Printf("Turn %d\n", i+1)
		fmt.Printf("Human: %s\n", question)

		// Load memory into variables
		memoryVars, err := bufferMemory.LoadMemoryVariables(ctx, map[string]any{})
		if err != nil {
			log.Printf("Error loading memory: %v\n", err)
			return
		}

		// Get chat history as string
		chatHistory := ""
		if history, ok := memoryVars["history"]; ok {
			chatHistory = fmt.Sprintf("%v", history)
		}

		// Prepare inputs
		inputs := map[string]any{
			"history": chatHistory,
			"input":   question,
		}

		// Create chain and run
		chain := chains.NewLLMChain(llm, conversationPrompt)
		result, err := chains.Call(ctx, chain, inputs)
		if err != nil {
			log.Printf("Error: %v\n", err)
			return
		}

		response := result[chain.OutputKey].(string)
		fmt.Printf("AI: %s\n\n", response)

		// Save to memory
		err = bufferMemory.SaveContext(ctx, map[string]any{"input": question}, map[string]any{"output": response})
		if err != nil {
			log.Printf("Error saving to memory: %v\n", err)
			return
		}
	}
}
