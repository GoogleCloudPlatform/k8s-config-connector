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
	"os"
	"os/exec"
	"strings"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/option"
	"k8s.io/klog/v2"
)

// BuildGeminiClient builds a client for the Gemini API.
func BuildGeminiClient(ctx context.Context) (*genai.Client, error) {
	log := klog.FromContext(ctx)

	var opts []option.ClientOption

	if s := os.Getenv("GEMINI_API_KEY"); s != "" {
		opts = append(opts, option.WithAPIKey(s))
	}
	// else {
	// 	creds, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/generative-language", "https://www.googleapis.com/auth/cloud-platform")
	// 	if err != nil {
	// 		return nil, fmt.Errorf("finding default credentials: %w", err)
	// 	}
	// 	opts = append(opts, option.WithCredentials(creds))
	// }

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
		return nil, fmt.Errorf("building gemini client: %w", err)
	}
	return client, nil
}
