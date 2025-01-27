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

	"k8s.io/klog/v2"
)

// BuildLlamaCPPClient builds a client for the llama-cpp server API.
func BuildLlamaCPPClient(ctx context.Context) (*LlamaCPPClient, error) {
	host := os.Getenv("LLAMACPP_HOST")
	if host == "" {
		host = "http://127.0.0.1:11434/"
	}

	baseURL, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("parsing host %q: %w", host, err)
	}
	klog.Infof("using llama.cpp with base url %v", baseURL.String())

	return &LlamaCPPClient{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}, nil
}

type LlamaCPPClient struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func (c *LlamaCPPClient) Close() error {
	return nil
}

func (c *LlamaCPPClient) StartChat(systemPrompt string) Chat {
	panic("chat not implemented")
}

type llamacppCompletionRequest struct {
	// See https://github.com/ggerganov/llama.cpp/blob/master/examples/server/README.md#post-completion-given-a-prompt-it-returns-the-predicted-completion

	Prompt string `json:"prompt,omitempty"`
}

type llamacppCompletionResponse struct {
	Content string `json:"content,omitempty"`
}

func (c *LlamaCPPClient) GenerateCompletion(ctx context.Context, request *CompletionRequest) (CompletionResponse, error) {
	llamacppRequest := &llamacppCompletionRequest{
		Prompt: request.Prompt,
	}

	llamacppResponse, err := c.doCompletion(ctx, llamacppRequest)
	if err != nil {
		return nil, err
	}

	if llamacppResponse.Content == "" {
		return nil, fmt.Errorf("no response returned from llamacpp")
	}

	response := &LlamaCPPCompletionResponse{llamacppResponse: llamacppResponse}
	return response, nil
}

func (c *LlamaCPPClient) doCompletion(ctx context.Context, req *llamacppCompletionRequest) (*llamacppCompletionResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("building json body: %w", err)
	}
	u := c.baseURL.JoinPath("completion")
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

	completionResponse := &llamacppCompletionResponse{}
	if err := json.Unmarshal(b, completionResponse); err != nil {
		return nil, fmt.Errorf("unmarshalling json response: %w", err)
	}
	return completionResponse, nil
}

type LlamaCPPCompletionResponse struct {
	llamacppResponse *llamacppCompletionResponse
}

var _ CompletionResponse = &LlamaCPPCompletionResponse{}

func (r *LlamaCPPCompletionResponse) Response() string {
	return r.llamacppResponse.Content
}

func (r *LlamaCPPCompletionResponse) UsageMetadata() any {
	return r.llamacppResponse
}
