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
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
)

type FunctionResult struct {
	Response any
	Error    error
}

type Toolbox struct {
	tools map[string]*toolInfo
}

func NewToolbox(tools []Tool) *Toolbox {
	m := make(map[string]*toolInfo)
	for _, tool := range tools {
		functionDefintion := tool.BuildFunctionDefinition()
		toolType := reflect.TypeOf(tool).Elem()
		m[functionDefintion.Name] = &toolInfo{
			functionDefinition: functionDefintion,
			tool:               tool,
			toolType:           toolType,
		}

	}
	return &Toolbox{
		tools: m,
	}
}

func (t *Toolbox) GetFunctionDefinitions() []*gollm.FunctionDefinition {
	var functionDefinitions []*gollm.FunctionDefinition
	for _, toolInfo := range t.tools {
		functionDefinitions = append(functionDefinitions, toolInfo.functionDefinition)
	}
	return functionDefinitions
}

type toolInfo struct {
	functionDefinition *gollm.FunctionDefinition
	tool               Tool
	toolType           reflect.Type
}

func (t *toolInfo) newInstance() Tool {
	obj := reflect.New(t.toolType).Interface()
	return obj.(Tool)
}

func (t *Toolbox) CallFunction(ctx context.Context, c *Chat, functionCall gollm.FunctionCall) (*FunctionResult, error) {
	toolInfo := t.tools[functionCall.Name]
	if toolInfo == nil {
		// TODO: Fatal or return an error?
		return nil, fmt.Errorf("unknown function %q", functionCall.Name)
	}

	tool := toolInfo.newInstance()
	result, err := tool.Run(ctx, c, functionCall.Arguments)
	return &FunctionResult{
		Response: result,
		Error:    err,
	}, nil

}
