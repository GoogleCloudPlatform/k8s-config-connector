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
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

func init() {
	// RegisterTool(&CreateSubtask{})
}

type CreateSubtask struct {
	Description string `json:"description"`
}

type CreateSubtaskResults struct {
	Success bool `json:"success"`
}

func (t *CreateSubtask) Run(ctx context.Context, c *Chat, args map[string]any) (*CreateSubtaskResults, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	klog.V(2).Infof("%T: %+v", t, t)

	results := &CreateSubtaskResults{
		Success: true,
	}
	return results, nil
}

func (t *CreateSubtask) BuildFunctionDefinition() *llm.FunctionDefinition {
	declaration := &llm.FunctionDefinition{
		Name:        "CreateSubtask",
		Description: `Create sub-tasks to break down a large amount of work into smaller chunks that.`,
		Parameters: &llm.Schema{
			Type:     llm.TypeObject,
			Required: []string{"description"},
			Properties: map[string]*llm.Schema{
				"description": {
					Type:        llm.TypeString,
					Description: `A description of the subtask.`,
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}
