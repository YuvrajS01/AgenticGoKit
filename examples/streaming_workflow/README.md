# Multi-Agent Workflow with Streaming

This example demonstrates how to create multi-agent workflows using v1beta.Workflow with real-time streaming support. Watch as specialized agents collaborate in sequence, with each agent's output flowing seamlessly to the next.

## 🎯 What This Example Shows

**Sequential Workflow with Two Agents:**
1. **🔍 Research Agent** - Conducts comprehensive topic research  
2. **📝 Summarizer Agent** - Creates concise summaries from research output

**Key Features Demonstrated:**
- **Real-time Streaming**: Watch tokens appear as each agent generates them
- **Automatic Data Flow**: Research output automatically feeds into summarizer
- **Progress Tracking**: See step-by-step progress with metadata and timing
- **Agent Specialization**: Different agents with specialized system prompts
- **Workflow Orchestration**: Using v1beta.Workflow to manage the pipeline

## 🚀 Running the Example

### Prerequisites
1. **Ollama** installed and running:
   ```bash
   # Install Ollama (if not already installed)
   # Visit: https://ollama.ai/download
   
   # Pull required model  
   ollama pull gemma3:1b
   
   # Start Ollama (if not running)
   ollama serve
   ```

2. **Go environment** set up in the AgenticGoKit project

### Running the Workflow
```bash
cd examples/streaming_workflow
go run main.go
```

### What You'll See
- **Workflow Start**: Initial setup and connection testing
- **Step 1**: Research agent streams detailed analysis in real-time
- **Step 2**: Summarizer agent processes research and creates summary
- **Real-time Progress**: Token-by-token streaming with step indicators
- **Final Results**: Completion summary with timing and statistics

## 📋 Example Output

```
🚀 v1beta.Workflow Streaming Showcase
====================================
Demonstrating v1beta.Workflow streaming!

🔍 Testing Ollama connection...
✅ Ollama connection successful

🌟 v1beta.Workflow Sequential Streaming
=====================================
Using real v1beta.Workflow with streaming support!

🎯 Topic: Benefits of streaming in AI applications
🔄 Processing through workflow...

💬 Real-time Workflow Streaming:
─────────────────────────────────

📋 [WORKFLOW] Starting sequential workflow

🔄 [STEP: RESEARCH] Step 1/2: research
─────────────────────
Streaming is a really cool way to access content – like videos, music...
[Real-time tokens continue streaming...]

🔄 [STEP: SUMMARIZE] Step 2/2: summarize  
─────────────────────
Based on the research findings, here are the key points:
[Real-time summary tokens continue streaming...]

============================================================
🎉 v1beta.WORKFLOW STREAMING COMPLETED!
============================================================
✅ Success: true
⏱️ Duration: 75.98 seconds
📊 Total Chunks: 1084
📄 Final Output Length: 5424 characters

📋 Step Breakdown:
  🔸 Research: 4810 chars
  🔸 Summarize: 307 chars
```

## 🔍 Code Architecture

### Key Components

1. **v1beta.Workflow**: Orchestrates multi-agent sequences with streaming
2. **WorkflowStep**: Defines individual agents and their transformations
3. **Real-time Streaming**: Uses `workflow.RunStream()` to process tokens as they arrive
4. **Automatic Data Flow**: Output from one step becomes input to the next

### Workflow Setup
```go
// Create the workflow
workflow, err := v1beta.NewSequentialWorkflow(&v1beta.WorkflowConfig{
    Mode:    v1beta.Sequential,
    Timeout: 180 * time.Second,
})

// Add steps with agents and transformations
workflow.AddStep(v1beta.WorkflowStep{
    Name:  "research",
    Agent: researcherAgent,
    Transform: func(input string) string {
        return fmt.Sprintf("Research the topic: %s", input)
    },
})

workflow.AddStep(v1beta.WorkflowStep{
    Name:  "summarize",
    Agent: summarizerAgent,
    Transform: func(input string) string {
        return fmt.Sprintf("Please summarize this research:\n\n%s", input)
    },
})
```

### Streaming Execution
```go
// Run workflow with streaming
stream, err := workflow.RunStream(ctx, topic)
if err != nil {
    log.Fatalf("Workflow streaming failed: %v", err)
}

// Process streaming chunks
for chunk := range stream.Chunks() {
    switch chunk.Type {
    case v1beta.ChunkTypeMetadata:
        // Display step information
        if stepName, ok := chunk.Metadata["step_name"].(string); ok {
            fmt.Printf("🔄 [STEP: %s] %s\n", stepName, chunk.Content)
        }
    case v1beta.ChunkTypeDelta:
        // Display real-time tokens
        fmt.Print(chunk.Delta)
    case v1beta.ChunkTypeDone:
        fmt.Println("✅ Step completed!")
    }
}
```

## 💡 Why Streaming Matters for Workflows

### Without Streaming
```
User: "Research AI streaming benefits"
System: [Working... 75 seconds of silence]
System: [Complete results appear all at once]
```

### With v1beta.Workflow Streaming
```
User: "Research AI streaming benefits"
System: "🔄 [STEP: RESEARCH] Step 1/2: research"
System: "Streaming is a really cool way..." [tokens stream live]
System: "🔄 [STEP: SUMMARIZE] Step 2/2: summarize"
System: "Based on the research findings..." [tokens stream live]
System: "✅ v1beta.WORKFLOW STREAMING COMPLETED!"
```

### Benefits
1. **Real-time Feedback**: See progress as it happens
2. **Step Visibility**: Clear indication of which step is executing
3. **Automatic Data Flow**: Results flow seamlessly between steps
4. **Better Engagement**: Users stay engaged during long operations
5. **Early Assessment**: Evaluate output quality as it's generated

## 🛠 Customization

### Different LLM Providers
```go
// OpenAI
LLM: v1beta.LLMConfig{
    Provider: "openai",
    Model:    "gpt-4",
    APIKey:   os.Getenv("OPENAI_API_KEY"),
}

// Azure OpenAI
LLM: v1beta.LLMConfig{
    Provider: "azure",
    Model:    "gpt-4",
    BaseURL:  "https://your-resource.openai.azure.com/",
    APIKey:   os.Getenv("AZURE_OPENAI_KEY"),
}
```

### Adding More Steps
```go
// Add a third analysis step
workflow.AddStep(v1beta.WorkflowStep{
    Name:  "analyze",
    Agent: analysisAgent,
    Transform: func(input string) string {
        return fmt.Sprintf("Analyze this summary for key insights:\n\n%s", input)
    },
})
```

### Custom Workflow Types
```go
// Try parallel workflow
workflow, err := v1beta.NewParallelWorkflow(&v1beta.WorkflowConfig{
    Mode:    v1beta.Parallel,
    Timeout: 120 * time.Second,
})

// Try DAG workflow  
workflow, err := v1beta.NewDAGWorkflow(&v1beta.WorkflowConfig{
    Mode:    v1beta.DAG,
    Timeout: 180 * time.Second,
})
```

## 🎭 Use Cases

This workflow pattern is ideal for:

- **Research & Analysis**: Multi-step research processes
- **Content Creation**: Planning → Drafting → Editing → Finalizing
- **Data Processing**: Ingestion → Analysis → Reporting → Summary
- **Decision Making**: Information gathering → Analysis → Recommendations
- **Creative Workflows**: Brainstorming → Concept development → Refinement

## 🚀 Next Steps

1. **Try different topics**: Modify the research topics in `main()`
2. **Add more steps**: Extend the workflow with additional phases
3. **Custom workflows**: Try parallel or DAG workflow modes
4. **Error handling**: Add retry logic and graceful degradation
5. **Different LLMs**: Test with OpenAI, Azure, or other providers