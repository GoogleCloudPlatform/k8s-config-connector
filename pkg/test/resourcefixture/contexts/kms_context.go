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
	resourceContextMap["kmskeyring"] = ResourceContext{
		ResourceKind: "KMSKeyRing",
		SkipUpdate:   true,
	}
	resourceContextMap["kmscryptokey"] = ResourceContext{
		ResourceKind: "KMSCryptoKey",
		// TestCreateNoChangeUpdateDelete/basic-kmscryptokey: dynamic_controller_integration_test.go:149: value
		//   mismatch for label with key 'key-one': got 'value-two', want 'value-one'
		// TestCreateNoChangeUpdateDelete/basic-kmscryptokey: dynamic_controller_integration_test.go:282: reconcile
		//   returned unexpected error: Delete call failed: error deleting resource: googleapi: Error 400: The request
		//   cannot be fulfilled. Resource projects/cnrm-test-tgooo56g38yqbn3k/locations/us-central1/keyRings/kmscryptokey-i94182u4ku0un1f0nsna/cryptoKeys/kmscryptokey-i94182u4ku0un1f0nsna/cryptoKeyVersions/1
		//   has value DESTROY_SCHEDULED in field crypto_key_version.state., failedPrecondition
		SkipDriftDetection: true,
	}
	resourceContextMap["kmsautokeyconfig"] = ResourceContext{
		ResourceKind: "KMSAutokeyConfig",
		// The AutokeyConfig resource does not support delete operation.
		SkipDriftDetection: true,
		SkipDelete:         true,
	}
	resourceContextMap["kmskeyhandle"] = ResourceContext{
		ResourceKind: "KMSKeyHandle",
		// The KMSKeyHandle resource does not support update and delete operation.
		SkipDriftDetection: true,
		SkipUpdate:         true,
		SkipDelete:         true,
	}
	resourceContextMap["basickmsimportjob"] = ResourceContext{
		ResourceKind: "KMSImportJob",
		// The KMSImportJob resource does not support update and delete operation.
		SkipDriftDetection: true,
		SkipUpdate:         true,
		SkipDelete:         true,
	}
}
