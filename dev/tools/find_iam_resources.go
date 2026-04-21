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

package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	// We need these to call supportedgvks.All
	smLoader, err := servicemappingloader.New()
	if err != nil {
		panic(err)
	}
	serviceMetaLoader := metadata.New()

	gvks, err := supportedgvks.All(smLoader, serviceMetaLoader)
	if err != nil {
		panic(err)
	}

	var results []schema.GroupVersionKind
	for _, gvk := range gvks {
		if supportedgvks.SupportsIAMByGVK(gvk) {
			results = append(results, gvk)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Group != results[j].Group {
			return results[i].Group < results[j].Group
		}
		return results[i].Kind < results[j].Kind
	})

	fmt.Printf("%-50s | %-10s | %-40s\n", "Group", "Version", "Kind")
	fmt.Println(strings.Repeat("-", 105))
	for _, gvk := range results {
		fmt.Printf("%-50s | %-10s | %-40s\n", gvk.Group, gvk.Version, gvk.Kind)
	}
}
