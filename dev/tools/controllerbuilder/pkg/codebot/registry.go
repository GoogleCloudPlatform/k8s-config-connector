// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codebot

import (
	"context"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
)

func RegisterTool(tool Tool) {
	tools = append(tools, tool)
}

var tools []Tool

type Tool interface {
	BuildFunctionDefinition() *gollm.FunctionDefinition

	Run(ctx context.Context, c *Chat, args map[string]any) (any, error)
}

func GetAllTools() []Tool {
	return tools
}

// GetFilteredTools returns a list of tools that match the filter criteria.
// It excludes tools that are marked for exclusion and includes those that are marked for inclusion.
func GetFilteredTools(filter Filter) []Tool {
	filtered := make([]Tool, 0, len(tools))
	for _, tool := range tools {
		if filter.Exclude(tool) {
			continue
		}
		if filter.Include(tool) {
			filtered = append(filtered, tool)
		}
	}
	return filtered
}
