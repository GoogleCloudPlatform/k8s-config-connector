// Copyright 2024 Google LLC
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
	"fmt"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"k8s.io/klog/v2"
)

// BuildGeminiClient builds a client for the Gemini API.
func BuildGeminiClient(ctx context.Context) (Client, error) {
	var opts []option.ClientOption

	if s := os.Getenv("GEMINI_API_KEY"); s != "" {
		opts = append(opts, option.WithAPIKey(s))
	}

	client, err := genai.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building gemini client: %w", err)
	}

	return &GeminiClient{
		client: client,
		model:  "gemini-2.0-pro-exp-02-05",
	}, nil
}

type GeminiClient struct {
	client *genai.Client
	model  string
}

func (c *GeminiClient) Close() error {
	return c.client.Close()
}

func (c *GeminiClient) GenerateCompletion(ctx context.Context, request *CompletionRequest) (CompletionResponse, error) {
	return nil, fmt.Errorf("GeminiClient::GenerateCompletion not implemented")
}

func (c *GeminiClient) StartChat(systemPrompt string) Chat {
	model := c.client.GenerativeModel(c.model)
	// model := c.client.GenerativeModel("gemini-1.5-pro-002")

	// Some values that are recommended by aistudio
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"

	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	// if model.ToolConfig == nil {
	// 	model.ToolConfig = &genai.ToolConfig{}
	// }
	// if model.ToolConfig.FunctionCallingConfig == nil {
	// 	model.ToolConfig.FunctionCallingConfig = &genai.FunctionCallingConfig{}
	// }
	// model.ToolConfig.FunctionCallingConfig.Mode = genai.FunctionCallingAny

	chat := model.StartChat()

	return &GeminiChat{
		model: model,
		chat:  chat,
	}
}

type GeminiChat struct {
	model *genai.GenerativeModel
	chat  *genai.ChatSession
}

func (c *GeminiChat) SetFunctionDefinitions(functionDefinitions []*FunctionDefinition) error {
	var geminiFunctionDefinitions []*genai.FunctionDeclaration
	for _, functionDefinition := range functionDefinitions {
		parameters, err := toGenaiSchema(functionDefinition.Parameters)
		if err != nil {
			return err
		}
		geminiFunctionDefinitions = append(geminiFunctionDefinitions, &genai.FunctionDeclaration{
			Name:        functionDefinition.Name,
			Description: functionDefinition.Description,
			Parameters:  parameters,
		})
	}

	c.model.Tools = append(c.model.Tools, &genai.Tool{
		FunctionDeclarations: geminiFunctionDefinitions,
	})
	return nil
}

// Converts our generic Schema to a genai.Schema
func toGenaiSchema(schema *Schema) (*genai.Schema, error) {
	ret := &genai.Schema{
		Description: schema.Description,
		Required:    schema.Required,
	}

	switch schema.Type {
	case TypeObject:
		ret.Type = genai.TypeObject
	case TypeString:
		ret.Type = genai.TypeString
	default:
		return nil, fmt.Errorf("type %q not handled by genai.Schema", schema.Type)
	}
	if schema.Properties != nil {
		ret.Properties = make(map[string]*genai.Schema)
		for k, v := range schema.Properties {
			geminiV, err := toGenaiSchema(v)
			if err != nil {
				return nil, err
			}
			ret.Properties[k] = geminiV
		}
	}
	return ret, nil
}

//	func (c *GeminiChat) AdditionalUserInput(s string) {
//		// c.model.SystemInstruction.Parts = append(c.model.SystemInstruction.Parts, genai.Text(s))
//		c.chat.History = append(c.chat.History, &genai.Content{
//			Role:  "user",
//			Parts: []genai.Part{genai.Text(s)},
//		})
//	}
func (c *GeminiChat) sendMessageWithRetries(ctx context.Context, geminiParts ...genai.Part) (*genai.GenerateContentResponse, error) {
	opt := gax.WithRetry(func() gax.Retryer {
		// https://cloud.google.com/vertex-ai/generative-ai/docs/error-code-429
		return gax.OnCodes([]codes.Code{codes.ResourceExhausted}, gax.Backoff{
			Initial:    5 * time.Second,
			Max:        10 * time.Minute,
			Multiplier: 2,
		})
	})
	var resp *genai.GenerateContentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.chat.SendMessage(ctx, geminiParts...)
		return err
	}, opt)
	return resp, err
}

func (c *GeminiChat) SendMessage(ctx context.Context, parts ...string) (Response, error) {
	log := klog.FromContext(ctx)
	var geminiParts []genai.Part
	for _, part := range parts {
		geminiParts = append(geminiParts, genai.Text(part))
	}
	log.Info("sending LLM request", "user", parts)

	geminiResponse, err := c.sendMessageWithRetries(ctx, geminiParts...)
	if err != nil {
		return nil, err
	}
	return &GeminiResponse{geminiResponse: geminiResponse}, nil
}

func (c *GeminiChat) SendFunctionResults(ctx context.Context, functionResults []FunctionCallResult) (Response, error) {
	var geminiFunctionResults []genai.Part
	for _, functionResult := range functionResults {
		geminiFunctionResults = append(geminiFunctionResults, genai.FunctionResponse{
			Name:     functionResult.Name,
			Response: functionResult.Result,
		})
	}

	geminiResponse, err := c.sendMessageWithRetries(ctx, geminiFunctionResults...)
	if err != nil {
		return nil, err
	}
	return &GeminiResponse{geminiResponse: geminiResponse}, nil
}

type GeminiResponse struct {
	geminiResponse *genai.GenerateContentResponse
}

func (r *GeminiResponse) UsageMetadata() any {
	return r.geminiResponse.UsageMetadata
}

func (r *GeminiResponse) Candidates() []Candidate {
	var candidates []Candidate
	for _, candidate := range r.geminiResponse.Candidates {
		klog.Infof("candidate: %+v", candidate)
		candidates = append(candidates, &GeminiCandidate{candidate: candidate})
	}
	return candidates
}

type GeminiCandidate struct {
	candidate *genai.Candidate
}

func (r *GeminiCandidate) Parts() []Part {
	var parts []Part
	if r.candidate.Content != nil {
		for _, part := range r.candidate.Content.Parts {
			parts = append(parts, &GeminiPart{part: part})
		}
	}
	return parts
}

type GeminiPart struct {
	part genai.Part
}

func (p *GeminiPart) AsText() (string, bool) {
	if text, ok := p.part.(genai.Text); ok {
		return string(text), true
	}
	return "", false
}

func (p *GeminiPart) AsFunctionCalls() ([]FunctionCall, bool) {
	if functionCall, ok := p.part.(genai.FunctionCall); ok {
		var ret []FunctionCall
		ret = append(ret, FunctionCall{
			Name:      functionCall.Name,
			Arguments: functionCall.Args,
		})
		return ret, true
	}
	return nil, false
}
