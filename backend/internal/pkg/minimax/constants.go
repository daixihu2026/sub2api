// Package minimax provides helpers for MiniMax API integration.
package minimax

import "github.com/Wei-Shaw/sub2api/internal/pkg/openai"

// DefaultModels MiniMax models list
var DefaultModels = []openai.Model{
	{ID: "abab6.5-chat", Object: "model", Created: 1700000000, OwnedBy: "minimax", Type: "model", DisplayName: "Abab6.5 Chat"},
	{ID: "abab6.5s-chat", Object: "model", Created: 1700000000, OwnedBy: "minimax", Type: "model", DisplayName: "Abab6.5s Chat"},
	{ID: "abab6.5s-chat-pro", Object: "model", Created: 1700000000, OwnedBy: "minimax", Type: "model", DisplayName: "Abab6.5s Chat Pro"},
	{ID: "abab6-chat", Object: "model", Created: 1700000000, OwnedBy: "minimax", Type: "model", DisplayName: "Abab6 Chat"},
}

// DefaultTestModel default model for testing MiniMax accounts
const DefaultTestModel = "abab6.5s-chat"
