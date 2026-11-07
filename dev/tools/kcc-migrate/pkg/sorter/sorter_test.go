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

package sorter

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSortResources(t *testing.T) {
	yamlInput := `
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: app-member
---
apiVersion: kms.cnrm.cloud.google.com/v1beta1
kind: KMSKeyRing
metadata:
  name: app-ring
---
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: app-sql
---
apiVersion: kms.cnrm.cloud.google.com/v1beta1
kind: KMSCryptoKey
metadata:
  name: app-key
---
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLDatabase
metadata:
  name: app-db
`
	items, err := ParseResourcesFromData([]byte(yamlInput))
	if err != nil {
		t.Fatalf("ParseResourcesFromData failed: %v", err)
	}

	if len(items) != 5 {
		t.Fatalf("Expected 5 items, got %d", len(items))
	}

	sorted := SortResources(items)

	expectedOrder := []string{"KMSKeyRing", "KMSCryptoKey", "SQLInstance", "SQLDatabase", "IAMPolicyMember"}
	for i, exp := range expectedOrder {
		if sorted[i].Kind != exp {
			t.Errorf("At index %d expected %s, got %s", i, exp, sorted[i].Kind)
		}
	}
}

func TestSortDirectory(t *testing.T) {
	tmpSrc := t.TempDir()
	tmpDst := t.TempDir()

	manifest := `
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: member-1
---
apiVersion: kms.cnrm.cloud.google.com/v1beta1
kind: KMSKeyRing
metadata:
  name: ring-1
`
	if err := os.WriteFile(filepath.Join(tmpSrc, "manifest.yaml"), []byte(manifest), 0644); err != nil {
		t.Fatal(err)
	}

	count, err := SortDirectory(tmpSrc, tmpDst)
	if err != nil {
		t.Fatalf("SortDirectory failed: %v", err)
	}
	if count != 2 {
		t.Fatalf("Expected 2 resources, got %d", count)
	}

	tier0 := filepath.Join(tmpDst, "tier0_ordered.yaml")
	if _, err := os.Stat(tier0); os.IsNotExist(err) {
		t.Fatalf("Expected tier0 file to exist")
	}

	tier5 := filepath.Join(tmpDst, "tier5_ordered.yaml")
	if _, err := os.Stat(tier5); os.IsNotExist(err) {
		t.Fatalf("Expected tier5 file to exist")
	}
}
