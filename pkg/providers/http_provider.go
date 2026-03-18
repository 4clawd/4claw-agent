// 4claw - Ultra-lightweight personal AI agent
// Built and maintained by 4claw contributors.
// License: MIT
//
// Copyright (c) 2026 4claw contributors

package providers

import (
	"context"
	"time"

	"github.com/4claw/4claw/pkg/providers/openai_compat"
)

type HTTPProvider struct {
	delegate *openai_compat.Provider
}

func NewHTTPProvider(apiKey, apiBase, proxy string) *HTTPProvider {
	return &HTTPProvider{
		delegate: openai_compat.NewProvider(apiKey, apiBase, proxy),
	}
}

func NewHTTPProviderWithTimeout(apiKey, apiBase, proxy string, timeout time.Duration) *HTTPProvider {
	return &HTTPProvider{
		delegate: openai_compat.NewProviderWithTimeout(apiKey, apiBase, proxy, timeout),
	}
}

func NewHTTPProviderWithMaxTokensField(apiKey, apiBase, proxy, maxTokensField string) *HTTPProvider {
	return &HTTPProvider{
		delegate: openai_compat.NewProviderWithMaxTokensField(apiKey, apiBase, proxy, maxTokensField),
	}
}

func NewHTTPProviderWithMaxTokensFieldAndTimeout(
	apiKey, apiBase, proxy, maxTokensField string,
	timeout time.Duration,
) *HTTPProvider {
	return &HTTPProvider{
		delegate: openai_compat.NewProviderWithMaxTokensFieldAndTimeout(
			apiKey,
			apiBase,
			proxy,
			maxTokensField,
			timeout,
		),
	}
}

func (p *HTTPProvider) Chat(
	ctx context.Context,
	messages []Message,
	tools []ToolDefinition,
	model string,
	options map[string]any,
) (*LLMResponse, error) {
	return p.delegate.Chat(ctx, messages, tools, model, options)
}

func (p *HTTPProvider) GetDefaultModel() string {
	return ""
}
