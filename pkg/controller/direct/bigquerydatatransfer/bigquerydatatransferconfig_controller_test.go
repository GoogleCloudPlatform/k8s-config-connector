// Copyright 2026 Google LLC
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

package bigquerydatatransfer

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	bigquerydatatransferpb "cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

type recordingRoundTripper struct {
	lastRequest *http.Request
}

func (r *recordingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	r.lastRequest = req
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("{}")),
	}, nil
}

func TestClientWithDefaultQuotaProject(t *testing.T) {
	tests := []struct {
		name                  string
		userProjectOverride   bool
		billingProject        string
		targetProject         string
		wantUserProjectHeader string
	}{
		{
			name:                  "RESOURCE_PROJECT policy (UserProjectOverride=true, BillingProject empty)",
			userProjectOverride:   true,
			billingProject:        "",
			targetProject:         "test-target-project",
			wantUserProjectHeader: "test-target-project",
		},
		{
			name:                  "BILLING_PROJECT policy (UserProjectOverride=true, BillingProject set)",
			userProjectOverride:   true,
			billingProject:        "custom-billing-project",
			targetProject:         "test-target-project",
			wantUserProjectHeader: "custom-billing-project",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			rt := &recordingRoundTripper{}
			httpClient := &http.Client{Transport: rt}

			controllerCfg := &config.ControllerConfig{
				UserProjectOverride: tc.userProjectOverride,
				BillingProject:      tc.billingProject,
				HTTPClient:          httpClient,
			}

			m, err := NewModel(ctx, controllerCfg)
			if err != nil {
				t.Fatalf("NewModel() failed: %v", err)
			}

			modelImpl := m.(*model)
			client, err := modelImpl.client(ctx, tc.targetProject)
			if err != nil {
				t.Fatalf("model.client() failed: %v", err)
			}
			defer client.Close()

			req := &bigquerydatatransferpb.GetTransferConfigRequest{
				Name: "projects/test-target-project/locations/us/transferConfigs/12345",
			}
			_, _ = client.GetTransferConfig(ctx, req)

			if rt.lastRequest == nil {
				t.Fatalf("expected HTTP request to be recorded, but got nil")
			}

			got := rt.lastRequest.Header.Get("X-goog-user-project")
			if got != tc.wantUserProjectHeader {
				t.Errorf("X-goog-user-project header = %q, want %q", got, tc.wantUserProjectHeader)
			}
		})
	}
}
