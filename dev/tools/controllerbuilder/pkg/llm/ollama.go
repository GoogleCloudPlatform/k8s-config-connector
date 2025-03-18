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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"k8s.io/klog/v2"
)

// BuildOllamaClient builds a client for the Ollama API.
func BuildOllamaClient(ctx context.Context) (*OllamaClient, error) {
	host := os.Getenv("OLLAMA_HOST")
	if host == "" {
		host = "http://127.0.0.1:11434/"
	}

	baseURL, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("parsing host %q: %w", host, err)
	}
	klog.Infof("using ollama with base url %v", baseURL.String())

	model := os.Getenv("OLLAMA_MODEL")
	if model == "" {
		klog.Fatalf("OLLAMA_MODEL not set")
	}

	return &OllamaClient{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
		model:      model,
	}, nil
}

type OllamaClient struct {
	baseURL    *url.URL
	httpClient *http.Client
	model      string
}

func (c *OllamaClient) Close() error {
	return nil
}

func (c *OllamaClient) StartChat(systemPrompt string) Chat {
	session := &chatRequest{
		Model: c.model,
	}

	// HACK: Setting the system prompt seems to really mess up some ollama models
	// session.Messages = append(session.Messages, chatMessage{
	// 	Role:    "system",
	// 	Content: systemPrompt,
	// })

	return &OllamaChat{
		session: session,
		client:  c,
	}
}

type OllamaChat struct {
	session *chatRequest
	client  *OllamaClient
}

type chatRequest struct {
	// model: (required) the model name
	Model string `json:"model,omitempty"`
	// messages: the messages of the chat, this can be used to keep a chat memory
	Messages []chatMessage `json:"messages,omitempty"`

	// tools: tools for the model to use if supported. Requires stream to be set to false
	Tools []chatTool `json:"tools,omitempty"`

	// format: the format to return a response in. Format can be json or a JSON schema.

	// options: additional model parameters listed in the documentation for the Modelfile such as temperature

	// stream: if false the response will be returned as a single response object, rather than a stream of objects
	Stream *bool `json:"stream,omitempty"`

	// keep_alive: controls how long the model will stay loaded into memory following the request (default: 5m)
}

type chatResponse struct {
	Model              string       `json:"model"`
	CreatedAt          string       `json:"created_at"`
	Message            *chatMessage `json:"message"`
	Done               bool         `json:"done"`
	TotalDuration      int64        `json:"total_duration"`
	LoadDuration       int64        `json:"load_duration"`
	PromptEvalCount    int64        `json:"prompt_eval_count"`
	PromptEvalDuration int64        `json:"prompt_eval_duration"`
	EvalCount          int64        `json:"eval_count"`
	EvalDuration       int64        `json:"eval_duration"`
}

type completionRequest struct {
	// model: (required) the model name
	Model string `json:"model,omitempty"`
	// prompt: the prompt to generate a response for
	Prompt string `json:"prompt,omitempty"`

	// suffix: the text after the model response

	// images: (optional) a list of base64-encoded images (for multimodal models such as llava)

	// format: the format to return a response in. Format can be json or a JSON schema

	// options: additional model parameters listed in the documentation for the Modelfile such as temperature
	Options map[string]any `json:"options,omitempty"`

	// system: system message to (overrides what is defined in the Modelfile)

	// template: the prompt template to use (overrides what is defined in the Modelfile)

	// stream: if false the response will be returned as a single response object, rather than a stream of objects
	Stream *bool `json:"stream,omitempty"`

	// raw: if true no formatting will be applied to the prompt. You may choose to use the raw parameter if you are specifying a full templated prompt in your request to the API

	// keep_alive: controls how long the model will stay loaded into memory following the request (default: 5m)

	// context (deprecated): the context parameter returned from a previous request to /generate, this can be used to keep a short conversational memory
}

type completionResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`

	//  "context": [1, 2, 3],

	TotalDuration      int64 `json:"total_duration"`
	LoadDuration       int64 `json:"load_duration"`
	PromptEvalCount    int64 `json:"prompt_eval_count"`
	PromptEvalDuration int64 `json:"prompt_eval_duration"`
	EvalCount          int64 `json:"eval_count"`
	EvalDuration       int64 `json:"eval_duration"`
}

type chatMessage struct {
	// role: the role of the message, either system, user, assistant, or tool
	Role string `json:"role,omitempty"`
	// content: the content of the message
	Content string `json:"content,omitempty"`

	// images (optional): a list of images to include in the message (for multimodal models such as llava)
	// Images []chatImage `json:"images,omitempty"`

	// tool_calls (optional): a list of tools the model wants to use
	ToolCalls []chatToolCall `json:"tool_calls,omitempty"`
}

type chatTool struct {
	Type string `json:"type,omitempty"`

	Function *FunctionDefinition `json:"function,omitempty"`
}

type chatToolCall struct {
	Function *FunctionCall `json:"function,omitempty"`
}

type FunctionCall struct {
	Name      string         `json:"name,omitempty"`
	Arguments map[string]any `json:"arguments,omitempty"`
}

type FunctionDefinition struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Parameters  *Schema `json:"parameters,omitempty"`
}

type Schema struct {
	Type        string             `json:"type,omitempty"`
	Properties  map[string]*Schema `json:"properties,omitempty"`
	Description string             `json:"description,omitempty"`
	Required    []string           `json:"required,omitempty"`
}

const TypeObject = "object"
const TypeString = "string"
const TypeBoolean = "boolean"

type FunctionCallResult struct {
	Name   string         `json:"name,omitempty"`
	Result map[string]any `json:"result,omitempty"`
}

func (c *OllamaChat) SetFunctionDefinitions(functionDefinitions []*FunctionDefinition) error {
	// Hack: try trimming whitespace from function definitions
	for _, functionDefinition := range functionDefinitions {
		functionDefinition.Description = strings.TrimSpace(functionDefinition.Description)
		if functionDefinition.Parameters != nil {
			for _, v := range functionDefinition.Parameters.Properties {
				v.Description = strings.TrimSpace(v.Description)
			}
		}
	}

	for _, functionDefinition := range functionDefinitions {
		c.session.Tools = append(c.session.Tools, chatTool{
			Type:     "function",
			Function: functionDefinition,
		})
	}
	return nil
}

// func (c *OllamaChat) AdditionalUserInput(s string) {
// 	c.session.Messages = append(c.session.Messages, chatMessage{
// 		Role:    "user",
// 		Content: s,
// 	})
// }

func (c *OllamaChat) SendMessage(ctx context.Context, parts ...string) (Response, error) {
	for _, part := range parts {
		c.session.Messages = append(c.session.Messages, chatMessage{
			Role:    "user",
			Content: part,
		})
		klog.Infof("sending user:\n%v", part)
	}

	ollamaResponse, err := c.client.doChat(ctx, c.session)
	if err != nil {
		return nil, err
	}

	if ollamaResponse.Message == nil {
		return nil, fmt.Errorf("no message returned from ollama")
	}

	c.session.Messages = append(c.session.Messages, *ollamaResponse.Message)

	response := &OllamaResponse{ollamaResponse: ollamaResponse}

	return response, nil
}

func (c *OllamaClient) GenerateCompletion(ctx context.Context, request *CompletionRequest) (CompletionResponse, error) {
	ollamaRequest := &completionRequest{
		Model:  c.model,
		Prompt: request.Prompt,
		Options: map[string]any{
			"num_ctx": 128 * 1024,
		},
	}

	ollamaResponse, err := c.doCompletion(ctx, ollamaRequest)
	if err != nil {
		return nil, err
	}

	if ollamaResponse.Response == "" {
		return nil, fmt.Errorf("no response returned from ollama")
	}

	response := &OllamaCompletionResponse{ollamaResponse: ollamaResponse}
	return response, nil
}

func (c *OllamaClient) doCompletion(ctx context.Context, req *completionRequest) (*completionResponse, error) {
	stream := false
	req.Stream = &stream

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("building json body: %w", err)
	}
	u := c.baseURL.JoinPath("api", "generate")
	klog.V(2).Infof("sending POST request to %v: %v", u.String(), string(body))
	httpRequest, err := http.NewRequestWithContext(ctx, "POST", u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("building http request: %w", err)
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("performing http request: %w", err)
	}
	defer httpResponse.Body.Close()

	b, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	klog.Infof("response is: %v", string(b))

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected http status: %q with response %q", httpResponse.Status, string(b))
	}

	completionResponse := &completionResponse{}
	if err := json.Unmarshal(b, completionResponse); err != nil {
		return nil, fmt.Errorf("unmarshalling json response: %w", err)
	}
	return completionResponse, nil
}

func (c *OllamaClient) doChat(ctx context.Context, req *chatRequest) (*chatResponse, error) {
	stream := false
	req.Stream = &stream

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("building json body: %w", err)
	}
	u := c.baseURL.JoinPath("api", "chat")
	klog.V(2).Infof("sending POST request to %v: %v", u.String(), string(body))
	httpRequest, err := http.NewRequestWithContext(ctx, "POST", u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("building http request: %w", err)
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("performing http request: %w", err)
	}
	defer httpResponse.Body.Close()

	b, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	klog.Infof("response is: %v", string(b))

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected http status: %q with response %q", httpResponse.Status, string(b))
	}

	chatResponse := &chatResponse{}
	if err := json.Unmarshal(b, chatResponse); err != nil {
		return nil, fmt.Errorf("unmarshalling json response: %w", err)
	}
	return chatResponse, nil
}

func (c *OllamaChat) SendFunctionResults(ctx context.Context, functionResults []FunctionCallResult) (Response, error) {
	for _, functionResult := range functionResults {
		b, err := json.Marshal(functionResult)
		if err != nil {
			return nil, fmt.Errorf("building json for function result: %w", err)
		}
		c.session.Messages = append(c.session.Messages, chatMessage{
			Role:    "tool",
			Content: string(b),
		})
	}
	ollamaResponse, err := c.client.doChat(ctx, c.session)
	if err != nil {
		return nil, err
	}
	if ollamaResponse.Message == nil {
		return nil, fmt.Errorf("no message returned from ollama")
	}

	c.session.Messages = append(c.session.Messages, *ollamaResponse.Message)

	response := &OllamaResponse{ollamaResponse: ollamaResponse}

	return response, nil
}

type OllamaResponse struct {
	ollamaResponse *chatResponse
}

func (r *OllamaResponse) UsageMetadata() any {
	return r.ollamaResponse
}

func (r *OllamaResponse) Candidates() []Candidate {
	var candidates []Candidate
	candidates = append(candidates, &OllamaCandidate{candidate: r.ollamaResponse.Message})
	return candidates
}

type OllamaCandidate struct {
	candidate *chatMessage
}

func (c *OllamaCandidate) Parts() []Part {
	var parts []Part
	parts = append(parts, &OllamaPart{part: c.candidate})
	return parts
}

type OllamaPart struct {
	part *chatMessage
}

func (p *OllamaPart) AsText() (string, bool) {
	if p.part.Role == "assistant" && p.part.Content != "" {
		return p.part.Content, true
	}
	return "", false
}

func (p *OllamaPart) AsFunctionCalls() ([]FunctionCall, bool) {
	var functionCalls []FunctionCall
	for _, toolCall := range p.part.ToolCalls {
		if toolCall.Function != nil {
			functionCalls = append(functionCalls, *toolCall.Function)
		}
	}
	return functionCalls, true
}

type OllamaCompletionResponse struct {
	ollamaResponse *completionResponse
}

var _ CompletionResponse = &OllamaCompletionResponse{}

func (r *OllamaCompletionResponse) Response() string {
	return r.ollamaResponse.Response
}

func (r *OllamaCompletionResponse) UsageMetadata() any {
	return r.ollamaResponse
}
