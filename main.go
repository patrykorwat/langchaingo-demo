package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/patrykorwat/langchaingo-demo/examples"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printMenu()
		fmt.Print("\nSelect an example (1-8, or 'q' to quit): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		if input == "q" || input == "Q" {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Println("\n" + strings.Repeat("=", 80))

		switch input {
		case "1":
			examples.RunBasicLLM()
		case "2":
			examples.RunChains()
		case "3":
			examples.RunPromptTemplates()
		case "4":
			examples.RunMemory()
		case "5":
			examples.RunAgents()
		case "6":
			examples.RunDocumentProcessing()
		case "7":
			examples.RunOutputParsers()
		case "8":
			examples.RunStreamingExample()
		default:
			fmt.Println("Invalid selection. Please try again.")
		}

		fmt.Println(strings.Repeat("=", 80) + "\n")
	}
}

func printMenu() {
	fmt.Println("\n╔══════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                   LangChain Go - Feature Demonstrations                     ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("  1. Basic LLM - Simple text completion with Claude")
	fmt.Println("  2. Chains - Sequential operations and LLM chains")
	fmt.Println("  3. Prompt Templates - Reusable prompts with variables")
	fmt.Println("  4. Memory - Conversation history and context management")
	fmt.Println("  5. Agents & Tools - Autonomous decision-making with custom tools")
	fmt.Println("  6. Document Processing - Text splitting and chunking")
	fmt.Println("  7. Output Parsers - Structured output from LLMs")
	fmt.Println("  8. Streaming - Real-time streaming responses")
	fmt.Println()
	fmt.Println("  Q. Quit")
}
