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

package sampleconversion

import (
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"
)

type DependencyGraph struct {
	// edges is the map of directed edges from the vertex (key of the map) to
	// the list of adjacent vertices (value of the map) in the graph.
	edges map[string][]string
}

func NewDependencyGraph() *DependencyGraph {
	edges := make(map[string][]string)
	return &DependencyGraph{edges: edges}
}

// AddDependencyWithTFRefVal adds an edge in the dependency graph using the
// referencer's TF type, and the referencee's "TF Reference Value" in the format
// of `[referenced_tf_type].[referenced_resource_name].[referenced_field_name]`.
func (d *DependencyGraph) AddDependencyWithTFRefVal(tfRefValOfReferencee, referencer string) {
	referencee, _ := extractReferencedTFTypeAndField(tfRefValOfReferencee)
	d.addDependency(referencee, referencer)
}

// addDependency adds an edge in the dependency graph using the referencer's TF
// type, and the referencee's TF type.
func (d *DependencyGraph) addDependency(referencee, referencer string) {
	referencerList := make([]string, 0)
	if _, ok := d.edges[referencee]; ok {
		referencerList = d.edges[referencee]
	}
	referencerList = append(referencerList, referencer)
	d.edges[referencee] = referencerList
}

func (d *DependencyGraph) TopologicalSort() []string {
	visited := make(map[string]bool)
	referencees := make([]string, 0)
	result := make([]string, 0)

	for referencee := range d.edges {
		visited[referencee] = false
		referencees = append(referencees, referencee)
	}
	sort.Strings(referencees)

	// Iterate the referencees following the alphabetical order.
	for _, referencee := range referencees {
		result = topologicalSort(referencee, d.edges, visited, result)
	}

	slice.Reverse(result)
	return result
}

func topologicalSort(vertex string, edges map[string][]string, visited map[string]bool, result []string) (updatedResult []string) {
	// If the vertex has been visited, or if the vertex doesn't have any
	// outgoing edges (thus not covered by the "visited" map), skip the
	// topological sorting for the vertex.
	if isVisited, ok := visited[vertex]; isVisited || !ok {
		return result
	}
	visited[vertex] = true
	adjacentList, ok := edges[vertex]
	if !ok || len(adjacentList) == 0 {
		result = append(result, vertex)
		return result
	}
	for _, adjacentVertex := range adjacentList {
		result = topologicalSort(adjacentVertex, edges, visited, result)
	}
	result = append(result, vertex)
	return result
}
