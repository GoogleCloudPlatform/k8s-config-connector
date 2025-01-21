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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
)

type FunctionResult struct {
	Response any
	Error    error
}

func (c *Chat) runFunctionCall(ctx context.Context, functionCall llm.FunctionCall) (*FunctionResult, error) {
	switch functionCall.Name {
	case "EditFile":
		t := &EditFile{}
		result, err := t.Run(ctx, c, functionCall.Arguments)
		return &FunctionResult{
			Response: result,
			Error:    err,
		}, nil
	case "VerifyCode":
		t := &VerifyCode{}
		result, err := t.Run(ctx, c, functionCall.Arguments)
		return &FunctionResult{
			Response: result,
			Error:    err,
		}, nil
	case "ReadFile":
		t := &ReadFile{}
		result, err := t.Run(ctx, c, functionCall.Arguments)
		return &FunctionResult{
			Response: result,
			Error:    err,
		}, nil
	case "CreateFile":
		t := &CreateFile{}
		result, err := t.Run(ctx, c, functionCall.Arguments)
		return &FunctionResult{
			Response: result,
			Error:    err,
		}, nil
	case "ASTEdit":
		t := &ASTEdit{}
		result, err := t.Run(ctx, c, functionCall.Arguments)
		return &FunctionResult{
			Response: result,
			Error:    err,
		}, nil
	case "FindInWorkspace":
		t := &FindInWorkspace{}
		result, err := t.Run(ctx, c, functionCall.Arguments)
		return &FunctionResult{
			Response: result,
			Error:    err,
		}, nil
	case "ListFilesInWorkspace":
		t := &FindInWorkspace{}
		result, err := t.Run(ctx, c, functionCall.Arguments)
		return &FunctionResult{
			Response: result,
			Error:    err,
		}, nil
	case "RunShellCommand":
		t := &RunShellCommand{}
		result, err := t.Run(ctx, c, functionCall.Arguments)
		return &FunctionResult{
			Response: result,
			Error:    err,
		}, nil
	default:
		// TODO: Fatal or return an error?
		return nil, fmt.Errorf("unknown function %q", functionCall.Name)
	}
}
