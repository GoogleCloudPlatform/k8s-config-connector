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

package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRepoRoot(t *testing.T) {
	for _, tc := range []struct {
		baseDir string
		want    string
	}{
		// generate-types' default, from GenerateCRDOptions.InitDefaults.
		{"/repo/apis/", "/repo"},
		{"/repo/apis", "/repo"},
		{"/tmp/out", "/tmp/out"},
		{"/tmp/out/", "/tmp/out"},
	} {
		t.Run(tc.baseDir, func(t *testing.T) {
			scaffolder := &APIScaffolder{BaseDir: tc.baseDir}
			got := scaffolder.repoRoot()
			if got != tc.want {
				t.Errorf("repoRoot() with BaseDir %q = %q, want %q", tc.baseDir, got, tc.want)
			}
		})
	}
}

func TestServiceName(t *testing.T) {
	for _, tc := range []struct {
		name      string
		goPackage string
		group     string
		want      string
	}{
		{"from GoPackage simple", "chronicle/v1alpha1", "", "chronicle"},
		{"from GoPackage nested", "sql/v1beta1", "", "sql"},
		{"from Group fallback", "", "bigquery.cnrm.cloud.google.com", "bigquery"},
		{"empty scaffolder", "", "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a := &APIScaffolder{GoPackage: tc.goPackage, Group: tc.group}
			if got := a.serviceName(); got != tc.want {
				t.Errorf("serviceName() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestServiceMatchesPrefix(t *testing.T) {
	for _, tc := range []struct {
		service string
		prefix  string
		want    bool
	}{
		{"sql", "sql", true},
		{"sqladmin", "sql", true},
		{"sql", "sqladmin", true},
		{"chronicle", "sql", false},
		{"chronicle", "compute", false},
		{"compute", "compute", true},
		{"bigquery", "bigquery", true},
		{"iam", "gcpserviceaccount", true},
		{"iam", "iam", true},
		{"secretmanager", "secret", true},
		{"secretmanager", "secretmanager", true},
		{"alloydb", "alloydb", true},
		{"appengine", "appengine", true},
		{"kms", "kms", true},
		{"", "sql", false},
		{"sql", "", false},
		{"", "", false},
	} {
		got := serviceMatchesPrefix(tc.service, tc.prefix)
		if got != tc.want {
			t.Errorf("serviceMatchesPrefix(%q, %q) = %v, want %v", tc.service, tc.prefix, got, tc.want)
		}
	}
}

func TestIsAllowedSharedRef(t *testing.T) {
	for _, tc := range []struct {
		typeName       string
		currentService string
		wantRef        string
		wantAllowed    bool
	}{
		{"ProjectRef", "chronicle", "projectref", true},
		{"FolderRef", "chronicle", "folderref", true},
		{"OrganizationRef", "chronicle", "organizationref", true},
		{"BillingAccountRef", "chronicle", "billingaccountref", true},
		{"SQLInstanceRef", "chronicle", "instanceref", false},
		{"SQLInstanceRef", "sql", "instanceref", true},
		{"SQLInstanceRef", "sqladmin", "instanceref", true},
		{"SQLInstanceRef", "", "instanceref", false},
		{"ComputeNetworkRef", "chronicle", "networkref", false},
		{"ComputeNetworkRef", "compute", "networkref", true},
		{"BigQueryTableRef", "spanner", "tableref", false},
		{"BigQueryTableRef", "bigquery", "tableref", true},
	} {
		got := isAllowedSharedRef(tc.typeName, tc.currentService, tc.wantRef)
		if got != tc.wantAllowed {
			t.Errorf("isAllowedSharedRef(%q, %q, %q) = %v, want %v", tc.typeName, tc.currentService, tc.wantRef, got, tc.wantAllowed)
		}
	}
}

func TestCandidatesForTarget(t *testing.T) {
	repo := t.TempDir()
	sharedDir := filepath.Join(repo, sharedRefsPackage)
	if err := os.MkdirAll(sharedDir, 0755); err != nil {
		t.Fatal(err)
	}
	const sharedContent = `package v1beta1

type ProjectRef struct{}
type SQLInstanceRef struct{}
`
	if err := os.WriteFile(filepath.Join(sharedDir, "refs.go"), []byte(sharedContent), 0644); err != nil {
		t.Fatal(err)
	}

	svcDir := filepath.Join(repo, "apis", "chronicle", "v1alpha1")
	if err := os.MkdirAll(svcDir, 0755); err != nil {
		t.Fatal(err)
	}
	const svcContent = `package v1alpha1

type ChronicleInstanceRef struct{}
`
	if err := os.WriteFile(filepath.Join(svcDir, "instance_reference.go"), []byte(svcContent), 0644); err != nil {
		t.Fatal(err)
	}

	scaffolder := &APIScaffolder{
		BaseDir:   filepath.Join(repo, "apis"),
		GoPackage: "chronicle/v1alpha1",
	}

	t.Run("matches reference in serviceDir when target service matches current", func(t *testing.T) {
		got := scaffolder.candidatesForTarget("chronicle.googleapis.com/Instance", svcDir)
		if len(got) != 1 || got["ChronicleInstanceRef"] != "" {
			t.Errorf("candidatesForTarget(chronicle Instance) = %v, want map[ChronicleInstanceRef:]", got)
		}
	})

	t.Run("matches generic reference in shared refs package", func(t *testing.T) {
		got := scaffolder.candidatesForTarget("resourcemanager.googleapis.com/Project", svcDir)
		if len(got) != 1 || got["ProjectRef"] != "refsv1beta1" {
			t.Errorf("candidatesForTarget(Project) = %v, want map[ProjectRef:refsv1beta1]", got)
		}
	})

	t.Run("does not match foreign service-prefixed reference in shared refs for chronicle", func(t *testing.T) {
		emptySvcDir := t.TempDir()
		got := scaffolder.candidatesForTarget("chronicle.googleapis.com/Instance", emptySvcDir)
		if len(got) != 0 {
			t.Errorf("candidatesForTarget(chronicle Instance) = %v, want empty", got)
		}
	})

	t.Run("matches service-prefixed reference in shared refs when service matches", func(t *testing.T) {
		sqlScaffolder := &APIScaffolder{
			BaseDir:   filepath.Join(repo, "apis"),
			GoPackage: "sql/v1beta1",
		}
		got := sqlScaffolder.candidatesForTarget("sqladmin.googleapis.com/Instance", filepath.Join(repo, "apis", "sql", "v1beta1"))
		if len(got) != 1 || got["SQLInstanceRef"] != "refsv1beta1" {
			t.Errorf("candidatesForTarget(sqladmin Instance) for sql = %v, want map[SQLInstanceRef:refsv1beta1]", got)
		}
	})

	t.Run("returns nil for empty target", func(t *testing.T) {
		if got := scaffolder.candidatesForTarget("", svcDir); got != nil {
			t.Errorf("candidatesForTarget(\"\") = %v, want nil", got)
		}
	})

	t.Run("returns empty for unknown target", func(t *testing.T) {
		got := scaffolder.candidatesForTarget("foo.googleapis.com/Unknown", svcDir)
		if len(got) != 0 {
			t.Errorf("candidatesForTarget(unknown) = %v, want empty", got)
		}
	})
}

func TestParentRefTypes(t *testing.T) {
	repo := t.TempDir()
	sharedDir := filepath.Join(repo, sharedRefsPackage)
	if err := os.MkdirAll(sharedDir, 0755); err != nil {
		t.Fatal(err)
	}
	const sharedContent = `package v1beta1

type ProjectRef struct{}
type SQLInstanceRef struct{}
type KMSCryptoKeyRef struct{}
type kmsCryptoKeyRef struct{}
`
	if err := os.WriteFile(filepath.Join(sharedDir, "refs.go"), []byte(sharedContent), 0644); err != nil {
		t.Fatal(err)
	}

	svcDir := filepath.Join(repo, "apis", "chronicle", "v1alpha1")
	if err := os.MkdirAll(svcDir, 0755); err != nil {
		t.Fatal(err)
	}
	const svcContent = `package v1alpha1

type ChronicleInstanceRef struct{}
type InstanceRef struct{}
`
	if err := os.WriteFile(filepath.Join(svcDir, "instance_reference.go"), []byte(svcContent), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("service package ref wins over shared package", func(t *testing.T) {
		got := parentRefTypes(repo, svcDir, "chronicle", "instance")
		if len(got) != 2 || got["ChronicleInstanceRef"] != "" || got["InstanceRef"] != "" {
			t.Errorf("parentRefTypes = %v, want 2 service package types with empty qualifier", got)
		}
	})

	t.Run("shared package ref matches generic un-prefixed", func(t *testing.T) {
		emptySvcDir := t.TempDir()
		got := parentRefTypes(repo, emptySvcDir, "chronicle", "project")
		if len(got) != 1 || got["ProjectRef"] != "refsv1beta1" {
			t.Errorf("parentRefTypes = %v, want map[ProjectRef:refsv1beta1]", got)
		}
	})

	t.Run("shared package ref does not match foreign service", func(t *testing.T) {
		emptySvcDir := t.TempDir()
		got := parentRefTypes(repo, emptySvcDir, "chronicle", "instance")
		if len(got) != 0 {
			t.Errorf("parentRefTypes for chronicle instance = %v, want empty", got)
		}
	})

	t.Run("shared package ref matches matching service", func(t *testing.T) {
		emptySvcDir := t.TempDir()
		got := parentRefTypes(repo, emptySvcDir, "sql", "instance")
		if len(got) != 1 || got["SQLInstanceRef"] != "refsv1beta1" {
			t.Errorf("parentRefTypes for sql instance = %v, want map[SQLInstanceRef:refsv1beta1]", got)
		}
	})

	t.Run("unexported types in shared package are ignored", func(t *testing.T) {
		emptySvcDir := t.TempDir()
		got := parentRefTypes(repo, emptySvcDir, "kms", "cryptoKey")
		if len(got) != 1 || got["KMSCryptoKeyRef"] != "refsv1beta1" {
			t.Errorf("parentRefTypes for kms cryptoKey = %v, want map[KMSCryptoKeyRef:refsv1beta1]", got)
		}
	})
}
