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
	resourceContextMap["sqlinstance-authorizednetworks"] = ResourceContext{
		// SQL instances appear to need a bit of additional time before attempting to recreate
		// with the exact same name. Otherwise, the GCP API returns "unknown error".
		RecreateDelay: time.Second * 60,
		ResourceKind:  "SQLInstance",
	}

	resourceContextMap["sqlinstance-clone-minimal"] = ResourceContext{
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

	resourceContextMap["sqlinstance-locationpreference"] = ResourceContext{
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

	resourceContextMap["sqlinstance-mysql-minimal"] = ResourceContext{
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

	resourceContextMap["sqlinstance-postgres-minimal"] = ResourceContext{
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

	resourceContextMap["sqlinstance-sqlserver"] = ResourceContext{
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

	resourceContextMap["sqlinstance-ssl"] = ResourceContext{
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
