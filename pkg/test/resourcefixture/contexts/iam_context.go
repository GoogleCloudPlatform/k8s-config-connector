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
	resourceContextMap["iamserviceaccountkey"] = ResourceContext{
		ResourceKind: "IAMServiceAccountKey",
		SkipUpdate:   true,
	}
	resourceContextMap["projectrole"] = ResourceContext{
		ResourceKind: "IAMCustomRole",
		// TestCreateNoChangeUpdateDelete/basic-iamcustomrole: dynamic_controller_integration_test.go:282:
		//   reconcile returned unexpected error: Delete call failed: error deleting resource: Error deleting
		//  the custom project role Example Custom Role: googleapi: Error 400: You can't delete role_id
		// (projects/cnrm-test-tgooo56g38yqbn3k/roles/roleq1s4vklq6gpalcvrachc) as it is already deleted., failedPrecondition
		SkipDriftDetection: true,
	}
}
