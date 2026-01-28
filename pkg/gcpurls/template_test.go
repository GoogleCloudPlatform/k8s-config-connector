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

package gcpurls

import (
	"testing"
)

type ProjectIdentity struct {
	Project string
}

type TableIdentity struct {
	Project string
	Dataset string
	Table   string
}

// AlloyDBBackupIdentity corresponds to alloydb.googleapis.com/Backup
// Format: //alloydb.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/backups/{{BACKUP}}
type AlloyDBBackupIdentity struct {
	Project  string
	Location string
	Backup   string
}

// ApigeeInstanceIdentity corresponds to apigee.googleapis.com/Instance
// Format: //apigee.googleapis.com/organizations/{{ORGANIZATION_ID}}/instances/{{INSTANCE_ID}}
type ApigeeInstanceIdentity struct {
	OrganizationID string
	InstanceID     string
}

// AppEngineServiceIdentity corresponds to appengine.googleapis.com/Service
// Format: //appengine.googleapis.com/apps/{{APP}}/services/{{SERVICE}}
type AppEngineServiceIdentity struct {
	App     string
	Service string
}

// K8sDeploymentIdentity corresponds to apps.k8s.io/Deployment in GKE
// Format: //container.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/clusters/{{CLUSTER}}/k8s/namespaces/{{NAMESPACE}}/apps/deployments/{{DEPLOYMENT}}
type K8sDeploymentIdentity struct {
	Project    string
	Location   string
	Cluster    string
	Namespace  string
	Deployment string
}

func TestTemplate(t *testing.T) {
	// Test case 1: Simple project
	tmpl1 := Template[ProjectIdentity]("cloudresourcemanager.googleapis.com", "projects/{project}")
	id1, match, err := tmpl1.Parse("projects/my-project")
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	if !match {
		t.Errorf("Expected match for 'projects/my-project'")
	}
	if id1.Project != "my-project" {
		t.Errorf("Expected project 'my-project', got '%s'", id1.Project)
	}

	// Test with host
	id1Host, match, err := tmpl1.Parse("//cloudresourcemanager.googleapis.com/projects/my-project")
	if err != nil {
		t.Errorf("Parse with host failed: %v", err)
	}
	if !match {
		t.Errorf("Expected match for url with host")
	}
	if id1Host.Project != "my-project" {
		t.Errorf("Expected project 'my-project' (with host), got '%s'", id1Host.Project)
	}

	str1 := tmpl1.ToString(*id1)
	if str1 != "projects/my-project" {
		t.Errorf("ToString mismatch: expected 'projects/my-project', got '%s'", str1)
	}

	// Test case 3: Table
	tmpl3 := Template[TableIdentity]("bigquery.googleapis.com", "projects/{project}/datasets/{dataset}/tables/{table}")
	id3, match, err := tmpl3.Parse("projects/p1/datasets/d1/tables/t1")
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	if !match {
		t.Errorf("Expected match")
	}
	if id3.Project != "p1" || id3.Dataset != "d1" || id3.Table != "t1" {
		t.Errorf("Parse mismatch: got %v", id3)
	}

	// Test with host for Table
	id3Host, match, err := tmpl3.Parse("//bigquery.googleapis.com/projects/p1/datasets/d1/tables/t1")
	if err != nil {
		t.Errorf("Parse with host failed: %v", err)
	}
	if !match {
		t.Errorf("Expected match with host")
	}
	if id3Host.Project != "p1" || id3Host.Dataset != "d1" || id3Host.Table != "t1" {
		t.Errorf("Parse mismatch (with host): got %v", id3Host)
	}

	// Test case insensitive field matching
	type LowerCaseIdentity struct {
		Foo string
	}
	tmpl4 := Template[LowerCaseIdentity]("", "items/{foo}")
	id4, match, err := tmpl4.Parse("items/bar")
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	if !match {
		t.Errorf("Expected match")
	}
	if id4.Foo != "bar" {
		t.Errorf("Expected Foo='bar', got '%s'", id4.Foo)
	}

	// Test Mismatch
	_, match, _ = tmpl1.Parse("folders/my-folder")
	if match {
		t.Errorf("Expected no match for mismatching static segment")
	}
}

func TestCAIExamples(t *testing.T) {
	// 1. AlloyDB Backup
	alloyDBTmpl := Template[AlloyDBBackupIdentity]("alloydb.googleapis.com", "projects/{project}/locations/{location}/backups/{backup}")
	alloyDBURL := "//alloydb.googleapis.com/projects/my-project/locations/us-central1/backups/my-backup"

	alloyDBId, match, err := alloyDBTmpl.Parse(alloyDBURL)
	if err != nil {
		t.Fatalf("Failed to parse AlloyDB URL: %v", err)
	}
	if !match {
		t.Fatalf("Expected AlloyDB URL to match")
	}
	expectedAlloyDB := AlloyDBBackupIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Backup:   "my-backup",
	}
	if *alloyDBId != expectedAlloyDB {
		t.Errorf("AlloyDB match failed: got %+v, expected %+v", *alloyDBId, expectedAlloyDB)
	}

	// Reverse
	s := alloyDBTmpl.ToString(*alloyDBId)
	if s != "projects/my-project/locations/us-central1/backups/my-backup" {
		t.Errorf("AlloyDB ToString mismatch: got %q", s)
	}

	// 2. Apigee Instance
	apigeeTmpl := Template[ApigeeInstanceIdentity]("apigee.googleapis.com", "organizations/{OrganizationID}/instances/{InstanceID}")
	apigeeURL := "//apigee.googleapis.com/organizations/my-org/instances/my-instance"

	apigeeId, match, err := apigeeTmpl.Parse(apigeeURL)
	if err != nil {
		t.Fatalf("Failed to parse Apigee URL: %v", err)
	}
	if !match {
		t.Fatalf("Expected Apigee URL to match")
	}
	expectedApigee := ApigeeInstanceIdentity{
		OrganizationID: "my-org",
		InstanceID:     "my-instance",
	}
	if *apigeeId != expectedApigee {
		t.Errorf("Apigee match failed: got %+v, expected %+v", *apigeeId, expectedApigee)
	}

	// 3. AppEngine Service
	appEngineTmpl := Template[AppEngineServiceIdentity]("appengine.googleapis.com", "apps/{app}/services/{service}")
	appEngineURL := "//appengine.googleapis.com/apps/my-app/services/default"

	appEngineId, match, err := appEngineTmpl.Parse(appEngineURL)
	if err != nil {
		t.Fatalf("Failed to parse AppEngine URL: %v", err)
	}
	if !match {
		t.Fatalf("Expected AppEngine URL to match")
	}
	expectedAppEngine := AppEngineServiceIdentity{
		App:     "my-app",
		Service: "default",
	}
	if *appEngineId != expectedAppEngine {
		t.Errorf("AppEngine match failed: got %+v, expected %+v", *appEngineId, expectedAppEngine)
	}

	// 4. K8s Deployment
	k8sTmpl := Template[K8sDeploymentIdentity]("container.googleapis.com", "projects/{project}/locations/{location}/clusters/{cluster}/k8s/namespaces/{namespace}/apps/deployments/{deployment}")
	k8sURL := "//container.googleapis.com/projects/p1/locations/l1/clusters/c1/k8s/namespaces/ns1/apps/deployments/d1"

	k8sId, match, err := k8sTmpl.Parse(k8sURL)
	if err != nil {
		t.Fatalf("Failed to parse K8s URL: %v", err)
	}
	if !match {
		t.Fatalf("Expected K8s URL to match")
	}
	expectedK8s := K8sDeploymentIdentity{
		Project:    "p1",
		Location:   "l1",
		Cluster:    "c1",
		Namespace:  "ns1",
		Deployment: "d1",
	}
	if *k8sId != expectedK8s {
		t.Errorf("K8s match failed: got %+v, expected %+v", *k8sId, expectedK8s)
	}
}

func TestEdgeCases(t *testing.T) {
	// Mismatch Host
	tmpl := Template[ProjectIdentity]("example.com", "projects/{project}")
	_, match, _ := tmpl.Parse("//other.com/projects/my-project")
	if match {
		t.Errorf("Expected mismatch for wrong host")
	}

	// Relative path should match even if host is configured
	id, match, err := tmpl.Parse("projects/my-project")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !match {
		t.Errorf("Expected match for relative path")
	}
	if id.Project != "my-project" {
		t.Errorf("Expected project 'my-project', got %s", id.Project)
	}

	// Leading slash on relative path
	idLeading, match, err := tmpl.Parse("/projects/my-project")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !match {
		t.Errorf("Expected match for leading slash")
	}
	if idLeading.Project != "my-project" {
		t.Errorf("Expected project 'my-project', got %s", idLeading.Project)
	}

	// Extra segments
	_, match, _ = tmpl.Parse("projects/my-project/extra")
	if match {
		t.Errorf("Expected mismatch for extra segments")
	}

	// Missing segments
	_, match, _ = tmpl.Parse("projects")
	if match {
		t.Errorf("Expected mismatch for missing segments")
	}

	// Wrong static segment
	_, match, _ = tmpl.Parse("folders/my-project")
	if match {
		t.Errorf("Expected mismatch for wrong static segment")
	}
}

func TestPanics(t *testing.T) {
	// Helper to check for panic
	assertPanic := func(msg string, f func()) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic for %s", msg)
			}
		}()
		f()
	}

	// 1. Duplicate case-insensitive fields
	type Ambiguous struct {
		Foo string
		FOO string
	}
	assertPanic("ambiguous fields", func() {
		Template[Ambiguous]("", "items/{foo}")
	})

	// 1b. Duplicate case-insensitive fields with exact match
	// Even if one matches exactly, it should panic because it is ambiguous case-insensitively
	assertPanic("ambiguous fields with exact match", func() {
		Template[Ambiguous]("", "items/{Foo}")
	})

	// 2. Field not string
	type NotString struct {
		Foo int
	}
	assertPanic("non-string field", func() {
		Template[NotString]("", "items/{foo}")
	})

	// 3. Field not found
	type MissingField struct {
		Bar string
	}
	assertPanic("missing field", func() {
		Template[MissingField]("", "items/{foo}")
	})

	// 4. Not a struct
	assertPanic("not a struct", func() {
		Template[string]("", "items/{foo}")
	})
}
