# LangChain Go Demo

A comprehensive demonstration project showcasing the main features and capabilities of [langchaingo](https://github.com/tmc/langchaingo), a Go implementation of LangChain for building applications with Large Language Models (LLMs).

## Features

This demo includes interactive examples of:

1. **Basic LLM Usage** - Simple text completion with various parameters
2. **Chains** - Sequential operations and composition of LLM calls
3. **Prompt Templates** - Reusable prompts with variable substitution
4. **Memory** - Conversation history and context management
5. **Agents & Tools** - Autonomous decision-making with custom tools
6. **Document Processing** - Text splitting and chunking strategies
7. **Output Parsers** - Extracting structured data from LLM responses
8. **Streaming** - Real-time streaming responses

## Prerequisites

- Go 1.24.2 or higher
- Anthropic API key (for Claude)

## Installation

1. Clone this repository:
```bash
git clone https://github.com/patrykorwat/langchaingo-demo.git
cd langchaingo-demo
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your Anthropic API key:
```bash
export ANTHROPIC_API_KEY="your-api-key-here"
```

## Usage

Run the demo:
```bash
go run main.go
```

You'll see an interactive menu where you can select different examples to run:

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                   LangChain Go - Feature Demonstrations                     ║
╚══════════════════════════════════════════════════════════════════════════════╝

  1. Basic LLM - Simple text completion with Claude
  2. Chains - Sequential operations and LLM chains
  3. Prompt Templates - Reusable prompts with variables
  4. Memory - Conversation history and context management
  5. Agents & Tools - Autonomous decision-making with custom tools
  6. Document Processing - Text splitting and chunking
  7. Output Parsers - Structured output from LLMs
  8. Streaming - Real-time streaming responses

  Q. Quit
```

## Example Descriptions

### 1. Basic LLM

Demonstrates fundamental LLM interactions:
- Simple text completion
- Temperature control (creativity vs determinism)
- Max tokens limitation
- Stop words

**Key concepts:** LLM initialization, prompt engineering, parameter tuning

### 2. Chains

Shows how to chain multiple LLM calls together:
- Sequential LLM chains
- Multi-step workflows
- Chain composition
- Conversation chains with history

**Key concepts:** Chain creation, sequential processing, data flow between chains

### 3. Prompt Templates

Explores reusable prompt patterns:
- Single variable templates
- Multiple variable substitution
- Few-shot prompting
- Role-based prompts
- Conditional templates

**Key concepts:** Template design, variable injection, prompt reusability

### 4. Memory

Demonstrates conversation context management:
- Conversation buffer memory (full history)
- Conversation window memory (recent K interactions)
- Manual memory management
- Context building

**Key concepts:** State persistence, conversation tracking, context windows

### 5. Agents & Tools

Showcases autonomous decision-making:
- Custom tool creation
- Calculator tool
- String manipulation tools
- Math tools (square root, power, absolute value)
- Agent executors with multiple tools

**Key concepts:** Tool interfaces, agent initialization, autonomous reasoning

### 6. Document Processing

Covers text splitting strategies:
- Character-based splitting
- Token-based splitting
- Markdown-aware splitting
- Custom separators
- Chunk overlap for context preservation

**Key concepts:** Text chunking, document loaders, context management

### 7. Output Parsers

Shows how to extract structured data:
- JSON parsing
- List extraction
- Key-value pairs
- Boolean classification
- Complex nested structures
- Markdown parsing

**Key concepts:** Output formatting, data extraction, structured responses

### 8. Streaming

Demonstrates real-time response handling:
- Basic streaming
- Character/word counting during stream
- Timing metrics (time to first token)
- Streaming vs non-streaming comparison
- Custom stream processing

**Key concepts:** Real-time responses, progressive loading, user experience optimization

## Project Structure

```
langchaingo-demo/
├── main.go                           # Main entry point with interactive menu
├── examples/
│   ├── 01_basic_llm.go              # Basic LLM operations
│   ├── 02_chains.go                  # Chain compositions
│   ├── 03_prompt_templates.go        # Template examples
│   ├── 04_memory.go                  # Memory management
│   ├── 05_agents.go                  # Agents and tools
│   ├── 06_document_processing.go     # Text splitting
│   ├── 07_output_parsers.go          # Structured output
│   └── 08_streaming.go               # Streaming responses
├── go.mod
├── go.sum
└── README.md
```

## Key Dependencies

- `github.com/tmc/langchaingo` - Main LangChain Go library
- `github.com/tmc/langchaingo/llms/anthropic` - Anthropic/Claude integration
- `github.com/tmc/langchaingo/chains` - Chain compositions
- `github.com/tmc/langchaingo/prompts` - Prompt templates
- `github.com/tmc/langchaingo/memory` - Memory management
- `github.com/tmc/langchaingo/agents` - Agent framework
- `github.com/tmc/langchaingo/tools` - Tool interfaces
- `github.com/tmc/langchaingo/textsplitter` - Document processing

## Learning Path

Recommended order for exploring the examples:

1. **Start with Basic LLM** to understand fundamental interactions
2. **Move to Prompt Templates** to learn about reusable prompts
3. **Explore Chains** to see how to compose multiple operations
4. **Try Memory** to understand conversation context
5. **Experiment with Output Parsers** to extract structured data
6. **Dive into Document Processing** for working with long texts
7. **Check out Streaming** for real-time responses
8. **Finally, Agents** to see autonomous decision-making

## Use Cases

This demo is useful for:

- **Learning LangChain Go**: Interactive examples of all major features
- **Prototyping**: Quick reference for implementing LLM features
- **Teaching**: Educational resource for LLM application development
- **Experimentation**: Sandbox for testing different approaches

## Common Patterns

### Basic LLM Call
```go
llm, _ := anthropic.New(anthropic.WithModel("claude-sonnet-4-5-20250929"))
response, _ := llms.GenerateFromSinglePrompt(ctx, llm, "Your prompt here")
```

### Using Chains
```go
prompt := prompts.NewPromptTemplate("Template: {{.variable}}", []string{"variable"})
chain := chains.NewLLMChain(llm, prompt)
result, _ := chains.Run(ctx, chain, "value")
```

### Streaming Responses
```go
llm.GenerateContent(ctx, []llms.MessageContent{
    llms.TextParts(llms.ChatMessageTypeHuman, "prompt"),
}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
    fmt.Print(string(chunk))
    return nil
}))
```

## Configuration

### Environment Variables

- `ANTHROPIC_API_KEY` - Your Anthropic API key (required)
- `ANTHROPIC_BASE_URL` - Custom API endpoint (optional)

### Model Configuration

The demo uses `claude-sonnet-4-5-20250929` by default. You can modify the model in each example file by changing:

```go
llm, err := anthropic.New(
    anthropic.WithModel("your-preferred-model"),
)
```

## Troubleshooting

### API Key Issues
- Ensure `ANTHROPIC_API_KEY` is set in your environment
- Verify your API key is valid and has sufficient credits

### Import Errors
- Run `go mod tidy` to ensure all dependencies are installed
- Check that you're using Go 1.24.2 or higher

### Rate Limits
- The examples make multiple API calls
- Add delays between requests if you hit rate limits
- Consider using a lower-tier model for testing

## Contributing

This is a demo project. Feel free to:
- Add new examples
- Improve existing demonstrations
- Add support for other LLM providers
- Enhance documentation

## Resources

- [LangChain Go Repository](https://github.com/tmc/langchaingo)
- [LangChain Documentation](https://docs.langchain.com/)
- [Anthropic Claude Documentation](https://docs.anthropic.com/)
- [Go Programming Language](https://go.dev/)

## License

This demo project is provided as-is for educational purposes.

## Acknowledgments

- [tmc/langchaingo](https://github.com/tmc/langchaingo) - The excellent LangChain Go implementation
- [Anthropic](https://www.anthropic.com/) - For the Claude API
- The LangChain community for patterns and best practices
