package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/agenticgokit/agenticgokit/plugins/llm/ollama"
	"github.com/agenticgokit/agenticgokit/v1beta"
)

func main() {
	fmt.Println("===========================================")
	fmt.Println("  Ollama Short Answer Agent - v1beta API")
	fmt.Println("===========================================\n")

	// Create a simple chat agent using Ollama with short, concise responses
	agent, err := createShortAnswerAgent()
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	// Initialize the agent
	ctx := context.Background()
	if err := agent.Initialize(ctx); err != nil {
		log.Fatalf("Failed to initialize agent: %v", err)
	}
	defer agent.Cleanup(ctx)

	// Run example queries
	queries := []string{
		"What is 2+29?",
		"Explain what Docker is.",
	}

	for i, query := range queries {
		fmt.Printf("[Query %d] %s\n", i+1, query)
		fmt.Println("---")

		// Create context with timeout
		queryCtx, cancel := context.WithTimeout(ctx, 30*time.Second)

		// Run the agent
		result, err := agent.Run(queryCtx, query)
		if err != nil {
			fmt.Printf("❌ Error: %v\n\n", err)
			cancel()
			continue
		}

		// Display the result
		fmt.Printf("✓ Answer: %s\n", result.Content)
		fmt.Printf("   Duration: %v\n", result.Duration)
		fmt.Printf("   Success: %v\n\n", result.Success)

		cancel()
	}

	fmt.Println("===========================================")
	fmt.Println("  Demo completed successfully!")
	fmt.Println("===========================================")
}

// createShortAnswerAgent creates an Ollama-based agent that provides short, concise answers
func createShortAnswerAgent() (v1beta.Agent, error) {
	systemPrompt := `You are a helpful AI assistant that provides short, concise answers.
Keep your responses to 2-3 sentences maximum.
Be direct and to the point.
Do not provide long explanations unless specifically asked.`

	config := &v1beta.Config{
		Name:         "short-answer-agent",
		SystemPrompt: systemPrompt,
		Timeout:      30 * time.Second,
		LLM: v1beta.LLMConfig{
			Provider:    "ollama",
			Model:       "gemma3:1b",
			Temperature: 0.3,
			MaxTokens:   200,
		},
	}

	agent, err := v1beta.NewBuilder("short-answer-agent").
		WithConfig(config).
		Build()

	if err != nil {
		return nil, fmt.Errorf("failed to build agent: %w", err)
	}

	return agent, nil
}
