// Package zhipu provides helpers for Zhipu AI (智谱) API integration.
package zhipu

import "github.com/Wei-Shaw/sub2api/internal/pkg/openai"

// DefaultModels Zhipu models list
var DefaultModels = []openai.Model{
	{ID: "glm-4", Object: "model", Created: 1700000000, OwnedBy: "zhipu", Type: "model", DisplayName: "GLM-4"},
	{ID: "glm-4v", Object: "model", Created: 1700000000, OwnedBy: "zhipu", Type: "model", DisplayName: "GLM-4V"},
	{ID: "glm-4-plus", Object: "model", Created: 1700000000, OwnedBy: "zhipu", Type: "model", DisplayName: "GLM-4 Plus"},
	{ID: "glm-4-air", Object: "model", Created: 1700000000, OwnedBy: "zhipu", Type: "model", DisplayName: "GLM-4 Air"},
	{ID: "glm-4-flash", Object: "model", Created: 1700000000, OwnedBy: "zhipu", Type: "model", DisplayName: "GLM-4 Flash"},
	{ID: "glm-4.5", Object: "model", Created: 1700000000, OwnedBy: "zhipu", Type: "model", DisplayName: "GLM-4.5"},
	{ID: "glm-4.6", Object: "model", Created: 1700000000, OwnedBy: "zhipu", Type: "model", DisplayName: "GLM-4.6"},
}

// DefaultTestModel default model for testing Zhipu accounts
const DefaultTestModel = "glm-4-flash"
