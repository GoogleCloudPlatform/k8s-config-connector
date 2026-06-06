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

package e2e

import (
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
)

func TestListComputeFixtures(t *testing.T) {
	lightFilter := func(name string, testType resourcefixture.TestType) bool {
		return !strings.Contains(name, "iam-bigqueryconnectionconnectionref") &&
			!strings.Contains(name, "iam-logsinkref") &&
			!strings.Contains(name, "iam-serviceaccountref") &&
			!strings.Contains(name, "iam-serviceidentityref") &&
			!strings.Contains(name, "iam-sqlinstanceref")
	}
	pathFilter := func(path string) bool {
		return !strings.Contains(path, "testdata/iam/iampartialpolicy")
	}

	fixtures := resourcefixture.LoadWithPathFilter(t, pathFilter, lightFilter, nil)
	for _, fixture := range fixtures {
		if fixture.GVK.Group == "compute.cnrm.cloud.google.com" {
			fmt.Printf("FIXTURE_COMPUTE_TEST: %s\n", fixture.TestKey)
		}
	}
}
