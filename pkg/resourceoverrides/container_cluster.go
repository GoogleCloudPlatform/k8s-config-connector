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

package resourceoverrides

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var (
	identityNamespaceFieldPath = []string{"workloadIdentityConfig", "identityNamespace"}
	workloadPoolFieldPath      = []string{"workloadIdentityConfig", "workloadPool"}
)

func GetContainerClusterResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ContainerCluster",
	}
	// Preserve the legacy field 'workloadIdentityConfig.identityNamespace' that has been removed from Terraform 4.x upgrade.
	// See b/206133327 for context.
	ro.Overrides = append(ro.Overrides, keepIdentityNamespaceField())
	return ro
}

func keepIdentityNamespaceField() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		workloadIdentityConfig := schema.Properties["spec"].Properties["workloadIdentityConfig"]
		workloadIdentityConfig.Properties["identityNamespace"] = apiextensions.JSONSchemaProps{
			Description: "DEPRECATED — This field will be removed in a future major release as it has been deprecated in the API. Use `workloadPool` instead; `workloadPool` field will supersede this field.\n" +
				"Enables workload identity.",
			Type: "string",
		}
		return nil
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := FavorAuthoritativeFieldOverLegacyField(r, identityNamespaceFieldPath, workloadPoolFieldPath); err != nil {
			return fmt.Errorf("error handling '%v' and '%v' fields in pre-actuation transformation: %w", strings.Join(identityNamespaceFieldPath, "."), strings.Join(workloadPoolFieldPath, "."), err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource) error {
		if err := PreserveUserSpecifiedLegacyField(original, reconciled, identityNamespaceFieldPath...); err != nil {
			return fmt.Errorf("error preserving '%v' in post-actuation transformation: %w", strings.Join(identityNamespaceFieldPath, "."), err)
		}
		if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified(original, reconciled, workloadPoolFieldPath, workloadPoolFieldPath); err != nil {
			return fmt.Errorf("error conditionally pruning '%v' in post-actuation transformation: %w", strings.Join(workloadPoolFieldPath, "."), err)
		}
		return nil
	}
	return o
}
