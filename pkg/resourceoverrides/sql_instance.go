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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetSQLInstanceResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "SQLInstance",
	}
	// Keep the 'databaseVersion' field optional with default for backwards compatibility.
	// See b/206145549 for context.
	ro.Overrides = append(ro.Overrides, keepDatabaseVersionFieldOptionalWithDefault())
	// Preserve the legacy fields for first generation sql instances that have been removed from Terraform 4.x upgrade.
	// See b/206145549 for context.
	ro.Overrides = append(ro.Overrides, keepFirstGenerationFields())

	ro.Overrides = append(ro.Overrides, copyInstanceTypeFieldToStatus())
	return ro
}

func copyInstanceTypeFieldToStatus() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		spec := schema.Properties["spec"]
		status := schema.Properties["status"]
		status.Properties["instanceType"] = spec.Properties["instanceType"]

		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		if tfState != nil {
			reconciled.Status["instanceType"] = tfState.Attributes["instance_type"]
		}

		return nil
	}
	return o
}

func keepDatabaseVersionFieldOptionalWithDefault() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return KeepTopLevelFieldOptionalWithDefault(crd, "MYSQL_5_6", "databaseVersion")
	}
	return o
}

func keepFirstGenerationFields() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		settings := schema.Properties["spec"].Properties["settings"]
		settings.Properties["authorizedGaeApplications"] = apiextensions.JSONSchemaProps{
			Description: "DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.\n" +
				"Specifying this field has no-ops; it's recommended to remove this field from your configuration.",
			Type: "array",
			Items: &apiextensions.JSONSchemaPropsOrArray{
				Schema: &apiextensions.JSONSchemaProps{
					Type: "string",
				},
			},
		}
		settings.Properties["crashSafeReplication"] = apiextensions.JSONSchemaProps{
			Description: "DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.\n" +
				"Specifying this field has no-ops; it's recommended to remove this field from your configuration.",
			Type: "boolean",
		}
		settings.Properties["replicationType"] = apiextensions.JSONSchemaProps{
			Description: "DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.\n" +
				"Specifying this field has no-ops; it's recommended to remove this field from your configuration.",
			Type: "string",
		}
		return nil
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := PruneNoOpsField(r, "settings", "authorizedGaeApplications"); err != nil {
			return fmt.Errorf("error pruning no-ops field 'settings.authorizedGaeApplications' in pre-actuation transformation: %w", err)
		}
		if err := PruneNoOpsField(r, "settings", "crashSafeReplication"); err != nil {
			return fmt.Errorf("error pruning no-ops field 'settings.crashSafeReplication' in pre-actuation transformation: %w", err)
		}
		if err := PruneNoOpsField(r, "settings", "replicationType"); err != nil {
			return fmt.Errorf("error pruning no-ops field 'settings.replicationType' in pre-actuation transformation: %w", err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		if err := PreserveUserSpecifiedLegacyField(original, reconciled, "settings", "authorizedGaeApplications"); err != nil {
			return fmt.Errorf("error preserving no-ops field 'settings.authorizedGaeApplications' in post-actuation transformation: %w", err)
		}
		if err := PreserveUserSpecifiedLegacyField(original, reconciled, "settings", "crashSafeReplication"); err != nil {
			return fmt.Errorf("error preserving no-ops field 'settings.crashSafeReplication' in post-actuation transformation: %w", err)
		}
		if err := PreserveUserSpecifiedLegacyField(original, reconciled, "settings", "replicationType"); err != nil {
			return fmt.Errorf("error preserving no-ops field 'settings.replicationType' in post-actuation transformation: %w", err)
		}
		return nil
	}
	return o
}
