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

package dataplex

import (
	"slices"
	"testing"

	dataplexpb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
)

func TestDataplexAsset_StorageBucketDiffConvergence(t *testing.T) {
	t.Run("no_diff_when_gcp_returns_bare_bucket_name", func(t *testing.T) {
		desired := &krm.DataplexAssetSpec{
			Description: direct.LazyPtr("Initial description"),
			ResourceSpec: &krm.Asset_ResourceSpec{
				Type:           direct.LazyPtr("STORAGE_BUCKET"),
				ReadAccessMode: direct.LazyPtr("DIRECT"),
				StorageBucketRef: &storagev1beta1.StorageBucketRef{
					External: "projects/my-proj/buckets/my-bucket",
				},
			},
		}
		actual := &dataplexpb.Asset{
			Description: "Initial description",
			ResourceSpec: &dataplexpb.Asset_ResourceSpec{
				Type:           dataplexpb.Asset_ResourceSpec_STORAGE_BUCKET,
				ReadAccessMode: dataplexpb.Asset_ResourceSpec_DIRECT,
				Name:           "my-bucket",
			},
		}

		merged, diff, updateMask, err := common.CompareGreenfieldSpec(
			t.Context(),
			desired,
			actual,
			DataplexAssetSpec_FromProto,
			DataplexAssetSpec_ToProto,
			normalizeAssetForDiff,
		)

		if err != nil {
			t.Fatalf("CompareGreenfieldSpec() unexpected error: %v", err)
		}
		if diff.HasDiff() || len(updateMask.GetPaths()) != 0 {
			t.Errorf("CompareGreenfieldSpec() paths = %v, want empty updateMask", updateMask.GetPaths())
		}
		if merged.GetResourceSpec().GetName() != "projects/my-proj/buckets/my-bucket" {
			t.Errorf("CompareGreenfieldSpec() merged ResourceSpec.Name = %q, want %q", merged.GetResourceSpec().GetName(), "projects/my-proj/buckets/my-bucket")
		}
	})

	t.Run("description_update_excludes_resource_spec_from_mask", func(t *testing.T) {
		desired := &krm.DataplexAssetSpec{
			Description: direct.LazyPtr("Updated description"),
			ResourceSpec: &krm.Asset_ResourceSpec{
				Type:           direct.LazyPtr("STORAGE_BUCKET"),
				ReadAccessMode: direct.LazyPtr("DIRECT"),
				StorageBucketRef: &storagev1beta1.StorageBucketRef{
					External: "projects/my-proj/buckets/my-bucket",
				},
			},
		}
		actual := &dataplexpb.Asset{
			Description: "Initial description",
			ResourceSpec: &dataplexpb.Asset_ResourceSpec{
				Type:           dataplexpb.Asset_ResourceSpec_STORAGE_BUCKET,
				ReadAccessMode: dataplexpb.Asset_ResourceSpec_DIRECT,
				Name:           "my-bucket",
			},
		}

		_, _, updateMask, err := common.CompareGreenfieldSpec(
			t.Context(),
			desired,
			actual,
			DataplexAssetSpec_FromProto,
			DataplexAssetSpec_ToProto,
			normalizeAssetForDiff,
		)

		if err != nil {
			t.Fatalf("CompareGreenfieldSpec() unexpected error: %v", err)
		}
		if slices.Contains(updateMask.GetPaths(), "resource_spec") {
			t.Errorf("CompareGreenfieldSpec() paths = %v, did not want resource_spec", updateMask.GetPaths())
		}
		if !slices.Contains(updateMask.GetPaths(), "description") {
			t.Errorf("CompareGreenfieldSpec() paths = %v, want description", updateMask.GetPaths())
		}
	})

	t.Run("detects_immutable_bucket_change", func(t *testing.T) {
		desired := &krm.DataplexAssetSpec{
			ResourceSpec: &krm.Asset_ResourceSpec{
				Type:           direct.LazyPtr("STORAGE_BUCKET"),
				ReadAccessMode: direct.LazyPtr("DIRECT"),
				StorageBucketRef: &storagev1beta1.StorageBucketRef{
					External: "projects/my-proj/buckets/different-bucket",
				},
			},
		}
		actual := &dataplexpb.Asset{
			ResourceSpec: &dataplexpb.Asset_ResourceSpec{
				Type:           dataplexpb.Asset_ResourceSpec_STORAGE_BUCKET,
				ReadAccessMode: dataplexpb.Asset_ResourceSpec_DIRECT,
				Name:           "my-bucket",
			},
		}

		_, diff, updateMask, err := common.CompareGreenfieldSpec(
			t.Context(),
			desired,
			actual,
			DataplexAssetSpec_FromProto,
			DataplexAssetSpec_ToProto,
			normalizeAssetForDiff,
		)

		if err != nil {
			t.Fatalf("CompareGreenfieldSpec() unexpected error: %v", err)
		}
		if !diff.HasDiff() {
			t.Errorf("CompareGreenfieldSpec() HasDiff() = false, want true")
		}
		if !slices.Contains(updateMask.GetPaths(), "resource_spec") {
			t.Errorf("CompareGreenfieldSpec() paths = %v, want resource_spec", updateMask.GetPaths())
		}
	})
}

func TestAssetAdapter_ResolveStorageBucketProject(t *testing.T) {
	assetID := &krm.DataplexAssetIdentity{
		Project:  "asset-proj",
		Location: "us-central1",
		Lake:     "my-lake",
		Zone:     "my-zone",
		Asset:    "my-asset",
	}

	tests := []struct {
		name     string
		id       *krm.DataplexAssetIdentity
		specType dataplexpb.Asset_ResourceSpec_Type
		specName string
		wantName string
	}{
		{
			name:     "wildcard_project_replaced",
			id:       assetID,
			specType: dataplexpb.Asset_ResourceSpec_STORAGE_BUCKET,
			specName: "projects/_/buckets/bare-bucket",
			wantName: "projects/asset-proj/buckets/bare-bucket",
		},
		{
			name:     "explicit_project_preserved",
			id:       assetID,
			specType: dataplexpb.Asset_ResourceSpec_STORAGE_BUCKET,
			specName: "projects/other-proj/buckets/bucket",
			wantName: "projects/other-proj/buckets/bucket",
		},
		{
			name:     "bigquery_dataset_untouched",
			id:       assetID,
			specType: dataplexpb.Asset_ResourceSpec_BIGQUERY_DATASET,
			specName: "projects/other-proj/datasets/my_dataset",
			wantName: "projects/other-proj/datasets/my_dataset",
		},
		{
			name:     "nil_identity_noop",
			specType: dataplexpb.Asset_ResourceSpec_STORAGE_BUCKET,
			specName: "projects/_/buckets/bare-bucket",
			wantName: "projects/_/buckets/bare-bucket",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &AssetAdapter{id: tc.id}
			asset := &dataplexpb.Asset{
				ResourceSpec: &dataplexpb.Asset_ResourceSpec{
					Type: tc.specType,
					Name: tc.specName,
				},
			}

			adapter.resolveStorageBucketProject(asset)

			if got := asset.GetResourceSpec().GetName(); got != tc.wantName {
				t.Errorf("resolveStorageBucketProject(%q) Name = %q, want %q", tc.specName, got, tc.wantName)
			}
		})
	}
}
