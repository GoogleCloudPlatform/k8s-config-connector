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

package gcp

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/version"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/storage/v1"
)

// The user agent to track KCC's attribution to GCP usages
func KCCUserAgent() string {
	kccVersion := version.GetVersion()
	// Note: try to keep in sync with third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/fwtransport/framework_utils.go
	userAgent := fmt.Sprintf("kcc/%s (+https://github.com/GoogleCloudPlatform/k8s-config-connector) kcc/controller-manager/%s", kccVersion, kccVersion)
	return userAgent
}

func NewIAMClient(ctx context.Context) (*iam.Service, error) {
	httpClient, err := google.DefaultClient(ctx, iam.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	client, err := iam.New(httpClient)
	if err != nil {
		return nil, err
	}
	client.UserAgent = KCCUserAgent()
	return client, nil
}

func NewStorageClient(ctx context.Context) (*storage.Service, error) {
	httpClient, err := google.DefaultClient(ctx, storage.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	client, err := storage.New(httpClient)
	if err != nil {
		return nil, err
	}
	client.UserAgent = KCCUserAgent()
	return client, nil
}

// NewCloudResourceManagerClient returns a GCP Cloud Resource Manager service.
func NewCloudResourceManagerClient(ctx context.Context) (*cloudresourcemanager.Service, error) {
	httpClient, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	client, err := cloudresourcemanager.New(httpClient)
	if err != nil {
		return nil, err
	}
	client.UserAgent = KCCUserAgent()
	return client, nil
}
