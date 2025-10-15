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

	// Example 1: Simple Template
	fmt.Println("Example 1: Simple Template with Single Variable")
	fmt.Println("Template: 'Explain {{.concept}} to a 5-year-old'\n")

	simpleTemplate := prompts.NewPromptTemplate(
		"Explain {{.concept}} to a 5-year-old in 2-3 sentences.",
		[]string{"concept"},
	)

	chain := chains.NewLLMChain(llm, simpleTemplate)

	concepts := []string{"gravity", "photosynthesis", "democracy"}
	for _, concept := range concepts {
		result, err := chains.Run(ctx, chain, concept)
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Concept: %s\n%s\n\n", concept, result)
	}

	// Example 2: Multiple Variables
	fmt.Println("Example 2: Template with Multiple Variables")
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
	fmt.Printf("Generated Email:\n%s\n\n", emailText)

	// Example 3: Few-Shot Prompting
	fmt.Println("Example 3: Few-Shot Prompting")
	fmt.Println("Teaching the LLM a pattern through examples\n")

	fewShotTemplate := prompts.NewPromptTemplate(
		`Convert the following text to a professional tone:

Example 1:
Input: "hey can u help me with this?"
Output: "Hello, could you please assist me with this matter?"

Example 2:
Input: "gonna be late tmrw"
Output: "I will be arriving late tomorrow."

Example 3:
Input: "thx for ur help!"
Output: "Thank you for your assistance."

Now convert this:
Input: "{{.input}}"
Output:`,
		[]string{"input"},
	)

	fewShotChain := chains.NewLLMChain(llm, fewShotTemplate)

	casualTexts := []string{
		"wanna grab lunch?",
		"sorry cant make it",
		"great job on the presentation!",
	}

	for _, text := range casualTexts {
		result, err := chains.Run(ctx, fewShotChain, text)
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Casual: %s\nProfessional: %s\n\n", text, result)
	}

	// Example 4: Chat Prompt Template
	fmt.Println("Example 4: Role-Based Prompt")
	fmt.Println("Setting system context and user message\n")

	roleTemplate := prompts.NewPromptTemplate(
		`You are a {{.role}}. {{.context}}

User question: {{.question}}

Provide a helpful response in {{.style}} style.`,
		[]string{"role", "context", "question", "style"},
	)

	roleChain := chains.NewLLMChain(llm, roleTemplate)

	roleInputs := map[string]any{
		"role":     "senior software architect",
		"context":  "You have 15 years of experience in distributed systems.",
		"question": "What are the key considerations when designing a microservices architecture?",
		"style":    "concise and practical",
	}

	result, err = chains.Call(ctx, roleChain, roleInputs)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	roleResponse := result[roleChain.OutputKey].(string)
	fmt.Printf("Question: %s\n\nResponse:\n%s\n", roleInputs["question"], roleResponse)

	// Example 5: Template Composition
	fmt.Println("\nExample 5: Conditional Template")
	fmt.Println("Different prompts based on difficulty level\n")

	difficulties := map[string]string{
		"beginner": "Explain {{.topic}} using simple everyday language and examples that anyone can understand.",
		"advanced": "Provide a technical deep-dive into {{.topic}}, including architectural considerations and implementation details.",
	}

	topic := "REST API design"

	for level, templateStr := range difficulties {
		template := prompts.NewPromptTemplate(templateStr, []string{"topic"})
		levelChain := chains.NewLLMChain(llm, template)

		result, err := chains.Run(ctx, levelChain, topic)
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("Level: %s\n%s\n\n", level, result)
	}
}
