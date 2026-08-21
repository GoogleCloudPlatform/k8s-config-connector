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

package gcpwatch

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"google.golang.org/api/option"
	"k8s.io/klog/v2"
)

type IAMFetcher struct {
	crmClient *resourcemanager.ProjectsClient
}

var _ Fetcher = &IAMFetcher{}

// todo acpana: house in direct?
func NewIAMFetcher(ctx context.Context, config *config.ControllerConfig) (*IAMFetcher, error) {
	opts, err := options(config)
	if err != nil {
		return nil, err
	}
	crmClient, err := resourcemanager.NewProjectsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building projects client: %w", err)
	}
	return &IAMFetcher{
		crmClient: crmClient,
	}, nil
}

func options(config *config.ControllerConfig) ([]option.ClientOption, error) {
	var opts []option.ClientOption
	if config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(config.UserAgent))
	}
	if config.HTTPClient != nil {
		// TODO: Set UserAgent in this scenario (error is: WithHTTPClient is incompatible with gRPC dial options)
		opts = append(opts, option.WithHTTPClient(config.HTTPClient))
	}
	if config.UserProjectOverride && config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(config.BillingProject))
	}

	// TODO: support endpoints?
	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	return opts, nil
}

func (f *IAMFetcher) IsSupported(kind string, external string) bool {
	tokens := strings.Split(external, "/")
	switch kind {
	case "Project":
		return len(tokens) == 2 && tokens[0] == "projects"
	}

	return false
}

func (f *IAMFetcher) Fetch(ctx context.Context, kind string, external string) (*ResourceInfo, error) {
	tokens := strings.Split(external, "/")
	switch kind {
	case "Project":
		if len(tokens) == 2 && tokens[0] == "projects" {
			return f.fetchProject(ctx, tokens[1])
		}
	}

	return nil, fmt.Errorf("%s/%s is not supported by fetcher", kind, external)
}

func (f *IAMFetcher) fetchProject(ctx context.Context, projectID string) (*ResourceInfo, error) {
	log := klog.FromContext(ctx)

	req := &iampb.GetIamPolicyRequest{
		Resource: fmt.Sprintf("projects/%s", projectID),
	}

	policy, err := f.crmClient.GetIamPolicy(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("fetching iam policy for %q: %w", req.Resource, err)
	}
	log.V(2).Info("got iam policy for project", "policy", policy, "projectID", projectID)
	etag := base64.StdEncoding.EncodeToString(policy.Etag)
	return &ResourceInfo{
		Etag: etag,
	}, nil
}
