// Copyright 2024 Google LLC
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

package direct

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ExportUsesDirect returns true if we have a direct-implemented exporter for the URL.
func ExportUsesDirect(ctx context.Context, url string, config *config.ControllerConfig) (bool, error) {
	adapter, err := registry.AdapterForURL(ctx, url)
	if err != nil {
		return false, err
	}
	return adapter != nil, nil
}

// Export attempts to export the resource specified by url.
// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
// If url is not recognized or not implemented by a direct controller, this returns (nil, nil)
func Export(ctx context.Context, url string, config *config.ControllerConfig) (*unstructured.Unstructured, error) {
	adapter, err := registry.AdapterForURL(ctx, url)
	if err != nil {
		return nil, err
	}
	if adapter != nil {
		found, err := adapter.Find(ctx)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("resource %q is not found", url)
		}

		u, err := adapter.Export(ctx)
		if err != nil {
			return nil, err
		}

		return u, nil
	}

	return nil, nil
}
