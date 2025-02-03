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
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"cloud.google.com/go/vertexai/genai"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"k8s.io/klog/v2"
)

// BuildVertexAIClient builds a client for the VertexAI API.
func BuildVertexAIClient(ctx context.Context) (*VertexAIClient, error) {
	log := klog.FromContext(ctx)

	var opts []option.ClientOption

	creds, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/generative-language", "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, fmt.Errorf("finding default credentials: %w", err)
	}
	opts = append(opts, option.WithCredentials(creds))

	projectID := ""
	location := ""

	if projectID == "" {
		cmd := exec.CommandContext(ctx, "gcloud", "config", "get", "project")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("cannot get project (using gcloud config get project): %w", err)
		}
		projectID = strings.TrimSpace(stdout.String())
		if projectID == "" {
			return nil, fmt.Errorf("project was not set in gcloud config")
		}
		log.Info("got project from gcloud config", "project", projectID)
	}

	client, err := genai.NewClient(ctx, projectID, location, opts...)
	if err != nil {
		return nil, fmt.Errorf("building vertexai client: %w", err)
	}
	return &VertexAIClient{client: client}, nil
}

type VertexAIClient struct {
	client *genai.Client
}

func (c *VertexAIClient) Close() error {
	return c.client.Close()
}

func (c *VertexAIClient) StartChat(systemPrompt string) Chat {
	// model := c.client.GenerativeModel("vertexai-1.5-flash")
	// model := c.client.GenerativeModel("vertexai-exp-1206")
	model := c.client.GenerativeModel("gemini-2.0-flash-exp")
	// model := c.client.GenerativeModel("gemma-2-27b-it")
	// model := c.client.GenerativeModel("gemini-1.5-pro-002")

	// Some values that are recommended by aistudio
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"

	if systemPrompt != "" {
		model.SystemInstruction = &genai.Content{
			Parts: []genai.Part{
				genai.Text(systemPrompt),
			},
		}
	} else {
		klog.Warningf("systemPrompt not provided")
	}

	// if model.ToolConfig == nil {
	// 	model.ToolConfig = &genai.ToolConfig{}
	// }
	// if model.ToolConfig.FunctionCallingConfig == nil {
	// 	model.ToolConfig.FunctionCallingConfig = &genai.FunctionCallingConfig{}
	// }
	// model.ToolConfig.FunctionCallingConfig.Mode = genai.FunctionCallingAny

	chat := model.StartChat()

	return &VertexAIChat{
		model: model,
		chat:  chat,
	}
}

type VertexAIChat struct {
	model *genai.GenerativeModel
	chat  *genai.ChatSession
}

func (c *VertexAIChat) SetFunctionDefinitions(functionDefinitions []*FunctionDefinition) error {
	var vertexaiFunctionDefinitions []*genai.FunctionDeclaration
	for _, functionDefinition := range functionDefinitions {
		parameters, err := toVertexAISchema(functionDefinition.Parameters)
		if err != nil {
			return err
		}
		vertexaiFunctionDefinitions = append(vertexaiFunctionDefinitions, &genai.FunctionDeclaration{
			Name:        functionDefinition.Name,
			Description: functionDefinition.Description,
			Parameters:  parameters,
		})
	}

	c.model.Tools = append(c.model.Tools, &genai.Tool{
		FunctionDeclarations: vertexaiFunctionDefinitions,
	})
	return nil
}

// Converts our generic Schema to a genai.Schema
func toVertexAISchema(schema *Schema) (*genai.Schema, error) {
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
			vertexaiValue, err := toVertexAISchema(v)
			if err != nil {
				return nil, err
			}
			ret.Properties[k] = vertexaiValue
		}
	}
	return ret, nil
}

// func (c *VertexAIChat) AdditionalUserInput(s string) {
// 	// c.model.SystemInstruction.Parts = append(c.model.SystemInstruction.Parts, genai.Text(s))
// 	c.chat.History = append(c.chat.History, &genai.Content{
// 		Role:  "user",
// 		Parts: []genai.Part{genai.Text(s)},
// 	})
// }

func (c *VertexAIChat) SendMessage(ctx context.Context, parts ...string) (Response, error) {
	log := klog.FromContext(ctx)
	var vertexaiParts []genai.Part
	for _, part := range parts {
		vertexaiParts = append(vertexaiParts, genai.Text(part))
	}
	log.Info("sending LLM request", "user", parts)
	vertexaiResponse, err := c.chat.SendMessage(ctx, vertexaiParts...)
	if err != nil {
		return nil, err
	}
	return &VertexAIResponse{vertexaiResponse: vertexaiResponse}, nil
}

func (c *VertexAIChat) SendFunctionResults(ctx context.Context, functionResults []FunctionCallResult) (Response, error) {
	var vertexaiFunctionResults []genai.Part
	for _, functionResult := range functionResults {
		vertexaiFunctionResults = append(vertexaiFunctionResults, genai.FunctionResponse{
			Name:     functionResult.Name,
			Response: functionResult.Result,
		})
	}

	vertexaiResponse, err := c.chat.SendMessage(ctx, vertexaiFunctionResults...)
	if err != nil {
		return nil, err
	}
	return &VertexAIResponse{vertexaiResponse: vertexaiResponse}, nil
}

type VertexAIResponse struct {
	vertexaiResponse *genai.GenerateContentResponse
}

func (r *VertexAIResponse) UsageMetadata() any {
	return r.vertexaiResponse.UsageMetadata
}

func (r *VertexAIResponse) Candidates() []Candidate {
	var candidates []Candidate
	for _, candidate := range r.vertexaiResponse.Candidates {
		// klog.Infof("candidate: %+v", candidate)
		candidates = append(candidates, &VertexAICandidate{candidate: candidate})
	}
	return candidates
}

type VertexAICandidate struct {
	candidate *genai.Candidate
}

func (r *VertexAICandidate) Parts() []Part {
	var parts []Part
	if r.candidate.Content != nil {
		for _, part := range r.candidate.Content.Parts {
			parts = append(parts, &VertexAIPart{part: part})
		}
	}
	return parts
}

type VertexAIPart struct {
	part genai.Part
}

func (p *VertexAIPart) AsText() (string, bool) {
	if text, ok := p.part.(genai.Text); ok {
		return string(text), true
	}
	return "", false
}

func (p *VertexAIPart) AsFunctionCalls() ([]FunctionCall, bool) {
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
