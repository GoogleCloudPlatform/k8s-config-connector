// Copyright 2023 Google LLC
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

func GetRedisInstanceResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "RedisInstance",
	}
	// Preserve the legacy fields for maintenanceSchedule that has been removed from Terraform 4.74.0 upgrade.
	// See b/293919820 and https://github.com/hashicorp/terraform-provider-google-beta/releases/tag/v4.74.0 for context.
	ro.Overrides = append(ro.Overrides, copyMaintenanceScheduleFieldToSpec())
	return ro
}

func copyMaintenanceScheduleFieldToSpec() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		spec := schema.Properties["spec"]
		status := schema.Properties["status"]
		spec.Properties["maintenanceSchedule"] = status.Properties["maintenanceSchedule"]
		return nil
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := PruneNoOpsField(r, "maintenanceSchedule"); err != nil {
			return fmt.Errorf("error pruning no-ops field 'maintenanceSchedule' in pre-actuation transformation: %w", err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		if reconciled.Status["maintenanceSchedule"] != nil {
			reconciled.Spec["maintenanceSchedule"] = reconciled.Status["maintenanceSchedule"]
		}
		return nil
	}
	return o
}
