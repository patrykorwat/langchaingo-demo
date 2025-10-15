package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/prompts"
)

// RunOutputParsers demonstrates parsing structured output from LLMs
func RunOutputParsers() {
	fmt.Println("🔍 Output Parsers Example")
	fmt.Println("Demonstrates parsing structured output from LLMs\n")

	// Initialize the LLM
	llm, err := anthropic.New(
		anthropic.WithModel("claude-sonnet-4-5-20250929"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Example 1: JSON Output
	fmt.Println("Example 1: JSON Output Parsing")
	fmt.Println("Request structured data in JSON format\n")

	jsonPrompt := prompts.NewPromptTemplate(
		`Extract information about the following person and return it as a valid JSON object with these fields:
- name (string)
- age (number)
- occupation (string)
- hobbies (array of strings)

Text: {{.text}}

Return only the JSON object, no additional text.`,
		[]string{"text"},
	)

	chain := chains.NewLLMChain(llm, jsonPrompt)

	personText := "John Smith is a 35-year-old software engineer who enjoys hiking, photography, and playing guitar."
	result, err := chains.Run(ctx, chain, personText)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Raw output:\n%s\n\n", result)

	// Parse the JSON
	type Person struct {
		Name       string   `json:"name"`
		Age        int      `json:"age"`
		Occupation string   `json:"occupation"`
		Hobbies    []string `json:"hobbies"`
	}

	// Clean the output (remove markdown code blocks if present)
	cleanJSON := strings.TrimSpace(result)
	cleanJSON = strings.TrimPrefix(cleanJSON, "```json")
	cleanJSON = strings.TrimPrefix(cleanJSON, "```")
	cleanJSON = strings.TrimSuffix(cleanJSON, "```")
	cleanJSON = strings.TrimSpace(cleanJSON)

	var person Person
	if err := json.Unmarshal([]byte(cleanJSON), &person); err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
	} else {
		fmt.Printf("Parsed Person:\n")
		fmt.Printf("  Name: %s\n", person.Name)
		fmt.Printf("  Age: %d\n", person.Age)
		fmt.Printf("  Occupation: %s\n", person.Occupation)
		fmt.Printf("  Hobbies: %v\n", person.Hobbies)
	}

	// Example 2: List Output
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Example 2: Comma-Separated List Parsing")
	fmt.Println("Extract items as a simple list\n")

	listPrompt := prompts.NewPromptTemplate(
		`List the main ingredients in {{.dish}}.
Return only the ingredients as a comma-separated list, nothing else.`,
		[]string{"dish"},
	)

	listChain := chains.NewLLMChain(llm, listPrompt)

	dishResult, err := chains.Run(ctx, listChain, "spaghetti carbonara")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Raw output: %s\n", dishResult)

	// Parse the list
	ingredients := strings.Split(dishResult, ",")
	fmt.Printf("\nParsed ingredients (%d items):\n", len(ingredients))
	for i, ingredient := range ingredients {
		fmt.Printf("  %d. %s\n", i+1, strings.TrimSpace(ingredient))
	}
}
