# Ollama Short Answer Agent - v1beta API Example

This example demonstrates how to create a simple, single-agent application using the AgenticGoKit v1beta public APIs. The agent uses Ollama as the LLM provider and is configured to provide short, concise answers to user queries.

## Features

- ✅ Uses **v1beta public APIs** (Builder pattern)
- ✅ **Ollama integration** for local LLM execution
- ✅ **Short answer optimization** with system prompts and token limits
- ✅ **Clean error handling** and timeout management
- ✅ **Simple, readable code** for learning purposes

## Current Status

✅ **Implementation Complete** - Full LLM integration working  
✅ **Compiles Successfully** - All code is syntactically correct

## Prerequisites

1. **Go 1.24+** installed
2. **Ollama** installed and running
3. **gemma3:1b model** pulled in Ollama

### Install Ollama

```bash
# macOS/Linux
curl -fsSL https://ollama.com/install.sh | sh

# Windows
# Download from https://ollama.com/download
```

### Pull the Model

```bash
ollama pull gemma3:1b
```

Verify Ollama is running:
```bash
curl http://localhost:11434/api/tags
```

## Project Structure

```
ollama-short-answer/
├── main.go           # Main application with agent setup and execution
├── go.mod            # Go module file
└── README.md         # This file
```

## Code Walkthrough

### 1. Agent Configuration

The agent is configured with:
- **System Prompt**: Instructs the LLM to provide short, 2-3 sentence answers
- **Low Temperature** (0.3): More focused and deterministic responses
- **Limited Tokens** (200): Enforces brevity
- **Ollama Provider**: Uses local gemma3:1b model

```go
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
```

### 2. Builder Pattern

The agent is built using the v1beta Builder pattern:

```go
agent, err := v1beta.NewBuilder("short-answer-agent").
    WithConfig(config).
    Build()
```


### 3. Agent Execution

The agent is initialized, runs queries, and cleaned up properly:

```go
ctx := context.Background()
agent.Initialize(ctx)
defer agent.Cleanup(ctx)

result, err := agent.Run(ctx, query)
```

## Running the Example

### Option 1: Direct Execution

```bash
cd examples/ollama-short-answer
go run main.go
```

### Option 2: Build and Run

```bash
cd examples/ollama-short-answer
go build -o ollama-agent
./ollama-agent
```

## Expected Output

```
==========================================
  Ollama Short Answer Agent - v1beta API
==========================================

[Query 1] What is 2+29?
---
✓ Answer: 31
   Duration: 1.2s
   Success: true

[Query 2] Explain what Docker is.
---
✓ Answer: Docker is a platform for developing, shipping, and running applications in containers.
   Duration: 1.1s
   Success: true

...
```

## Key v1beta APIs Used

### Agent Interface
- `agent.Run(ctx, input)` - Execute agent with input
- `agent.Initialize(ctx)` - Initialize agent resources
- `agent.Cleanup(ctx)` - Clean up agent resources

### Builder Pattern
- `v1beta.NewBuilder(name)` - Create new agent builder
- `WithConfig(config)` - Set complete configuration
- `Build()` - Build the final agent

### Configuration Types
- `v1beta.Config` - Main agent configuration
- `v1beta.LLMConfig` - LLM provider settings

### Result Type
- `result.Content` - Agent response text
- `result.Duration` - Execution duration
- `result.Success` - Success status

## Customization

### Change the Model

```go
config.LLM.Model = "llama3.2"  // or "mistral", "gemma3:1b", etc.
```

### Adjust Response Length

```go
config.LLM.MaxTokens = 500  // Longer responses
config.LLM.Temperature = 0.7  // More creative
```

### Custom System Prompt

```go
systemPrompt := "You are an expert in [TOPIC]. Provide detailed explanations..."
```

### Add Timeout

```go
queryCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
defer cancel()
```

## Troubleshooting

### Ollama Not Running
```
Error: failed to connect to Ollama
Solution: Start Ollama service
```

### Model Not Found
```
Error: model 'gemma3:1b' not found
Solution: Run 'ollama pull gemma3:1b'
```

### Timeout Errors
```
Error: context deadline exceeded
Solution: Increase timeout in config or query context
```

## Next Steps

- **Add Streaming**: Use `agent.RunStream()` for real-time responses
- **Add Memory**: Enable conversation history with `WithMemory()`
- **Add Tools**: Integrate external tools with `WithTools()`
- **Configuration File**: Load settings from TOML file


## References

- [v1beta Documentation](../../../v1beta/)
- [Builder Pattern Guide](../../../v1beta/builder.go)
- [Configuration Guide](../../../v1beta/config.go)
- [Ollama Documentation](https://github.com/ollama/ollama)
