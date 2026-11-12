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

package dlp

import (
	"context"
	"net/http"
	"strings"
	"testing"

	dlppb "cloud.google.com/go/dlp/apiv2/dlppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dlp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestContentPolicyAdapter_Export(t *testing.T) {
	// Arrange
	id := &krm.DLPContentPolicyIdentity{
		Project:       "my-project",
		Location:      "global",
		ContentPolicy: "policy-xyz",
	}
	actual := &dlppb.ContentPolicy{
		Name:        "projects/my-project/locations/global/contentPolicies/policy-xyz",
		DisplayName: "My Exported Content Policy",
		InspectConfig: &dlppb.InspectConfig{
			InfoTypes: []*dlppb.InfoType{
				{Name: "EMAIL_ADDRESS"},
			},
		},
	}
	adapter := newContentPolicyAdapter(id, nil, nil)
	adapter.actual = actual

	// Act
	u, err := adapter.Export(context.Background())

	// Assert
	if err != nil {
		t.Fatalf("Export() failed: %v", err)
	}
	if got, want := u.GetName(), "policy-xyz"; got != want {
		t.Errorf("u.GetName() = %q, want %q (must use resource ID, not full resource name)", got, want)
	}
	if got, want := u.GroupVersionKind(), krm.DLPContentPolicyGVK; got != want {
		t.Errorf("u.GroupVersionKind() = %v, want %v", got, want)
	}
	specMap, ok := u.Object["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec map in unstructured object, got: %v", u.Object["spec"])
	}
	if got, want := specMap["displayName"], "My Exported Content Policy"; got != want {
		t.Errorf("spec.displayName = %v, want %v", got, want)
	}
}

func TestContentPolicyAdapter_Export_UncalledFind(t *testing.T) {
	// Arrange
	id := &krm.DLPContentPolicyIdentity{
		Project:       "my-project",
		Location:      "global",
		ContentPolicy: "policy-xyz",
	}
	adapter := newContentPolicyAdapter(id, nil, nil)

	// Act
	u, err := adapter.Export(context.Background())

	// Assert
	if err == nil {
		t.Fatalf("expected error when Find() was not called, got u=%v", u)
	}
	if !strings.Contains(err.Error(), "Find() not called") {
		t.Errorf("expected error message to contain 'Find() not called', got %v", err)
	}
}

func TestModelContentPolicy_AdapterForURL(t *testing.T) {
	// Arrange
	model := &modelContentPolicy{
		config: config.ControllerConfig{
			HTTPClient: &http.Client{},
		},
	}
	ctx := context.Background()

	// Act & Assert for valid relative URL
	urlValidSlash := "//dlp.googleapis.com/projects/my-project/locations/global/contentPolicies/my-policy"
	adapterSlash, errSlash := model.AdapterForURL(ctx, urlValidSlash)

	// Assert
	if errSlash != nil {
		t.Fatalf("unexpected error for %q: %v", urlValidSlash, errSlash)
	}
	if adapterSlash == nil {
		t.Fatalf("expected non-nil adapter for %q", urlValidSlash)
	}
	cpAdapterSlash, ok := adapterSlash.(*ContentPolicyAdapter)
	if !ok {
		t.Fatalf("expected *ContentPolicyAdapter, got %T", adapterSlash)
	}
	if cpAdapterSlash.id.ContentPolicy != "my-policy" || cpAdapterSlash.id.Project != "my-project" {
		t.Errorf("parsed identity mismatch: got %+v", cpAdapterSlash.id)
	}

	// Arrange: valid HTTPS URL
	urlValidHTTPS := "https://dlp.googleapis.com/projects/my-project/locations/global/contentPolicies/my-policy"

	// Act
	adapterHTTPS, errHTTPS := model.AdapterForURL(ctx, urlValidHTTPS)

	// Assert
	if errHTTPS != nil {
		t.Fatalf("unexpected error for %q: %v", urlValidHTTPS, errHTTPS)
	}
	if adapterHTTPS == nil {
		t.Fatalf("expected non-nil adapter for %q", urlValidHTTPS)
	}
	cpAdapterHTTPS, ok := adapterHTTPS.(*ContentPolicyAdapter)
	if !ok {
		t.Fatalf("expected *ContentPolicyAdapter, got %T", adapterHTTPS)
	}
	if cpAdapterHTTPS.id.ContentPolicy != "my-policy" || cpAdapterHTTPS.id.Project != "my-project" {
		t.Errorf("parsed identity mismatch: got %+v", cpAdapterHTTPS.id)
	}

	// Arrange: unrecognized host prefix
	urlOtherHost := "https://compute.googleapis.com/compute/v1/projects/my-project/global/firewalls/default"

	// Act
	adapterOther, errOther := model.AdapterForURL(ctx, urlOtherHost)

	// Assert
	if errOther != nil {
		t.Errorf("expected nil error for unrecognized host, got %v", errOther)
	}
	if adapterOther != nil {
		t.Errorf("expected nil adapter for unrecognized host, got %v", adapterOther)
	}

	// Arrange: malformed URL for recognized host
	urlMalformed := "//dlp.googleapis.com/not-a-valid-resource-path"

	// Act
	adapterMalformed, errMalformed := model.AdapterForURL(ctx, urlMalformed)

	// Assert
	if errMalformed == nil {
		t.Errorf("expected error for malformed URL %q, got nil error and adapter=%v", urlMalformed, adapterMalformed)
	}
	if adapterMalformed != nil {
		t.Errorf("expected nil adapter on error, got %v", adapterMalformed)
	}
}

func TestNewContentPolicyAdapter_Constructor(t *testing.T) {
	// Arrange
	id := &krm.DLPContentPolicyIdentity{
		Project:       "test-project",
		Location:      "global",
		ContentPolicy: "policy-test",
	}
	desired := &krm.DLPContentPolicy{}
	desired.Spec.DisplayName = direct.LazyPtr("Test")

	// Act
	adapter := newContentPolicyAdapter(id, nil, desired)

	// Assert
	if adapter == nil {
		t.Fatalf("expected non-nil adapter")
	}
	if adapter.id != id {
		t.Errorf("adapter.id = %v, want %v", adapter.id, id)
	}
	if adapter.desired != desired {
		t.Errorf("adapter.desired = %v, want %v", adapter.desired, desired)
	}
	if adapter.gcpClient != nil {
		t.Errorf("adapter.gcpClient = %v, want nil", adapter.gcpClient)
	}
}
