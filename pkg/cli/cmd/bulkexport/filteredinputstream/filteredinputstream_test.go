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

package filteredinputstream

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
)

func TestServiceAccountKey(t *testing.T) {
	testAssetType(t, "iam.googleapis.com/ServiceAccountKey", false)
}

func TestServiceAccount(t *testing.T) {
	testAssetType(t, "iam.googleapis.com/ServiceAccount", true)
}

func testAssetType(t *testing.T, assetType string, expectedResult bool) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	saKeyAsset := asset.Asset{
		AssetType: assetType,
	}

	config := &config.ControllerConfig{}

	result := isAssetSupported(ctx, smLoader, tfProvider, config, &saKeyAsset)
	if result != expectedResult {
		t.Fatalf("unexpected result for service '%v' asset: got '%v', want '%v'", assetType, result, expectedResult)
	}
}

func TestIsDefaultNetworkingAsset(t *testing.T) {
	testCases := []struct {
		Name                     string
		Asset                    *asset.Asset
		IsDefaultNetworkingAsset bool
	}{
		{
			Name: "default network",
			Asset: &asset.Asset{
				Name:      "//compute.googleapis.com/projects/foo/global/networks/default",
				AssetType: "compute.googleapis.com/Network",
				Ancestors: []string{
					"projects/1",
					"organizations/1",
				},
			},
			IsDefaultNetworkingAsset: true,
		},
		{
			Name: "non-default network",
			Asset: &asset.Asset{
				Name:      "//compute.googleapis.com/projects/foo/global/networks/bar",
				AssetType: "compute.googleapis.com/Network",
				Ancestors: []string{
					"projects/1",
					"organizations/1",
				},
			},
			IsDefaultNetworkingAsset: false,
		},
		{
			Name: "default route",
			Asset: &asset.Asset{
				Name:      "//compute.googleapis.com/projects/foo-project/global/routes/default-route-1200b933111e760d",
				AssetType: "compute.googleapis.com/Route",
				Ancestors: []string{
					"projects/1",
					"organizations/1",
				},
			},
			IsDefaultNetworkingAsset: true,
		},
		{
			Name: "non-default route",
			Asset: &asset.Asset{
				Name:      "//compute.googleapis.com/projects/foo-project/global/routes/bar",
				AssetType: "compute.googleapis.com/Route",
				Ancestors: []string{
					"projects/1",
					"organizations/1",
				},
			},
			IsDefaultNetworkingAsset: false,
		},
		{
			Name: "default subnetwork",
			Asset: &asset.Asset{
				Name:      "//compute.googleapis.com/projects/foo-project/regions/asia-southeast2/subnetworks/default",
				AssetType: "compute.googleapis.com/Subnetwork",
				Ancestors: []string{
					"projects/1",
					"organizations/1",
				},
			},
			IsDefaultNetworkingAsset: true,
		},
		{
			Name: "non-default subnetwork",
			Asset: &asset.Asset{
				Name:      "//compute.googleapis.com/projects/default-project-12/regions/asia-southeast2/subnetworks/foo",
				AssetType: "compute.googleapis.com/Subnetwork",
				Ancestors: []string{
					"projects/1",
					"organizations/1",
				},
			},
			IsDefaultNetworkingAsset: false,
		},
		{
			Name: "non-networking resource",
			Asset: &asset.Asset{
				Name:      "//iam.googleapis.com/projects/project-foo/roles/CustomRole",
				AssetType: "iam.googleapis.com/Role",
				Ancestors: []string{
					"projects/1",
					"organizations/1",
				},
			},
			IsDefaultNetworkingAsset: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := isDefaultNetworkingAsset(tc.Asset)
			if got != tc.IsDefaultNetworkingAsset {
				t.Errorf("isDefaultNetworkingAsset(..., '%v') = '%v', want '%v'", tc.Asset, got, tc.IsDefaultNetworkingAsset)
			}
		})
	}
}
