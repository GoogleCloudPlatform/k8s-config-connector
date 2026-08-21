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

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	bucketPolicyOnlyFieldPath         = []string{"bucketPolicyOnly"}
	uniformBucketLevelAccessFieldPath = []string{"uniformBucketLevelAccess"}
)

func GetStorageBucketResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "StorageBucket",
	}
	// Preserve the legacy field 'bucketPolicyOnly' that has been removed from Terraform 4.x upgrade.
	// See b/206156139 for context.
	ro.Overrides = append(ro.Overrides, preserveBucketPolicyOnlyField())
	// Keep the 'location' field optional and default it to 'US' for backwards compatibility.
	// See b/206156139 for context.
	ro.Overrides = append(ro.Overrides, keepLocationFieldOptionalWithDefault())
	return ro
}

func preserveBucketPolicyOnlyField() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		spec := schema.Properties["spec"]
		spec.Properties["bucketPolicyOnly"] = apiextensions.JSONSchemaProps{
			Type: "boolean",
			Description: "DEPRECATED. Please use the `uniformBucketLevelAccess` field as this field has been renamed by Google. The `uniformBucketLevelAccess` field will supersede this field.\n" +
				"Enables Bucket PolicyOnly access to a bucket.",
		}
		schema.Properties["spec"] = spec
		return nil
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := FavorAuthoritativeFieldOverLegacyField(r, bucketPolicyOnlyFieldPath, uniformBucketLevelAccessFieldPath); err != nil {
			return fmt.Errorf("error handling '%v' and '%v' fields in pre-actuation transformation: %w", strings.Join(bucketPolicyOnlyFieldPath, "."), strings.Join(uniformBucketLevelAccessFieldPath, "."), err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		if err := PreserveUserSpecifiedLegacyField(original, reconciled, bucketPolicyOnlyFieldPath...); err != nil {
			return fmt.Errorf("error preserving '%v' in post-actuation transformation: %w", strings.Join(bucketPolicyOnlyFieldPath, "."), err)
		}
		if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified(original, reconciled, bucketPolicyOnlyFieldPath, uniformBucketLevelAccessFieldPath); err != nil {
			return fmt.Errorf("error conditionally pruning '%v' in post-actuation transformation: %w", strings.Join(uniformBucketLevelAccessFieldPath, "."), err)
		}
		return nil
	}
	return o
}

func keepLocationFieldOptionalWithDefault() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return KeepTopLevelFieldOptionalWithDefault(crd, "US", "location")
	}
	return o
}
