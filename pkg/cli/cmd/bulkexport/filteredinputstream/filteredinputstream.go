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
	"fmt"
	"regexp"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/log"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/klog/v2"
)

var defaultNetworkingNameRegexByAssetType = map[string]string{
	"compute.googleapis.com/Network":    ".*networks/default$",
	"compute.googleapis.com/Subnetwork": ".*subnetworks/default$",
	"compute.googleapis.com/Route":      ".*default-route-.*$",
}

func isAssetSupported(ctx context.Context, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider, config *config.ControllerConfig, a *asset.Asset) bool {
	log := klog.FromContext(ctx)

	// First check if this resource uses our direct-reconciliation model
	exportUsesDirect, err := direct.ExportUsesDirect(ctx, a.Name, config)
	if err != nil {
		log.Error(err, "checking if resource is direct-implemented", "url", a.Name)
	} else if exportUsesDirect {
		return true
	}

	_, rc, err := asset.GetServiceMappingAndResourceConfig(smLoader, a)
	if err != nil {
		// ignore resources that we don't have service mappings and resource configs for
		return false
	}
	// every value for rc.Name should be in the ResourcesMap
	resource := tfProvider.ResourcesMap[rc.Name]
	return resource.Importer != nil
}

// isDefaultNetworkingAsset returns true if the asset is a default networking asset.
// Default networking assets are not always acquirable via Config Connector or Terraform
// due to their non-standard configuration.
func isDefaultNetworkingAsset(a *asset.Asset) bool {
	if defaultRegexMatch, ok := defaultNetworkingNameRegexByAssetType[a.AssetType]; ok {
		matched, err := regexp.MatchString(defaultRegexMatch, a.Name)
		if err != nil {
			log.Error("unable to match string to detect networking asset type: regex '%v', string '%v', err '%v'", defaultRegexMatch, a.Name, err)
			return false
		}
		return matched
	}
	return false
}

func NewFilteredAssetStream(ctx context.Context, assetStream *asset.Stream, tfProvider *schema.Provider, config *config.ControllerConfig) (stream.AssetStream, error) {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		return nil, fmt.Errorf("error loading service mappings: %w", err)
	}
	filter := func(a *asset.Asset) bool {
		if !isAssetSupported(ctx, smLoader, tfProvider, config, a) {
			log.Verbose("skipping unsupported asset: %v", a.AssetType)
			return false
		}
		if isDefaultNetworkingAsset(a) {
			log.Verbose("skipping default asset, as it cannot be normally acquired or imported: %v/%v", a.AssetType, a.Name)
			return false
		}
		return true
	}
	return stream.NewFilteredAssetStream(assetStream, filter), nil
}
