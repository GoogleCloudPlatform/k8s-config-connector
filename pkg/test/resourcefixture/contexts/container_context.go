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

package contexts

func init() {
	resourceContextMap["containercluster"] = ResourceContext{
		ResourceKind: "ContainerCluster",
		// TestCreateNoChangeUpdateDelete/basic-containercluster: dynamic_controller_integration_test.go:239:
		//    reconcile returned unexpected error: Update call failed: error applying desired state: node_version and
		//    min_master_version must be set to equivalent values on create
		SkipDriftDetection: true,
	}
	resourceContextMap["removedefaultnodepool"] = ResourceContext{
		ResourceKind: "ContainerCluster",
		// TestCreateNoChangeUpdateDelete/directives-removedefaultnodepool: dynamic_controller_integration_test.go:239:
		//   reconcile returned unexpected error: Update call failed: error applying desired state: node_version and
		//   min_master_version must be set to equivalent values on create
		SkipDriftDetection: true,
	}
	resourceContextMap["containernodepool"] = ResourceContext{
		ResourceKind: "ContainerNodePool",
		// TestCreateNoChangeUpdateDelete/basic-containernodepool: dynamic_controller_integration_test.go:239: reconcile
		//   returned unexpected error: Update call failed: error applying desired state: Cannot set both initial_node_count
		//   and node_count on node pool nodepool-sample-5bb55e2rtb33bp3d4vvnTestCreateNoChangeUpdateDelete/directives-removedefaultnodepool:
		//   dynamic_controller_integration_test.go:239:
		SkipDriftDetection: true,
	}
}
