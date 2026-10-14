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

package saasservicemgmt

import (
	"context"
	"reflect"
	"testing"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/saasservicemgmt/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSaaSServiceMgmtTenant_Export(t *testing.T) {
	ctx := context.Background()

	id := &krm.SaaSServiceMgmtTenantIdentity{
		Project:  "test-project",
		Location: "us-central1",
		Tenant:   "my-tenant-123",
	}

	actual := &pb.Tenant{
		Name:             "projects/test-project/locations/us-central1/tenants/my-tenant-123",
		ConsumerResource: "projects/test-project/locations/us-central1/consumerResources/cr-1",
		Saas:             "projects/test-project/locations/us-central1/saas/my-saas",
		Labels: map[string]string{
			"env":  "production",
			"team": "saas",
		},
	}

	adapter := &SaaSServiceMgmtTenantAdapter{
		id:     id,
		actual: actual,
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() failed: %v", err)
	}
	if u == nil {
		t.Fatal("Export() returned nil unstructured object")
	}

	// Verify Kubernetes object name is the short tenant ID, NOT the full GCP path
	if got, want := u.GetName(), "my-tenant-123"; got != want {
		t.Errorf("GetName() = %q, want %q", got, want)
	}

	// Verify GVK
	if got, want := u.GroupVersionKind(), krm.SaaSServiceMgmtTenantGVK; got != want {
		t.Errorf("GroupVersionKind() = %v, want %v", got, want)
	}

	// Verify Labels
	if got, want := u.GetLabels(), actual.Labels; !reflect.DeepEqual(got, want) {
		t.Errorf("GetLabels() = %v, want %v", got, want)
	}

	// Verify Spec fields (ProjectRef, Location, ConsumerResource, SaasRef)
	spec, found, err := unstructured.NestedMap(u.Object, "spec")
	if err != nil {
		t.Fatalf("NestedMap spec failed: %v", err)
	}
	if !found {
		t.Fatal("spec field not found in unstructured object")
	}

	projectRef, found, _ := unstructured.NestedString(spec, "projectRef", "external")
	if !found || projectRef != "test-project" {
		t.Errorf("spec.projectRef.external = %q, want %q", projectRef, "test-project")
	}

	location, found, _ := unstructured.NestedString(spec, "location")
	if !found || location != "us-central1" {
		t.Errorf("spec.location = %q, want %q", location, "us-central1")
	}

	consumerResource, found, _ := unstructured.NestedString(spec, "consumerResource")
	if !found || consumerResource != "projects/test-project/locations/us-central1/consumerResources/cr-1" {
		t.Errorf("spec.consumerResource = %q, want %q", consumerResource, "projects/test-project/locations/us-central1/consumerResources/cr-1")
	}

	saasRef, found, _ := unstructured.NestedString(spec, "saasRef", "external")
	if !found || saasRef != "projects/test-project/locations/us-central1/saas/my-saas" {
		t.Errorf("spec.saasRef.external = %q, want %q", saasRef, "projects/test-project/locations/us-central1/saas/my-saas")
	}
}

func TestSaaSServiceMgmtTenant_Export_NoLabels(t *testing.T) {
	ctx := context.Background()

	id := &krm.SaaSServiceMgmtTenantIdentity{
		Project:  "test-project",
		Location: "us-central1",
		Tenant:   "tenant-nolabels",
	}

	actual := &pb.Tenant{
		Name: "projects/test-project/locations/us-central1/tenants/tenant-nolabels",
	}

	adapter := &SaaSServiceMgmtTenantAdapter{
		id:     id,
		actual: actual,
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() failed: %v", err)
	}

	if got, want := u.GetName(), "tenant-nolabels"; got != want {
		t.Errorf("GetName() = %q, want %q", got, want)
	}

	if len(u.GetLabels()) != 0 {
		t.Errorf("GetLabels() = %v, want empty", u.GetLabels())
	}
}

func TestSaaSServiceMgmtTenant_Export_FindNotCalled(t *testing.T) {
	ctx := context.Background()

	adapter := &SaaSServiceMgmtTenantAdapter{
		id: &krm.SaaSServiceMgmtTenantIdentity{
			Project:  "test-project",
			Location: "us-central1",
			Tenant:   "tenant-1",
		},
		actual: nil,
	}

	_, err := adapter.Export(ctx)
	if err == nil {
		t.Fatal("Export() expected error when actual is nil, got nil")
	}
}
