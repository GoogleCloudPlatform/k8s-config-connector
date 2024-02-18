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
	"fmt"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func withCustomFlatteners(config map[string]interface{}, kind string) (map[string]interface{}, error) {
	switch kind {
	case "ComputeInstance", "ComputeInstanceTemplate":
		return FlattenComputeInstanceMetadata(config)
	default:
		return config, nil
	}
}

func withCustomExpanders(state map[string]interface{}, prev *Resource, kind string) map[string]interface{} {
	switch kind {
	case "ComputeInstance", "ComputeInstanceTemplate":
		return ExpandComputeInstanceMetadata(state, prev)
	default:
		return state
	}
}

func withResourceCustomResolvers(config map[string]interface{}, liveState map[string]interface{}, kind string, r *tfschema.Resource) (map[string]interface{}, error) {
	switch kind {
	case "BigtableInstance":
		return MergeClusterConfigsFromLiveStateForBigtableInstance(config, liveState, r)
	default:
		return config, nil
	}
}

// MergeClusterConfigsFromLiveStateForBigtableInstance is a resource specific function to deal with the following edge case.
// BigtableInstance has a `cluster` field that takes a full list of clusters associated with the instance. The list of clusters read from the API is unordered.
// Due to the terraform SDK limitation, if some optional field e.g. num_nodes is omitted, terraform SDK will determine the current value of the field from the cluster on the same index
// rather than from the cluster with the same cluster_id; plus the returned list is not in the same order as user specified, the partial config with optional fields omitted will result in unexpected behaviors.
// As a workarounds until migrating this resource to DCL, KCC will maintain this following resource specific code to merge the cluster config for omitted fields from cluster's live state by cluster_id.
// DCL is expected to have the similar logic on its side to merge the partial desired intent with the live state; once this resource is migrated to DCL, we should be able to remove the bespoke code.
func MergeClusterConfigsFromLiveStateForBigtableInstance(config map[string]interface{}, liveState map[string]interface{}, r *tfschema.Resource) (map[string]interface{}, error) {
	if len(liveState) == 0 || len(config) == 0 {
		return config, nil
	}
	clustersRaw, found, err := unstructured.NestedFieldCopy(liveState, "cluster")
	if err != nil {
		return nil, err
	}
	if !found || clustersRaw == nil {
		return config, nil
	}
	clusters, ok := clustersRaw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the value of `cluster` field in config to be []interface{}, but was actually %T", clustersRaw)
	}

	newClustersRaw, found, err := unstructured.NestedFieldCopy(config, "cluster")
	if err != nil {
		return nil, err
	}
	if !found || newClustersRaw == nil {
		return config, nil
	}
	newClusters, ok := newClustersRaw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the value of `cluster` field in state to be []interface{}, but was actually %T", newClusters)
	}
	for _, item := range newClusters {
		newCluster, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the item value of `cluster` field to be map[string]interface{}, but was actually %T", item)
		}
		c, found, err := getClusterByID(clusters, newCluster["cluster_id"])
		if err != nil {
			return nil, err
		}
		if !found {
			continue
		}
		clusterSchema := r.Schema["cluster"].Elem.(*tfschema.Resource)
		mergeClusterConfigWithLiveState(newCluster, c, clusterSchema)
	}
	config["cluster"] = newClusters
	return config, nil
}

func getClusterByID(clusters []interface{}, id interface{}) (map[string]interface{}, bool, error) {
	for _, item := range clusters {
		c, ok := item.(map[string]interface{})
		if !ok {
			return nil, false, fmt.Errorf("expected the item value of `cluster` field to be map[string]interface{}, but was actually %T", item)
		}
		if id == c["cluster_id"] {
			return c, true, nil
		}
	}
	return nil, false, nil
}

func mergeClusterConfigWithLiveState(config map[string]interface{}, state map[string]interface{}, clusterSchema *tfschema.Resource) {
	for f := range clusterSchema.Schema {
		if f == "cluster_id" {
			continue
		}

		if _, ok := config[f]; !ok {
			if v, ok := state[f]; ok {
				config[f] = v
			}
		}
	}
}

func FlattenComputeInstanceMetadata(config map[string]interface{}) (map[string]interface{}, error) {
	metadataRaw, _ := config["metadata"]
	if metadataRaw == nil {
		return config, nil
	}
	metadataList, ok := metadataRaw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("metadata field not in expected list format")
	}
	config = deepcopy.MapStringInterface(config)
	metadataMap, err := convertStructuredListToMap(metadataList)
	if err != nil {
		return nil, fmt.Errorf("error converting structured list to map: %w", err)
	}
	config["metadata"] = metadataMap
	return config, nil
}

func ExpandComputeInstanceMetadata(state map[string]interface{}, prev *Resource) map[string]interface{} {
	// If previous resource already has metadata set, copy the metadata from
	// the previous resource onto the state to preserve the desired ordering
	// from the customer.
	prevMetadata, _ := prev.Spec["metadata"]
	if prevMetadata != nil {
		state = deepcopy.MapStringInterface(state)
		state["metadata"] = prevMetadata
		return state
	}
	// Otherwise, if the previous resource does not specify metadata, then
	// convert the metadata map in the state to a structured list.
	stateMetadataMapRaw, ok := state["metadata"]
	if !ok || stateMetadataMapRaw == nil {
		return state
	}
	stateMetadataMap, ok := stateMetadataMapRaw.(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("metadata in state unexpectedly is not of type map[string]interface{}: %v", stateMetadataMapRaw))
	}
	state["metadata"] = convertMapToStructuredList(stateMetadataMap)
	return state
}

func convertStructuredListToMap(l []interface{}) (map[string]interface{}, error) {
	// Converts a list of structured {"key":"foo", "value":"bar"} objects into a simple
	// map[string]interface{}.
	m := make(map[string]interface{})
	for _, elemRaw := range l {
		elem, ok := elemRaw.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("could not process metadata list element as object")
		}
		key, ok := elem["key"].(string)
		if !ok {
			return nil, fmt.Errorf("metadata list element's key is not string: %+v", key)
		}
		val := elem["value"]
		if existingVal, ok := m[key]; ok {
			return nil, fmt.Errorf("duplicate values set for key '%v': [%v, %v]", key, existingVal, val)
		}
		m[key] = val
	}
	return m, nil
}

func convertMapToStructuredList(m map[string]interface{}) []interface{} {
	// Converts a map[string]interface{} to a list of structured {"key":"foo",
	// "value":"bar"} objects. Resulting list is sorted by key to ensure
	// function is deterministic.
	l := make([]interface{}, 0)
	for k, v := range m {
		l = append(l, map[string]interface{}{
			"key":   k,
			"value": v,
		})
	}
	sort.Slice(l, func(i, j int) bool {
		iKey, err := keyForStructuredListElement(l[i])
		if err != nil {
			panic(err)
		}
		jKey, err := keyForStructuredListElement(l[j])
		if err != nil {
			panic(err)
		}
		return iKey < jKey
	})
	return l
}

func keyForStructuredListElement(rawElem interface{}) (string, error) {
	elem, ok := rawElem.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("structured list element unexpectedly does not have type map[string]interface{}: %v", rawElem)
	}
	rawKey, ok := elem["key"]
	if !ok {
		return "", fmt.Errorf("structured list element unexpectedly does not have a key: %v", rawElem)
	}
	key, ok := rawKey.(string)
	if !ok {
		return "", fmt.Errorf("structured list element's key unexpectedly does not have type string: %v", rawKey)
	}
	return key, nil
}
