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

package resourceoverrides

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetAppEngineStandardAppVersionResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "AppEngineStandardAppVersion",
	}
	ro.Overrides = append(ro.Overrides, handleAppEngineServiceRef())
	return ro
}

// handleAppEngineServiceRef ensures that serviceRef in format "apps/{app}/services/{service}"
// is trimmed to just "{service}" before Terraform actuation, and restored after actuation.
func handleAppEngineServiceRef() ResourceOverride {
	o := ResourceOverride{}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		val, found, err := unstructured.NestedString(r.Spec, "serviceRef", "external")
		if err != nil {
			return err
		}
		if !found || val == "" {
			return nil
		}
		if strings.Contains(val, "/services/") {
			parts := strings.Split(val, "/services/")
			service := parts[len(parts)-1]
			if err := unstructured.SetNestedField(r.Spec, service, "serviceRef", "external"); err != nil {
				return fmt.Errorf("error setting serviceRef.external: %w", err)
			}
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		return PreserveUserSpecifiedLegacyField(original, reconciled, "serviceRef", "external")
	}
	o.PostUpdateStatusTransform = func(r *k8s.Resource) error {
		val, found, err := unstructured.NestedString(r.Spec, "serviceRef", "external")
		if err != nil {
			return err
		}
		if !found || val == "" {
			return nil
		}
		if strings.Contains(val, "/services/") {
			parts := strings.Split(val, "/services/")
			service := parts[len(parts)-1]
			if err := unstructured.SetNestedField(r.Spec, service, "serviceRef", "external"); err != nil {
				return fmt.Errorf("error setting serviceRef.external in post-status update: %w", err)
			}
		}
		return nil
	}
	return o
}
