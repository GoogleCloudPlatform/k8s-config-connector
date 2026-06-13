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

package objects

import (
	"strings"
	"testing"
)

func TestParseObjectsFromStream(t *testing.T) {
	manifest := `
apiVersion: dns.cnrm.cloud.google.com/v1beta1
kind: DNSManagedZone
metadata:
  name: test-zone
  namespace: test-namespace
spec:
  dnsName: "test.example.com."
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: default
  namespace: test-namespace
`

	objs, err := ParseObjectsFromStream(strings.NewReader(manifest))
	if err != nil {
		t.Fatalf("ParseObjectsFromStream failed: %v", err)
	}

	if len(objs) != 2 {
		t.Fatalf("expected 2 objects, got %d", len(objs))
	}

	if objs[0].GetKind() != "DNSManagedZone" || objs[0].GetName() != "test-zone" {
		t.Errorf("unexpected first object: %v", objs[0])
	}

	if objs[1].GetKind() != "ComputeNetwork" || objs[1].GetName() != "default" {
		t.Errorf("unexpected second object: %v", objs[1])
	}
}
