package provider

import (
	"chatcli/internal/chat"
	"context"
	"os"
	"fmt"
)

// Provider streams a reply, calling onToken for each chunk.
// It returns the full reply when done.
type Provider interface {
	Name() string
	Stream(ctx context.Context, msgs []chat.Message, onToken func(string)) (string error)
}

func 