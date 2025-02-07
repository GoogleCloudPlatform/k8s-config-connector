// Copyright 2025 Google LLC
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

package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func main() {
	// Generate the list of kinds that supports state-into-spec merge.
	mergeableKinds := make(map[string]bool)
	for gvk, metadata := range supportedgvks.SupportedGVKs {
		if supportedgvks.IsDirectByGVK(gvk) {
			continue
		}
		// Some alpha stability resources have reference docs so this condition
		// has some exception(s).
		if value, ok := metadata.Labels["cnrm.cloud.google.com/stability-level"]; ok && value == "alpha" && gvk.Kind != "ConfigControllerInstance" {
			continue
		}
		if k8s.SupportsStateIntoSpecMerge(gvk) {
			mergeableKinds[gvk.Kind] = true
		}
	}
	sortedKinds := slices.Sorted(maps.Keys(mergeableKinds))
	for _, k := range sortedKinds {
		fmt.Printf("*   %s\n", k)
	}
}
