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

package vertexai

import (
	"context"
	"net/http"
	"testing"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestFeaturestoreEntityType_AdapterForURL(t *testing.T) {
	ctx := context.Background()
	m := &modelFeaturestoreEntityType{
		config: config.ControllerConfig{
			HTTPClient: &http.Client{},
		},
	}

	t.Run("valid URL", func(t *testing.T) {
		url := "//aiplatform.googleapis.com/projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et"
		adapter, err := m.AdapterForURL(ctx, url)
		if err != nil {
			t.Fatalf("AdapterForURL() unexpected error: %v", err)
		}
		if adapter == nil {
			t.Fatalf("AdapterForURL() returned nil adapter")
		}

		etAdapter, ok := adapter.(*FeaturestoreEntityTypeAdapter)
		if !ok {
			t.Fatalf("expected *FeaturestoreEntityTypeAdapter, got %T", adapter)
		}
		if etAdapter.id.Project != "my-project" ||
			etAdapter.id.Location != "us-central1" ||
			etAdapter.id.Featurestore != "my-fs" ||
			etAdapter.id.EntityType != "my-et" {
			t.Errorf("unexpected identity parsed: %+v", etAdapter.id)
		}
	})

	t.Run("invalid URL returns nil without error", func(t *testing.T) {
		url := "https://compute.googleapis.com/compute/v1/projects/p/zones/z/instances/i"
		adapter, err := m.AdapterForURL(ctx, url)
		if err != nil {
			t.Fatalf("AdapterForURL() unexpected error: %v", err)
		}
		if adapter != nil {
			t.Errorf("expected nil adapter for invalid url, got %v", adapter)
		}
	})
}

type fakeEntityTypeOp struct {
	directbase.Operation
	status *krm.VertexAIFeaturestoreEntityTypeStatus
}

func (f *fakeEntityTypeOp) UpdateStatus(ctx context.Context, status any, readyCondition *v1alpha1.Condition) error {
	f.status = status.(*krm.VertexAIFeaturestoreEntityTypeStatus)
	return nil
}

func TestFeaturestoreEntityType_UpdateStatus(t *testing.T) {
	ctx := context.Background()
	adapter := &FeaturestoreEntityTypeAdapter{
		id: &krm.VertexAIFeaturestoreEntityTypeIdentity{
			Project:      "my-project",
			Location:     "us-central1",
			Featurestore: "my-fs",
			EntityType:   "my-et",
		},
	}

	t.Run("handles nil proto gracefully", func(t *testing.T) {
		op := &fakeEntityTypeOp{}
		if err := adapter.updateStatus(ctx, op, nil); err != nil {
			t.Fatalf("updateStatus unexpected error: %v", err)
		}
		if op.status == nil {
			t.Fatalf("expected non-nil status")
		}
		if op.status.Region == nil || *op.status.Region != "us-central1" {
			t.Errorf("expected Region=us-central1, got %v", op.status.Region)
		}
	})

	t.Run("maps proto status correctly", func(t *testing.T) {
		op := &fakeEntityTypeOp{}
		proto := &pb.EntityType{
			CreateTime: timestamppb.Now(),
		}
		if err := adapter.updateStatus(ctx, op, proto); err != nil {
			t.Fatalf("updateStatus unexpected error: %v", err)
		}
		if op.status == nil || op.status.CreateTime == nil {
			t.Fatalf("expected populated status with CreateTime")
		}
		if op.status.Region == nil || *op.status.Region != "us-central1" {
			t.Errorf("expected Region=us-central1, got %v", op.status.Region)
		}
	})
}
