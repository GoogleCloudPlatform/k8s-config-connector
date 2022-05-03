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
	resourceContextMap["folderinfolder"] = ResourceContext{
		ResourceKind: "Folder",
	}
	resourceContextMap["folderinorg"] = ResourceContext{
		ResourceKind: "Folder",
	}
	resourceContextMap["projectinorg"] = ResourceContext{
		ResourceKind: "Project",
		// TestCreateNoChangeUpdateDelete/basic-projectinorg: dynamic_controller_integration_test.go:239: reconcile returned unexpected error:
		//   Update call failed: error applying desired state: error creating project project-kq4irqvgso8rud79s14j (KCC-2 kq4irqvgso8rud79s14j):
		//   googleapi: Error 409: Requested entity already exists, alreadyExists. If you received a 403 error, make sure you have the
		//   `roles/resourcemanager.projectCreator` permission
		SkipDriftDetection: true,
	}
	resourceContextMap["projectinfolder"] = ResourceContext{
		ResourceKind: "Project",
		// TestCreateNoChangeUpdateDelete/basic-projectinfolder: dynamic_controller_integration_test.go:239: reconcile returned unexpected error:
		//   Update call failed: error applying desired state: error creating project project-9cv84vufffhqgr8hlscf (KCC-2 9cv84vufffhqgr8hlscf):
		//   googleapi: Error 409: Requested entity already exists, alreadyExists. If you received a 403 error, make sure you have the
		//   `roles/resourcemanager.projectCreator` permission
		SkipDriftDetection: true,
	}
	resourceContextMap["resourcemanagerlien"] = ResourceContext{
		ResourceKind: "ResourceManagerLien",
		SkipUpdate:   true,
	}
}
