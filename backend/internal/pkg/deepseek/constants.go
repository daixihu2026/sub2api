// Package deepseek provides helpers for DeepSeek API integration.
package deepseek

import "github.com/Wei-Shaw/sub2api/internal/pkg/openai"

// DefaultModels DeepSeek models list
var DefaultModels = []openai.Model{
	{ID: "deepseek-chat", Object: "model", Created: 1700000000, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek Chat"},
	{ID: "deepseek-coder", Object: "model", Created: 1700000000, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek Coder"},
	{ID: "deepseek-reasoner", Object: "model", Created: 1700000000, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek Reasoner"},
	{ID: "deepseek-v3", Object: "model", Created: 1700000000, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek V3"},
	{ID: "deepseek-v3-0324", Object: "model", Created: 1700000000, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek V3 (0324)"},
	{ID: "deepseek-r1", Object: "model", Created: 1700000000, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek R1"},
	{ID: "deepseek-r1-0528", Object: "model", Created: 1700000000, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek R1 (0528)"},
}

// DefaultTestModel default model for testing DeepSeek accounts
const DefaultTestModel = "deepseek-chat"
