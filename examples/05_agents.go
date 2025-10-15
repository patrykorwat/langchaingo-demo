package examples

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/tools"
)

// RunAgents demonstrates agents with custom tools
func RunAgents() {
	fmt.Println("🤖 Agents & Tools Example")
	fmt.Println("Demonstrates autonomous decision-making with custom tools\n")

	// Initialize the LLM
	llm, err := anthropic.New(
		anthropic.WithModel("claude-sonnet-4-5-20250929"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Example 1: Custom Calculator Tool
	fmt.Println("Example 1: Calculator Tool")
	calculatorTool := tools.Calculator{}

	fmt.Println("Testing: 25 * 4 + 10")
	result, err := calculatorTool.Call(ctx, "25 * 4 + 10")
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n\n", result)
	}

	// Example 2: Agent with Multiple Tools
	fmt.Println("Example 2: Agent with Multiple Tools")
	fmt.Println("Agent can choose which tool to use\n")

	allTools := []tools.Tool{
		tools.Calculator{},
		&StringTool{},
	}

	// Create an agent executor
	executor, err := agents.Initialize(
		llm,
		allTools,
		agents.ZeroShotReactDescription,
		agents.WithMaxIterations(3),
	)
	if err != nil {
		log.Printf("Error creating agent: %v\n", err)
		return
	}

	// Test the agent with different queries
	queries := []string{
		"What is 15 + 27?",
		"Convert the text 'hello' to uppercase",
	}

	for _, query := range queries {
		fmt.Printf("Query: %s\n", query)

		result, err := chains.Run(ctx, executor, query)
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("Answer: %s\n\n", result)
	}
}

// StringTool is a custom tool for string manipulation
type StringTool struct{}

func (t *StringTool) Name() string {
	return "StringManipulation"
}

func (t *StringTool) Description() string {
	return `Useful for manipulating strings.
Input format: operation:text
Operations: uppercase, lowercase, reverse, length
Example: uppercase:hello`
}

func (t *StringTool) Call(ctx context.Context, input string) (string, error) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid input format, use operation:text")
	}

	operation := strings.ToLower(strings.TrimSpace(parts[0]))
	text := parts[1]

	switch operation {
	case "uppercase":
		return strings.ToUpper(text), nil
	case "lowercase":
		return strings.ToLower(text), nil
	case "reverse":
		runes := []rune(text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	case "length":
		return strconv.Itoa(len(text)), nil
	default:
		return "", fmt.Errorf("unknown operation: %s", operation)
	}
}

// SquareRootTool calculates square root
type SquareRootTool struct{}

func (t *SquareRootTool) Name() string {
	return "SquareRoot"
}

func (t *SquareRootTool) Description() string {
	return "Calculates the square root of a number. Input should be a single number."
}

func (t *SquareRootTool) Call(ctx context.Context, input string) (string, error) {
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return "", fmt.Errorf("invalid number: %v", err)
	}

	if num < 0 {
		return "", fmt.Errorf("cannot calculate square root of negative number")
	}

	result := math.Sqrt(num)
	return fmt.Sprintf("%.2f", result), nil
}
