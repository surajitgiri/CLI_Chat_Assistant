package provider

import (
	"CLI_Chat_Assistant/internal/chat"
	"context"
)

// Provider streams a reply, calling onToken for each chunk.
// It returns the full reply when done.
type Provider interface {
	Name() string
	Stream(ctx context.Context, msg []chat.Message)
}
