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

func TestFeaturestoreEntityTypeFeature_AdapterForURL(t *testing.T) {
	ctx := context.Background()
	m := &modelFeaturestoreEntityTypeFeature{
		config: config.ControllerConfig{
			HTTPClient: &http.Client{},
		},
	}

	t.Run("valid URL", func(t *testing.T) {
		url := "//aiplatform.googleapis.com/projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et/features/my-feat"
		adapter, err := m.AdapterForURL(ctx, url)
		if err != nil {
			t.Fatalf("AdapterForURL() unexpected error: %v", err)
		}
		if adapter == nil {
			t.Fatalf("AdapterForURL() returned nil adapter")
		}

		featAdapter, ok := adapter.(*FeaturestoreEntityTypeFeatureAdapter)
		if !ok {
			t.Fatalf("expected *FeaturestoreEntityTypeFeatureAdapter, got %T", adapter)
		}
		if featAdapter.id.Project != "my-project" ||
			featAdapter.id.Location != "us-central1" ||
			featAdapter.id.Featurestore != "my-fs" ||
			featAdapter.id.EntityType != "my-et" ||
			featAdapter.id.Feature != "my-feat" {
			t.Errorf("unexpected identity parsed: %+v", featAdapter.id)
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

type fakeOp struct {
	directbase.Operation
	status *krm.VertexAIFeaturestoreEntityTypeFeatureStatus
}

func (f *fakeOp) UpdateStatus(ctx context.Context, status any, readyCondition *v1alpha1.Condition) error {
	f.status = status.(*krm.VertexAIFeaturestoreEntityTypeFeatureStatus)
	return nil
}

func TestFeaturestoreEntityTypeFeature_UpdateStatus(t *testing.T) {
	ctx := context.Background()
	adapter := &FeaturestoreEntityTypeFeatureAdapter{
		id: &krm.VertexAIFeaturestoreEntityTypeFeatureIdentity{
			Project:      "my-project",
			Location:     "us-central1",
			Featurestore: "my-fs",
			EntityType:   "my-et",
			Feature:      "my-feat",
		},
	}

	t.Run("handles nil proto gracefully", func(t *testing.T) {
		op := &fakeOp{}
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
		op := &fakeOp{}
		proto := &pb.Feature{
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
