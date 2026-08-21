// Copyright 2022 Google LLC
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

// package serviceclient contains clients that are used
// to interact with various GCP services.
package serviceclient

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudasset/v1"
	resourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	cloudstorage "google.golang.org/api/storage/v1"
)

type serviceClient struct {
	httpClient *http.Client
}

type ServiceClient interface {
	GetProjectFromProjectIDOrNumber(string) (*resourcemanager.Project, error)
}

func NewServiceClient(httpClient *http.Client) serviceClient { //nolint:revive
	return serviceClient{httpClient: httpClient}
}

// GetProjectFromIdOrNumber retrieves a project, given a project ID or number.
func (s *serviceClient) GetProjectFromProjectIDOrNumber(projectIDOrNumber string) (*resourcemanager.Project, error) {
	resourceManagerClient, err := NewResourceManagerClient(context.Background(), s.httpClient)
	if err != nil {
		return nil, fmt.Errorf("unable to create resource manager client: %w", err)
	}
	// Technically this API is supposed to only supports the project ID, but
	// have verified that the project number also work.
	//
	// As of this commit, there is no API that officially supports retrieval of
	// project information by project number.
	project, err := resourceManagerClient.Projects.Get(projectIDOrNumber).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve project with project id or number %v : %w", projectIDOrNumber, err)
	}
	return project, nil
}

func NewHTTPClient(ctx context.Context, oAuth2Token string) (*http.Client, error) {
	if oAuth2Token == "" {
		client, err := google.DefaultClient(ctx, cloudasset.CloudPlatformScope, cloudstorage.CloudPlatformScope)
		if err != nil {
			return nil, fmt.Errorf("error creating default gcp client: %w", err)
		}
		return client, nil
	}
	token := oauth2.Token{AccessToken: oAuth2Token}
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&token))
	return client, nil
}

func NewResourceManagerClient(ctx context.Context, httpClient *http.Client) (*resourcemanager.Service, error) {
	client, err := resourcemanager.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("error creating resource manager client: %w", err)
	}
	return client, nil
}
