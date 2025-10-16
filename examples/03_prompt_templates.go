package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/prompts"
)

// RunPromptTemplates demonstrates various ways to use prompt templates
func RunPromptTemplates() {
	fmt.Println("📝 Prompt Templates Example")
	fmt.Println("Demonstrates reusable prompts with variable substitution\n")

	// Initialize the LLM
	llm, err := anthropic.New(
		anthropic.WithModel("claude-sonnet-4-5-20250929"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Template with Multiple Variables
	fmt.Println("Template with Multiple Variables")
	fmt.Println("Email generator with tone, recipient, and topic\n")

	emailTemplate := prompts.NewPromptTemplate(
		`Write a {{.tone}} email to {{.recipient}} about {{.topic}}.
Keep it brief (3-4 sentences) and professional.`,
		[]string{"tone", "recipient", "topic"},
	)

	emailChain := chains.NewLLMChain(llm, emailTemplate)

	emailInputs := map[string]any{
		"tone":      "friendly and enthusiastic",
		"recipient": "the team",
		"topic":     "the successful completion of our project milestone",
	}

	result, err := chains.Call(ctx, emailChain, emailInputs)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	emailText := result[emailChain.OutputKey].(string)
	fmt.Printf("Generated Email:\n%s\n", emailText)
}
