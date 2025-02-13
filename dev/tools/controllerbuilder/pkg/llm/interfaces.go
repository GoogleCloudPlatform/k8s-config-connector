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

package llm

import (
	"context"
	"io"
)

type Client interface {
	io.Closer
	StartChat(systemPrompt string) Chat

	GenerateCompletion(ctx context.Context, req *CompletionRequest) (CompletionResponse, error)
}

type Chat interface {
	SendMessage(ctx context.Context, userParts ...string) (Response, error)
	SetFunctionDefinitions(functionDefinitions []*FunctionDefinition) error
	SendFunctionResults(ctx context.Context, functionResults []FunctionCallResult) (Response, error)
	// AdditionalUserInput(s string)
}

type Response interface {
	UsageMetadata() any
	Candidates() []Candidate
}

type Candidate interface {
	Parts() []Part
}

type Part interface {
	AsText() (string, bool)
	AsFunctionCalls() ([]FunctionCall, bool)
}

type CompletionRequest struct {
	Prompt string
}

type CompletionResponse interface {
	Response() string
	UsageMetadata() any
}
