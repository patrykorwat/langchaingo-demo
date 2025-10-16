package examples

import (
	"fmt"

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
2. Chains: Chains go beyond just a single LLM call and are sequences of calls.
3. Data Augmented Generation: Involves specific types of chains that first interact with an external datasource.
4. Agents: Agents involve an LLM making decisions about which Actions to take.
5. Memory: Memory is the concept of persisting state between calls of a chain/agent.

LangChain provides a standard interface through which you can interact with many different types of LLMs.`

	// Character-based Text Splitting
	fmt.Println("Character-based Text Splitting")
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
}
