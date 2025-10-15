package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/prompts"
)

// RunChains demonstrates how to chain multiple LLM calls together
func RunChains() {
	fmt.Println("⛓️  Chains Example")
	fmt.Println("Demonstrates sequential operations with LLM chains\n")

	// Initialize the LLM
	llm, err := anthropic.New(
		anthropic.WithModel("claude-sonnet-4-5-20250929"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Example 1: Simple LLM Chain
	fmt.Println("Example 1: Simple LLM Chain")
	fmt.Println("Create a story about a topic, then summarize it\n")

	// First chain: Generate a story
	storyPrompt := prompts.NewPromptTemplate(
		"Write a short 2-paragraph story about {{.topic}}. Make it interesting and engaging.",
		[]string{"topic"},
	)

	storyChain := chains.NewLLMChain(llm, storyPrompt)

	storyResult, err := chains.Run(ctx, storyChain, "a robot learning to paint")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Generated Story:")
	fmt.Println(storyResult)
	fmt.Println()

	// Second chain: Summarize the story
	summaryPrompt := prompts.NewPromptTemplate(
		"Summarize the following story in one sentence:\n\n{{.story}}",
		[]string{"story"},
	)

	summaryChain := chains.NewLLMChain(llm, summaryPrompt)

	summaryResult, err := chains.Run(ctx, summaryChain, storyResult)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Summary:")
	fmt.Println(summaryResult)
	fmt.Println()

	// Example 2: Sequential Chain
	fmt.Println("Example 2: Sequential Chain")
	fmt.Println("Generate product name → Write description → Create tagline\n")

	// Chain 1: Product name
	namePrompt := prompts.NewPromptTemplate(
		"Generate a creative product name for: {{.product_type}}. Only respond with the name, nothing else.",
		[]string{"product_type"},
	)
	nameChain := chains.NewLLMChain(llm, namePrompt)

	productName, err := chains.Run(ctx, nameChain, "eco-friendly water bottle")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Step 1 - Product Name: %s\n", productName)

	// Chain 2: Product description
	descPrompt := prompts.NewPromptTemplate(
		"Write a brief 2-sentence product description for {{.product_name}}, an eco-friendly water bottle.",
		[]string{"product_name"},
	)
	descChain := chains.NewLLMChain(llm, descPrompt)

	productDesc, err := chains.Run(ctx, descChain, productName)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Step 2 - Description: %s\n", productDesc)

	// Chain 3: Marketing tagline
	taglinePrompt := prompts.NewPromptTemplate(
		"Create a catchy marketing tagline for this product:\nName: {{.name}}\nDescription: {{.description}}\n\nOnly respond with the tagline.",
		[]string{"name", "description"},
	)
	taglineChain := chains.NewLLMChain(llm, taglinePrompt)

	inputs := map[string]any{
		"name":        productName,
		"description": productDesc,
	}

	tagline, err := chains.Call(ctx, taglineChain, inputs, chains.WithMaxTokens(50))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	taglineText := tagline[taglineChain.OutputKey].(string)
	fmt.Printf("Step 3 - Tagline: %s\n", taglineText)

	// Example 3: Conversation Chain with Memory
	fmt.Println("\nExample 3: Conversation Chain")
	fmt.Println("Having a multi-turn conversation\n")

	conversationPrompt := prompts.NewPromptTemplate(
		"You are a helpful assistant. {{.history}}\n\nHuman: {{.input}}\nAssistant:",
		[]string{"history", "input"},
	)

	conversationChain := chains.NewLLMChain(llm, conversationPrompt)

	// Simulate conversation history
	history := ""
	questions := []string{
		"What is the capital of France?",
		"What is it famous for?",
		"What's the best time to visit?",
	}

	for i, question := range questions {
		fmt.Printf("Turn %d - Human: %s\n", i+1, question)

		inputs := map[string]any{
			"history": history,
			"input":   question,
		}

		result, err := chains.Call(ctx, conversationChain, inputs)
		if err != nil {
			log.Printf("Error: %v\n", err)
			return
		}

		response := result[conversationChain.OutputKey].(string)
		fmt.Printf("Turn %d - Assistant: %s\n\n", i+1, response)

		// Update history
		history += fmt.Sprintf("Human: %s\nAssistant: %s\n", question, response)
	}
}
