package llm

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func NewLLMClientFromEnvVar(ctx context.Context) (Client, error) {
	cfg := os.Getenv("LLM_CLIENT")
	if cfg == "" {
		return nil, fmt.Errorf("LLM_CLIENT env var not set")
	}

	// vertexai => vertexai://
	if !strings.Contains(cfg, "/") {
		cfg += "://"
	}

	u, err := url.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("parsing LLM_CLIENT env var: %w", err)
	}

	switch u.Scheme {
	case "gemini":
		return BuildGeminiClient(ctx)
	case "ollama":
		return BuildOllamaClient(ctx)
	case "vertexai":
		return BuildVertexAIClient(ctx)
	default:
		return nil, fmt.Errorf("LLM_CLIENT not recognized (use vertexai or gemini)")
	}
}
