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
	fmt.Println("An agent that can perform calculations\n")

	calculatorTool := tools.Calculator{}

	fmt.Println("Tool: Calculator")
	fmt.Println("Testing: 25 * 4 + 10")
	result, err := calculatorTool.Call(ctx, "25 * 4 + 10")
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n\n", result)
	}

	// Example 2: Custom String Tool
	fmt.Println("Example 2: Custom String Manipulation Tool")

	stringTool := &StringTool{}

	testCases := []string{
		"uppercase:hello world",
		"lowercase:HELLO WORLD",
		"reverse:abcdef",
		"length:testing",
	}

	for _, testCase := range testCases {
		result, err := stringTool.Call(ctx, testCase)
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Input: %s -> Output: %s\n", testCase, result)
	}
	fmt.Println()

	// Example 3: Math Tools
	fmt.Println("Example 3: Custom Math Tools")
	fmt.Println("Tools for square root, power, and absolute value\n")

	mathTools := []tools.Tool{
		&SquareRootTool{},
		&PowerTool{},
		&AbsoluteTool{},
	}

	for _, tool := range mathTools {
		fmt.Printf("Tool: %s\n", tool.Name())
		fmt.Printf("Description: %s\n", tool.Description())
	}
	fmt.Println()

	// Test square root
	sqrtTool := &SquareRootTool{}
	result, err = sqrtTool.Call(ctx, "16")
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("sqrt(16) = %s\n", result)
	}

	// Test power
	powerTool := &PowerTool{}
	result, err = powerTool.Call(ctx, "2,10")
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("power(2, 10) = %s\n", result)
	}

	// Test absolute
	absTool := &AbsoluteTool{}
	result, err = absTool.Call(ctx, "-42")
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("abs(-42) = %s\n\n", result)
	}

	// Example 4: Agent with Tools
	fmt.Println("Example 4: Agent with Multiple Tools")
	fmt.Println("Agent can choose which tool to use\n")

	allTools := []tools.Tool{
		tools.Calculator{},
		&StringTool{},
		&SquareRootTool{},
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
		"Convert the text 'hello agent' to uppercase",
		"What is the square root of 144?",
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

// PowerTool calculates power
type PowerTool struct{}

func (t *PowerTool) Name() string {
	return "Power"
}

func (t *PowerTool) Description() string {
	return "Calculates base raised to exponent. Input format: base,exponent (e.g., 2,10)"
}

func (t *PowerTool) Call(ctx context.Context, input string) (string, error) {
	parts := strings.Split(strings.TrimSpace(input), ",")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid input format, use base,exponent")
	}

	base, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return "", fmt.Errorf("invalid base: %v", err)
	}

	exponent, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return "", fmt.Errorf("invalid exponent: %v", err)
	}

	result := math.Pow(base, exponent)
	return fmt.Sprintf("%.2f", result), nil
}

// AbsoluteTool calculates absolute value
type AbsoluteTool struct{}

func (t *AbsoluteTool) Name() string {
	return "Absolute"
}

func (t *AbsoluteTool) Description() string {
	return "Calculates the absolute value of a number. Input should be a single number."
}

func (t *AbsoluteTool) Call(ctx context.Context, input string) (string, error) {
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return "", fmt.Errorf("invalid number: %v", err)
	}

	result := math.Abs(num)
	return fmt.Sprintf("%.2f", result), nil
}
