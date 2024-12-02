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
	resourceContextMap["sqlinstance-activationpolicy"] = ResourceContext{
		// TODO: After switching to use direct controller, we can update the direct controller
		// logic to support creating SQLInstances with `activationPolicy: "NEVER"`. Then, we
		// can enable this test. The TF based controller does not support creating
		// SQLInstances with `activationPolicy: "NEVER"`.
		SkipDriftDetection: true,
		ResourceKind:       "SQLInstance",
	}

	resourceContextMap["sqlinstance-activationpolicy-direct"] = ResourceContext{
		// TODO: After switching to use direct controller, we can update the direct controller
		// logic to support creating SQLInstances with `activationPolicy: "NEVER"`. Then, we
		// can enable this test. The TF based controller does not support creating
		// SQLInstances with `activationPolicy: "NEVER"`.
		SkipDriftDetection: true,
		ResourceKind:       "SQLInstance",
	}

	resourceContextMap["sqlinstance-backupconfiguration-binarylog"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		ResourceKind: "SQLInstance",
	}

	resourceContextMap["sqlinstance-backupconfiguration-binarylog-direct"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		ResourceKind: "SQLInstance",
	}

	resourceContextMap["sqlinstance-backupconfiguration-pitr"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		ResourceKind: "SQLInstance",
	}

	resourceContextMap["sqlinstance-backupconfiguration-pitr-direct"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		ResourceKind: "SQLInstance",
	}

	resourceContextMap["sqlinstance-datacacheconfig"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		ResourceKind: "SQLInstance",
	}

	resourceContextMap["sqlinstance-datacacheconfig-direct"] = ResourceContext{
		// TODO: Remove after switching to use direct controller.
		SkipNoChange: true,
		ResourceKind: "SQLInstance",
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
