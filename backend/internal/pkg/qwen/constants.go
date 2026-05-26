// Package qwen provides helpers for Qwen (通义千问) API integration.
package qwen

import "github.com/Wei-Shaw/sub2api/internal/pkg/openai"

// DefaultModels Qwen models list
var DefaultModels = []openai.Model{
	{ID: "qwen-turbo", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "Qwen Turbo"},
	{ID: "qwen-plus", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "Qwen Plus"},
	{ID: "qwen-max", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "Qwen Max"},
	{ID: "qwen-max-longcontext", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "Qwen Max LongContext"},
	{ID: "qwen-long", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "Qwen Long"},
	{ID: "qwen2.5-72b-instruct", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "Qwen2.5 72B Instruct"},
	{ID: "qwen2.5-32b-instruct", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "Qwen2.5 32B Instruct"},
	{ID: "qwen3-235b-a22b", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "Qwen3 235B A22B"},
	{ID: "qwq-32b", Object: "model", Created: 1700000000, OwnedBy: "qwen", Type: "model", DisplayName: "QWQ 32B"},
}

// DefaultTestModel default model for testing Qwen accounts
const DefaultTestModel = "qwen-plus"
