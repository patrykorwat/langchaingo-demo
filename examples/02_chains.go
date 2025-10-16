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

	// Sequential Chain: Generate product name → Write description → Create tagline
	fmt.Println("Sequential Chain Example")
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
}
