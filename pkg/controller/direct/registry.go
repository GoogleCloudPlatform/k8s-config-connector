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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/logging"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// IsDirect returns true if this resource uses the direct-reconciliation model.
func IsDirect(groupKind schema.GroupKind) bool {
	switch groupKind {
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
		return true
	}
	return false
}

// SupportsIAM returns true if this resource supports IAM (not all GCP resources do).
// An error will be returned if IsDirect(groupKind) is not true.
func SupportsIAM(groupKind schema.GroupKind) (bool, error) {
	switch groupKind {
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
		return false, nil
	}
	return false, fmt.Errorf("groupKind %v is not recognized as a direct kind", groupKind)
}

// Export attempts to export the resource specified by url.
// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
// If url is not recognized or not implemented by a direct controller, this returns (nil, nil)
func Export(ctx context.Context, url string, config *controller.Config) (*unstructured.Unstructured, error) {
	if strings.HasPrefix(url, "//logging.googleapis.com/") {
		tokens := strings.Split(strings.TrimPrefix(url, "//logging.googleapis.com/"), "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "metrics" {
			m := logging.GetModel(config)
			in := &unstructured.Unstructured{}
			in.SetName(tokens[3])
			if err := unstructured.SetNestedField(in.Object, tokens[1], "spec", "projectRef", "external"); err != nil {
				return nil, err
			}

			var reader client.Reader // TODO: Create erroring reader?
			a, err := m.AdapterForObject(ctx, reader, in)
			if err != nil {
				return nil, err
			}
			found, err := a.Find(ctx)
			if err != nil {
				return nil, err
			}
			if !found {
				return nil, fmt.Errorf("resource %q is not found", url)
			}

			u, err := a.Export(ctx)
			if err != nil {
				return nil, err
			}

			return u, nil
		}
	}
	return nil, nil
}
