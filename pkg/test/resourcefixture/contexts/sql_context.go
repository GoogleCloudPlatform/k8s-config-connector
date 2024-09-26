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

import "time"

func init() {
	resourceContextMap["sqlinstance-activationpolicy"] = ResourceContext{
		// TODO: After switching to use direct controller, we can update the direct controller
		// logic to support creating SQLInstances with `activationPolicy: "NEVER"`. Then, we
		// can enable this test. The TF based controller does not support creating
		// SQLInstances with `activationPolicy: "NEVER"`.
		SkipDriftDetection: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-activationpolicy-direct"] = ResourceContext{
		// TODO: After switching to use direct controller, we can update the direct controller
		// logic to support creating SQLInstances with `activationPolicy: "NEVER"`. Then, we
		// can enable this test. The TF based controller does not support creating
		// SQLInstances with `activationPolicy: "NEVER"`.
		SkipDriftDetection: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-activedirectoryconfig"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-activedirectoryconfig-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-auditconfig"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-auditconfig-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-authorizednetworks"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-authorizednetworks-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-backupconfiguration-binarylog"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-backupconfiguration-binarylog-direct"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-backupconfiguration-pitr"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-backupconfiguration-pitr-direct"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	/* Note: Terraform-based controller does not support create from clone.
	resourceContextMap["sqlinstance-clone-minimal"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}
	*/

	resourceContextMap["sqlinstance-clone-minimal-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-connectorenforcement"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-connectorenforcement-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-databaseflags"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-databaseflags-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-datacacheconfig"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-datacacheconfig-direct"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-deletionprotection"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-deletionprotection-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-denymaintenanceperiod"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-denymaintenanceperiod-direct"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-encryptionkey"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-encryptionkey-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-insightsconfig"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-insightsconfig-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-locationpreference"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-locationpreference-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-maintenancewindow"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-maintenancewindow-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-multithreading"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-multithreading-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-mysql"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-mysql-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-mysql-minimal"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-mysql-minimal-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-passwordvalidationpolicy"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-passwordvalidationpolicy-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-postgres"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-postgres-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-postgres-minimal"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-postgres-minimal-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-privatenetwork"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-privatenetwork-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-privatenetwork-legacyref"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-privatenetwork-legacyref-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-replica"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-replica-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-sqlserver"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-sqlserver-direct"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-sqlserver-minimal"] = ResourceContext{
		// SQL instances need a bit of additional time before attempting to recreate with
		// the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-sqlserver-minimal-direct"] = ResourceContext{
		// SQL instances need a bit of additional time before attempting to recreate with
		// the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-ssl"] = ResourceContext{
		// SQL instances need a bit of additional time before attempting to recreate with
		// the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-ssl-direct"] = ResourceContext{
		// SQL instances need a bit of additional time before attempting to recreate with
		// the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-storage"] = ResourceContext{
		// SQL instances need a bit of additional time before attempting to recreate with
		// the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-storage-direct"] = ResourceContext{
		// SQL instances need a bit of additional time before attempting to recreate with
		// the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqldatabase"] = ResourceContext{
		ResourceKind: "SQLDatabase",
	}

	resourceContextMap["sqlsslcert"] = ResourceContext{
		ResourceKind: "SQLSSLCert",
		SkipUpdate:   true,
	}

	resourceContextMap["sqluser"] = ResourceContext{
		ResourceKind: "SQLUser",
		// Skip update test, as the only field that is updatable is "password", which we
		// can't verify on the underlying API due to it not being passed back in a GET.
		SkipUpdate: true,
	}
}
