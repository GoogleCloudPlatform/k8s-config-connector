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

package krmtotf

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ResolveLegacyGCPManagedFields(r *Resource, liveState *terraform.InstanceState, config map[string]interface{}) error {
	// TODO(kcc-eng): This is a temporary workaround for the well-known cases of
	//  autoscaling/auto-updating fields. When the full GCP-managed fields story
	//  is in place, this will be made fully generic.
	if liveState.Empty() {
		// If the resource is being created, then there is nothing from GCP to
		// manage yet. Use the values explicitly from the customer.
		return nil
	}
	switch r.GroupVersionKind().Kind {
	case "SQLInstance":
		return resolveSQLInstanceDiskSize(r, config)
	case "ContainerCluster":
		if err := resolveContainerClusterNodeVersion(r, config); err != nil {
			return err
		}
		if err := resolveContainerClusterNodeConfig(r, liveState, config); err != nil {
			return err
		}
		return nil
	case "ContainerNodePool":
		if err := resolveContainerNodePoolVersion(r, config); err != nil {
			return err
		}
		if err := resolveContainerNodePoolInitialNodeCount(r, config); err != nil {
			return err
		}
		if err := resolveContainerNodePoolNodeCount(r, config); err != nil {
			return err
		}
		return nil
	case "ComputeBackendService":
		return resolveComputeBackendServiceBackend(r, config)
	case "BigtableInstance":
		return resolveBigtableInstanceNumNodes(r, config)
	default:
		return nil
	}
}

func isGCPManagedField(kind, field string) bool {
	// TODO(kcc-eng): This is a temporary workaround for the well-known cases of
	//  autoscaling/auto-updating fields. When the full GCP-managed fields story
	//  is in place, this will be made fully generic.
	switch kind {
	case "SQLInstance":
		return field == "settings.disk_size"
	case "ContainerCluster":
		return field == "node_version"
	case "ContainerNodePool":
		switch field {
		case "version", "node_count", "initial_node_count":
			return true
		default:
			return false
		}
	case "BigQueryConnectionConnection":
		return field == "cloud_resource.service_account_id"
	}
	return false
}

func resolveSQLInstanceDiskSize(r *Resource, config map[string]interface{}) error {
	// Customers can opt-in to automatic disk resizes. Customers can set this
	// manually, but GCP will also automatically update the field if the
	// capacity reaches a certain threshold. For more information, see:
	// https://cloud.google.com/sql/docs/mysql/instance-settings
	autoresizeEnabled, found, err := unstructured.NestedBool(config, "settings", "diskAutoresize")
	if err != nil {
		return fmt.Errorf("error determining if disk autoresize is set: %w", err)
	}
	if !found || !autoresizeEnabled {
		// Autoresize is disabled, so no special behavior required.
		return nil
	}
	if err := removeFromConfigIfNotApplied(r, config, "settings", "diskSize"); err != nil {
		return fmt.Errorf("error resolving disk size in config: %w", err)
	}
	return nil
}

func resolveContainerClusterNodeVersion(r *Resource, config map[string]interface{}) error {
	// If the customer sets a release channel on their cluster, then GKE assumes ownership
	// of the node version and will automatically revert any changes.
	releaseChannel, found, err := unstructured.NestedMap(config, "releaseChannel")
	if err != nil {
		return fmt.Errorf("error determining if release channel is set: %w", err)
	}
	if !found || releaseChannel == nil {
		// Release channel is not specified, so no special behavior required.
		return nil
	}
	if err := removeFromConfigIfNotApplied(r, config, "nodeVersion"); err != nil {
		return fmt.Errorf("error resolving node version in config: %w", err)
	}
	return nil
}

func resolveContainerNodePoolVersion(r *Resource, config map[string]interface{}) error {
	autoUpgrade, found, err := unstructured.NestedBool(config, "management", "autoUpgrade")
	if err != nil {
		return fmt.Errorf("error determining if autoupgrade is set: %w", err)
	}
	if !found || !autoUpgrade {
		// Autoupgrade is disabled, so no special behavior required.
		return nil
	}
	field := "version"
	if err := removeFromConfigIfNotApplied(r, config, field); err != nil {
		return fmt.Errorf("error resolving field '%v' in config: %w", field, err)
	}
	return nil
}

// Remove `nodeConfig` from the desired config when `remove-default-node-pool`
// directive is set to `true` and the liveState doesn't contain this field.
//
// When `remove-default-node-pool` directive is set to `true`, the default node
// pool will be removed, and `spec.nodeConfig` field should be managed by the API.
// However, because the service-generated value of `spec.nodeConfig` contains
// lists, which are preserved by KCC even if the live state of the GCP resource
// no longer has `nodeConfig` field, it triggers unexpected recreation of the
// resource. So in this case, we need to manually clean up `nodeConfig` field.
func resolveContainerClusterNodeConfig(r *Resource, liveState *terraform.InstanceState, config map[string]interface{}) error {
	removeDefaultNodePoolDirective := "remove-default-node-pool"
	nodeConfigFieldInTFState := "node_config"
	nodeConfigFieldInKRMConfig := text.SnakeCaseToLowerCamelCase(nodeConfigFieldInTFState)

	key := k8s.FormatAnnotation(removeDefaultNodePoolDirective)
	val, ok := k8s.GetAnnotation(key, r)
	if !ok || val != "true" {
		return nil
	}

	liveStateMap := InstanceStateToMap(r.TFResource, liveState)
	exists, err := topLevelObjectFieldExistsInStateMap(liveStateMap, nodeConfigFieldInTFState)
	if err != nil {
		return fmt.Errorf("error resolving field '%v' in 'ContainerCluster': %w", nodeConfigFieldInKRMConfig, err)
	}
	if exists {
		return nil
	}

	if err := removeFromConfigIfNotApplied(r, config, nodeConfigFieldInKRMConfig); err != nil {
		return fmt.Errorf("error removing field '%v' in config: %w", nodeConfigFieldInKRMConfig, err)
	}
	return nil
}

func topLevelObjectFieldExistsInStateMap(state map[string]interface{}, field string) (bool, error) {
	value, ok := state[field]
	if !ok {
		return false, nil
	}
	listVal, ok := value.([]interface{})
	if !ok {
		return false, fmt.Errorf("field '%v' is not an object field", field)
	}
	// An object field should be considered nonexistent if no sub-field is specified.
	if len(listVal) == 0 {
		return false, nil
	}
	// The response returned by terraform may insert a list of size 1 for nested fields.
	return listVal[0] != nil, nil
}

func resolveContainerNodePoolInitialNodeCount(r *Resource, config map[string]interface{}) error {
	// After create, `initialNodeCount` floats with the last value passed to the
	// `setSize` custom verb. For KCC, we will push users to use `spec.nodeCount`
	// instead, and treat initialNodeCount as GCP-managed.
	if err := removeFromConfigIfNotApplied(r, config, "initialNodeCount"); err != nil {
		return fmt.Errorf("error resolving initialNodeCount in config: %w", err)
	}
	return nil
}

func resolveContainerNodePoolNodeCount(r *Resource, config map[string]interface{}) error {
	// The `spec.nodeCount` field should be assumed to be GCP-managed when autoscaling
	// is enabled. This is determined by the presence of the "autoscaling" field.
	if val := config["autoscaling"]; val == nil {
		// Autoscaling is not enabled; so no special behavior required.
		return nil
	}
	// Autoscaling is enabled. Treat spec.nodeCount as GCP-managed.
	if err := removeFromConfigIfNotApplied(r, config, "nodeCount"); err != nil {
		return fmt.Errorf("error resolving nodeCount in config: %w", err)
	}
	return nil
}

func resolveComputeBackendServiceBackend(r *Resource, config map[string]interface{}) error {
	// If the customer omits backend definitions in their backendservice,
	// assume backend field is owned by another process
	if err := removeFromConfigIfNotApplied(r, config, "backend"); err != nil {
		return fmt.Errorf("error resolving backend in config: %w", err)
	}
	return nil
}

func resolveBigtableInstanceNumNodes(r *Resource, config map[string]interface{}) error {
	// If numNodes is not in a cluster's last applied configuration
	// remove from config
	applied, found, err := getLastAppliedValue(r, "cluster")
	if err != nil {
		return fmt.Errorf("error determining last applied clusters: %w", err)
	}
	if !found {
		return nil
	}
	appliedClusters, ok := applied.([]interface{})
	if !ok {
		return fmt.Errorf("cannot decode last applied clusters")
	}
	for _, c := range appliedClusters {
		c, ok := c.(map[string]interface{})
		if !ok {
			return fmt.Errorf("cannot decode cluster")
		}
		clusterID, found, err := unstructured.NestedString(c, "clusterId")
		if err != nil {
			return fmt.Errorf("error determining clusterId: %w", err)
		} else if !found {
			return fmt.Errorf("cannot determine clusterId")
		}
		_, found, err = unstructured.NestedFloat64(c, "numNodes")
		if err != nil {
			return fmt.Errorf("error determining numNodes: %w", err)
		}
		if !found {
			// remove from output config
			if err = removeNumNodesFromBigtableCluster(config, clusterID); err != nil {
				return fmt.Errorf("error removing numNodes: %w", err)
			}
		}
	}
	return nil
}

// removeNumNodesFromBigtableCluster removes the numNodes field from the named
// cluster specification if it exists
func removeNumNodesFromBigtableCluster(config map[string]interface{}, cluster string) error {
	// Fetch clusters from config
	clusters, found, err := unstructured.NestedSlice(config, "cluster")
	if err != nil {
		return fmt.Errorf("error finding clusters in config: %w", err)
	}
	if !found {
		return nil
	}
	for _, c := range clusters {
		c, ok := c.(map[string]interface{})
		if !ok {
			return fmt.Errorf("cannot decode cluster")
		}
		id, found, _ := unstructured.NestedString(c, "clusterId")
		if !found {
			return fmt.Errorf("cannot determine cluster id")
		}
		if id == cluster {
			unstructured.RemoveNestedField(c, "numNodes")
		}
	}
	if err := unstructured.SetNestedSlice(config, clusters, "cluster"); err != nil {
		return fmt.Errorf("error setting cluster list: %w", err)
	}
	return nil
}

func getLastAppliedValue(r *Resource, path ...string) (val interface{}, found bool, err error) {
	// Note: Only values from kubectl's last applied configuration will be recognized. Values
	// set manually with server-side apply or edited via POST/PUT/PATCH directly will not be
	// recognized. kubectl's annotation is present on all `kubectl apply`-ed resources for
	// GKE versions 1.14-1.16.
	lastAppliedConfigRaw, ok := k8s.GetAnnotation(k8s.LastAppliedConfigurationAnnotation, r)
	if !ok {
		return nil, false, nil
	}
	lastAppliedConfig := make(map[string]interface{})
	if err := json.Unmarshal([]byte(lastAppliedConfigRaw), &lastAppliedConfig); err != nil {
		return nil, false, fmt.Errorf("error unmarshalling last applied configuration: %w", err)
	}
	specPath := append([]string{"spec"}, path...)
	return unstructured.NestedFieldCopy(lastAppliedConfig, specPath...)
}

// removeFromConfigIfNotApplied removes the specified field from the config, unless it was
// applied explicitly by the customer.
func removeFromConfigIfNotApplied(r *Resource, config map[string]interface{}, path ...string) error {
	// Note that this does not perform any mappings. It is assumed the path and config are
	// in the KRM camelCase format.
	_, found, err := getLastAppliedValue(r, path...)
	if err != nil {
		return fmt.Errorf("error finding last applied value for disk size: %w", err)
	}
	if !found {
		// The value was not found in the last applied configuration. Delegate
		// instead to GCP by removing the field from the config. The live state's
		// value will be substituted during the diff calculation.
		unstructured.RemoveNestedField(config, path...)
	}
	return nil
}
