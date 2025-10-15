package examples

import (
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/textsplitter"
)

// RunDocumentProcessing demonstrates text splitting and document chunking
func RunDocumentProcessing() {
	fmt.Println("📄 Document Processing Example")
	fmt.Println("Demonstrates text splitting and chunking strategies\n")

	// Sample long text
	sampleText := `LangChain is a framework for developing applications powered by language models.
It enables applications that are context-aware and can reason about complex tasks.

The framework consists of several key components:
1. LLMs and Prompts: This includes prompt management, prompt optimization, and a generic interface for all LLMs.
2. Chains: Chains go beyond just a single LLM call and are sequences of calls (whether to an LLM or a different utility).
3. Data Augmented Generation: Data Augmented Generation involves specific types of chains that first interact with an external datasource to fetch data to use in the generation step.
4. Agents: Agents involve an LLM making decisions about which Actions to take, taking that Action, seeing an Observation, and repeating that until done.
5. Memory: Memory is the concept of persisting state between calls of a chain/agent.

LangChain provides a standard interface through which you can interact with many different types of LLMs.
It also provides a set of utilities for working with these LLMs, including prompt templates, output parsers, and more.

The framework is designed to be modular and extensible, allowing developers to easily swap out components or add new ones.
This makes it easy to experiment with different approaches and find the best solution for your specific use case.`

	// Example 1: Character-based Text Splitting
	fmt.Println("Example 1: Character-based Text Splitting")
	fmt.Println("Split text by character count with overlap\n")

	charSplitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(200),
		textsplitter.WithChunkOverlap(50),
	)

	chunks, err := charSplitter.SplitText(sampleText)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total chunks: %d\n", len(chunks))
	for i, chunk := range chunks {
		fmt.Printf("\n--- Chunk %d (length: %d) ---\n%s\n", i+1, len(chunk), chunk)
	}

	// Example 2: Token-based Text Splitting
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Example 2: Token-based Text Splitting")
	fmt.Println("Split by token count for LLM context windows\n")

	tokenSplitter := textsplitter.NewTokenSplitter(
		textsplitter.WithChunkSize(100),
		textsplitter.WithChunkOverlap(20),
	)

	tokenChunks, err := tokenSplitter.SplitText(sampleText)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total token-based chunks: %d\n", len(tokenChunks))
	for i, chunk := range tokenChunks {
		fmt.Printf("\n--- Token Chunk %d ---\n%s\n", i+1, chunk)
	}

	// Example 3: Markdown Text Splitting
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Example 3: Markdown-aware Text Splitting")
	fmt.Println("Respects markdown structure when splitting\n")

	markdownText := `# LangChain Documentation

## Introduction
LangChain is a framework for developing applications powered by language models.

## Core Concepts

### LLMs
Large Language Models are the foundation of LangChain applications.

### Chains
Chains combine multiple components together.

### Agents
Agents use LLMs to decide which actions to take.

## Getting Started

### Installation
Install LangChain using your package manager.

### Basic Usage
Here's a simple example to get you started.`

	markdownSplitter := textsplitter.NewMarkdownTextSplitter(
		textsplitter.WithChunkSize(150),
		textsplitter.WithChunkOverlap(20),
	)

	mdChunks, err := markdownSplitter.SplitText(markdownText)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total markdown chunks: %d\n", len(mdChunks))
	for i, chunk := range mdChunks {
		fmt.Printf("\n--- Markdown Chunk %d ---\n%s\n", i+1, chunk)
	}

	// Example 4: Code-aware Text Splitting
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Example 4: Different Chunk Sizes")
	fmt.Println("Comparing small vs large chunk sizes\n")

	codeText := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

func processData(data []int) []int {
    result := make([]int, len(data))
    for i, v := range data {
        result[i] = v * 2
    }
    return result
}`

	// Small chunks
	smallSplitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(50),
		textsplitter.WithChunkOverlap(10),
	)

	smallChunks, err := smallSplitter.SplitText(codeText)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Small chunks (size 50): %d chunks\n", len(smallChunks))

	// Large chunks
	largeSplitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(200),
		textsplitter.WithChunkOverlap(20),
	)

	largeChunks, err := largeSplitter.SplitText(codeText)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Large chunks (size 200): %d chunks\n\n", len(largeChunks))

	// Example 5: Custom Separators
	fmt.Println("Example 5: Custom Separators")
	fmt.Println("Split text using custom delimiters\n")

	csvText := `Name,Age,City
John,30,New York
Jane,25,Los Angeles
Bob,35,Chicago
Alice,28,Houston
Charlie,32,Phoenix`

	// Split by newlines
	lineSplitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(100),
		textsplitter.WithChunkOverlap(0),
		textsplitter.WithSeparators([]string{"\n", ","}),
	)

	csvChunks, err := lineSplitter.SplitText(csvText)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("CSV chunks: %d\n", len(csvChunks))
	for i, chunk := range csvChunks {
		fmt.Printf("Chunk %d: %s\n", i+1, chunk)
	}

	// Summary
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Summary:")
	fmt.Println("- Character splitting: Good for general text")
	fmt.Println("- Token splitting: Best for LLM context management")
	fmt.Println("- Markdown splitting: Preserves document structure")
	fmt.Println("- Custom separators: Flexible for specific formats")
	fmt.Println("- Chunk overlap: Maintains context between chunks")
}
