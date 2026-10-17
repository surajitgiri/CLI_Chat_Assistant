package provider

import (
	"chatcli/internal/chat"
	"context"
	"fmt"
	"os"
)

// Provider streams a reply, calling onToken for each chunk.
// It returns the full reply when done.
type Provider interface {
	Name() string
	Stream(ctx context.Context, msgs []chat.Message, onToken func(string)) (string error)
}

func New(name, model string) (Provider, error) {
	switch name {
	case "ollama":
		if model == "" {
			model = "llama3.2"
		}
		return &Ollama{Model: model, URL: "http://localhost:11434/api/chat"}, nil
	case "openai":
		key := os.Getenv("OPENAI_API_KEY")
		if key == "" {
			return nil, fmt.Errorf("OPENAI_API_KEY not set")
		}
		if model == "" {
			model = "gpt-4o-mini"
		}
		return &OpenAI{Model: model, APIKey: key}, nil
	}
	return nil, fmt.Errorf("unknown provider %q", name)
}
