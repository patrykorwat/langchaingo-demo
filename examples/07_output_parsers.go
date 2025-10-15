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

	personText := "John Smith is a 35-year-old software engineer who enjoys hiking, photography, and playing guitar in his free time."
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

	// Example 3: Key-Value Pairs
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Example 3: Key-Value Pair Parsing")
	fmt.Println("Extract structured information as key-value pairs\n")

	kvPrompt := prompts.NewPromptTemplate(
		`Analyze the following product and extract these attributes as key-value pairs (one per line, format: Key: Value):
- Brand
- Model
- Price
- Color

Product: {{.product}}`,
		[]string{"product"},
	)

	kvChain := chains.NewLLMChain(llm, kvPrompt)

	productResult, err := chains.Run(ctx, kvChain, "iPhone 15 Pro in titanium blue for $999")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Raw output:\n%s\n\n", productResult)

	// Parse key-value pairs
	lines := strings.Split(productResult, "\n")
	attributes := make(map[string]string)

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			attributes[key] = value
		}
	}

	fmt.Println("Parsed attributes:")
	for key, value := range attributes {
		fmt.Printf("  %s: %s\n", key, value)
	}

	// Example 4: Boolean Output
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Example 4: Boolean Classification")
	fmt.Println("Parse yes/no or true/false responses\n")

	boolPrompt := prompts.NewPromptTemplate(
		`Is the following statement a question?
Statement: {{.statement}}

Answer only with 'yes' or 'no'.`,
		[]string{"statement"},
	)

	boolChain := chains.NewLLMChain(llm, boolPrompt)

	statements := []string{
		"What time is it?",
		"The sky is blue.",
		"How do I install this package?",
		"Programming is fun.",
	}

	for _, stmt := range statements {
		result, err := chains.Run(ctx, boolChain, stmt)
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		isQuestion := strings.ToLower(strings.TrimSpace(result)) == "yes"
		fmt.Printf("'%s' -> Is question? %v\n", stmt, isQuestion)
	}

	// Example 5: Structured JSON with Multiple Objects
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Example 5: Complex JSON Structure")
	fmt.Println("Parse a list of objects\n")

	complexPrompt := prompts.NewPromptTemplate(
		`Extract information about movies from the text and return as a JSON array of objects.
Each object should have: title, year, director

Text: {{.text}}

Return only the JSON array.`,
		[]string{"text"},
	)

	complexChain := chains.NewLLMChain(llm, complexPrompt)

	moviesText := "The Shawshank Redemption (1994) directed by Frank Darabont and The Godfather (1972) directed by Francis Ford Coppola are considered masterpieces."
	moviesResult, err := chains.Run(ctx, complexChain, moviesText)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Raw output:\n%s\n\n", moviesResult)

	// Parse the JSON array
	type Movie struct {
		Title    string `json:"title"`
		Year     int    `json:"year"`
		Director string `json:"director"`
	}

	// Clean the output
	cleanMoviesJSON := strings.TrimSpace(moviesResult)
	cleanMoviesJSON = strings.TrimPrefix(cleanMoviesJSON, "```json")
	cleanMoviesJSON = strings.TrimPrefix(cleanMoviesJSON, "```")
	cleanMoviesJSON = strings.TrimSuffix(cleanMoviesJSON, "```")
	cleanMoviesJSON = strings.TrimSpace(cleanMoviesJSON)

	var movies []Movie
	if err := json.Unmarshal([]byte(cleanMoviesJSON), &movies); err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
	} else {
		fmt.Printf("Parsed Movies (%d total):\n", len(movies))
		for i, movie := range movies {
			fmt.Printf("  %d. %s (%d) - directed by %s\n", i+1, movie.Title, movie.Year, movie.Director)
		}
	}

	// Example 6: Markdown Output
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Example 6: Markdown Structure Parsing")
	fmt.Println("Parse markdown formatted output\n")

	mdPrompt := prompts.NewPromptTemplate(
		`Create a brief comparison of {{.topic1}} vs {{.topic2}}.
Format your response as markdown with sections:
- ## Overview
- ## Pros of {{.topic1}}
- ## Pros of {{.topic2}}
- ## Conclusion`,
		[]string{"topic1", "topic2"},
	)

	mdChain := chains.NewLLMChain(llm, mdPrompt)

	mdInputs := map[string]any{
		"topic1": "REST APIs",
		"topic2": "GraphQL",
	}

	mdResult, err := chains.Call(ctx, mdChain, mdInputs)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	mdText := mdResult[mdChain.OutputKey].(string)
	fmt.Printf("Markdown output:\n%s\n", mdText)

	// Parse sections
	sections := strings.Split(mdText, "##")
	fmt.Printf("\nFound %d sections\n", len(sections)-1) // -1 because first split is before any ##
}
