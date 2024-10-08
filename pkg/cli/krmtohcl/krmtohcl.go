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

package krmtohcl

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/serialization"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func UnstructuredToHCL(ctx context.Context, u *unstructured.Unstructured, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider) (string, error) {
	gvk := u.GroupVersionKind()
	sm, err := smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return "", err
	}
	krmResource, err := krmtotf.NewResource(u, sm, tfProvider)
	if err != nil {
		return "", fmt.Errorf("could not parse resource %s: %w", u.GetName(), err)
	}
	config, _, err := krmtotf.KRMResourceToTFResourceConfigFull(krmResource, k8s.NewErroringClient(), smLoader, nil, nil, true)
	if err != nil {
		return "", fmt.Errorf("error expanding resource configuration: %w", err)
	}
	configAsMap := krmtotf.ResourceConfigToMap(config)
	tfResource := krmResource.TFResource
	removeConflictingFields(u.GetObjectKind().GroupVersionKind().String(), configAsMap)
	state := terraform.NewInstanceStateShimmedFromValue(krmtotf.MapToCtyVal(configAsMap, tfResource.CoreConfigSchema().ImpliedType()), tfResource.SchemaVersion)
	// TODO: it has not been 100% verified what the canonical purpose of krmResource.TFInfo.ID is.
	// this assignment is being done due to the bespoke InstanceState converter that was ported over
	// from https://critique-ng.corp.google.com/cl/335955454 uses TFInfo.Id, and it's best not to
	// fork this copied code to enable drop-in replacement.
	//
	// The bespoke InstanceState code considers TFInfo.Id to be the resource name.
	krmResource.TFInfo.Id = krmtotf.KRMNameToTerraformID(krmResource.GetObjectMeta().GetName())

	exportOp := &operations.TerraformExport{
		TerraformState: state,
		TerraformInfo:  krmResource.TFInfo,
	}

	if err := resourceoverrides.Handler.PreTerraformExport(ctx, gvk, exportOp); err != nil {
		return "", err
	}

	hcl, err := serialization.InstanceStateToHCL(exportOp.TerraformState, exportOp.TerraformInfo, tfProvider)
	if err != nil {
		return "", fmt.Errorf("error generating hcl: %w", err)
	}

	importID, err := krmResource.GetImportID(k8s.NewErroringClient(), smLoader)
	if err != nil {
		return "", fmt.Errorf("error getting import id for '%v': %w", krmResource.GetName(), err)
	}
	// append a comment with terraform import command for two reasons:
	// 1. A human reading the output could use the value to perform an import
	// 2. gcloud is looking for this output and printing it out for their users
	//
	// any changes to the format of this output should be communicated to the gcloud team
	hcl = fmt.Sprintf("%v# terraform import %v.%v %v\n", hcl, exportOp.TerraformInfo.Type, krmResource.TFInfo.Id, importID)
	return hcl, nil
}

// removingConflictingFields removes values that conflict with each other
// as indicated by the Terraform Resource's ConflictsWith array.
//
// Currently, KCC behavior allows for conflicting Terraform fields to exist
// in a valid KCC object, and this method is used to clean up the output for
// config-connector CLI's terraform output only.
//
// Although Terraform schema denotes the conflicting fields, because removing fields can
// lead to unintended consequences, these cases are hard-coded to vet the right output.
func removeConflictingFields(groupVersionKind string, m map[string]interface{}) {
	switch groupVersionKind {
	case "storage.cnrm.cloud.google.com/v1beta1, Kind=StorageBucket":
		// conflicts with uniformBucketLevelAccess, with bucket_policy_only
		// being a deprecated field.
		// https://cloud.google.com/storage/docs/json_api/v1/buckets/insert
		delete(m, "bucket_policy_only")
	case "container.cnrm.cloud.google.com/v1beta1, Kind=ContainerCluster":
		// loggingService and metricService are only used to configure logging syncs,
		// which is valid only pre GKE-1.15.
		// https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters#Cluster.LoggingConfig
		delete(m, "logging_service")
		delete(m, "monitoring_service")
	}
}
