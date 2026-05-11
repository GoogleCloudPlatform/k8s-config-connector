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

package contexts

func init() {

	resourceContextMap["fullalloydbbackup"] = ResourceContext{
		ResourceKind: "AlloyDBBackup",
		SkipUpdate:   true,
	}

	resourceContextMap["basicalloydbbackup"] = ResourceContext{
		ResourceKind: "AlloyDBBackup",
		SkipUpdate:   true,
	}

	resourceContextMap["basicalloydbcluster"] = ResourceContext{
		ResourceKind: "AlloyDBCluster",
		SkipNoChange: true, // MBUR fields always triggers an update.
	}

	resourceContextMap["fullalloydbcluster"] = ResourceContext{
		ResourceKind: "AlloyDBCluster",
		SkipNoChange: true, // MBUR fields always triggers an update.
	}

	resourceContextMap["restorebackupalloydbcluster"] = ResourceContext{
		ResourceKind: "AlloyDBCluster",
		SkipUpdate:   true,
	}

	resourceContextMap["zonalalloydbinstance"] = ResourceContext{
		ResourceKind: "AlloyDBInstance",
		SkipUpdate:   true,
	}

	// SkipDelete and SkipDriftDetection reason justification:
	// No direct delete operation for secondary instance
	// dynamic_controller_integration_test.go:518: expected GCP client to return
	// NotFound for 'alloydbinstance-2-nwdw57rwwk37lvskfvya', instead got:
	// expected error, instead got 'nil'
	resourceContextMap["basicsecondaryalloydbinstance"] = ResourceContext{
		ResourceKind:       "AlloyDBInstance",
		SkipDriftDetection: true,
		SkipDelete:         true,
	}
}
