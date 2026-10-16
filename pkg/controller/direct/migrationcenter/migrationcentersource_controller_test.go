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

package migrationcenter

import (
	"context"
	"testing"

	gcp "cloud.google.com/go/migrationcenter/apiv1"
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestMigrationCenterSource_AdapterForURL(t *testing.T) {
	ctx := context.Background()
	model := &sourceModel{
		config:    config.ControllerConfig{},
		gcpClient: &gcp.Client{},
	}

	// Test invalid URL returns nil, nil
	adapter, err := model.AdapterForURL(ctx, "invalid-url")
	if err != nil {
		t.Fatalf("unexpected error for invalid url: %v", err)
	}
	if adapter != nil {
		t.Fatalf("expected nil adapter for invalid url, got: %v", adapter)
	}

	// Test valid external URL parses identity
	url := "//migrationcenter.googleapis.com/projects/test-project/locations/us-central1/sources/test-source"
	adapter, err = model.AdapterForURL(ctx, url)
	if err != nil {
		t.Fatalf("unexpected error for valid url: %v", err)
	}
	if adapter == nil {
		t.Fatalf("expected non-nil adapter for valid url")
	}
	sAdapter, ok := adapter.(*sourceAdapter)
	if !ok {
		t.Fatalf("expected *sourceAdapter, got: %T", adapter)
	}
	if sAdapter.id.Project != "test-project" || sAdapter.id.Location != "us-central1" || sAdapter.id.Source != "test-source" {
		t.Fatalf("unexpected parsed identity: %+v", sAdapter.id)
	}
}

func TestMigrationCenterSource_Export(t *testing.T) {
	ctx := context.Background()
	adapter := &sourceAdapter{
		id: &krm.MigrationCenterSourceIdentity{
			Project:  "test-project",
			Location: "us-central1",
			Source:   "test-source",
		},
		actual: &pb.Source{
			Name:        "projects/test-project/locations/us-central1/sources/test-source",
			DisplayName: "Test Display Name",
			Description: "Test Description",
			Type:        pb.Source_SOURCE_TYPE_CUSTOM,
			Priority:    10,
			Managed:     true,
		},
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() returned error: %v", err)
	}

	if u.GetName() != "test-source" {
		t.Errorf("Export() name = %q, want %q", u.GetName(), "test-source")
	}
	if u.GroupVersionKind() != krm.MigrationCenterSourceGVK {
		t.Errorf("Export() GVK = %v, want %v", u.GroupVersionKind(), krm.MigrationCenterSourceGVK)
	}

	var obj krm.MigrationCenterSource
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		t.Fatalf("failed to convert unstructured to MigrationCenterSource: %v", err)
	}

	if obj.Spec.ProjectRef == nil || obj.Spec.ProjectRef.External != "test-project" {
		t.Errorf("Export() ProjectRef.External = %v, want test-project", obj.Spec.ProjectRef)
	}
	if obj.Spec.Location == nil || *obj.Spec.Location != "us-central1" {
		t.Errorf("Export() Location = %v, want us-central1", obj.Spec.Location)
	}
	if obj.Spec.ResourceID == nil || *obj.Spec.ResourceID != "test-source" {
		t.Errorf("Export() ResourceID = %v, want test-source", obj.Spec.ResourceID)
	}
	if obj.Spec.DisplayName == nil || *obj.Spec.DisplayName != "Test Display Name" {
		t.Errorf("Export() DisplayName = %v, want 'Test Display Name'", obj.Spec.DisplayName)
	}
	if obj.Spec.Description == nil || *obj.Spec.Description != "Test Description" {
		t.Errorf("Export() Description = %v, want 'Test Description'", obj.Spec.Description)
	}
	if obj.Spec.Type == nil || *obj.Spec.Type != "SOURCE_TYPE_CUSTOM" {
		t.Errorf("Export() Type = %v, want 'SOURCE_TYPE_CUSTOM'", obj.Spec.Type)
	}
	if obj.Spec.Priority == nil || *obj.Spec.Priority != 10 {
		t.Errorf("Export() Priority = %v, want 10", obj.Spec.Priority)
	}
	if obj.Spec.Managed == nil || *obj.Spec.Managed != true {
		t.Errorf("Export() Managed = %v, want true", obj.Spec.Managed)
	}

	// Verify Find() not called error
	uninitializedAdapter := &sourceAdapter{}
	if _, err := uninitializedAdapter.Export(ctx); err == nil {
		t.Errorf("expected error when actual is nil")
	}
}
