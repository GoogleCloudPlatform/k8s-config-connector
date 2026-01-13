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

func TestTemplate(t *testing.T) {
	// Test case 1: Simple project
	tmpl1 := Template[ProjectIdentity]("projects/{project}")
	id1, err := tmpl1.Parse("projects/my-project")
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	if id1.Project != "my-project" {
		t.Errorf("Expected project 'my-project', got '%s'", id1.Project)
	}

	// Test with host
	id1Host, err := tmpl1.Parse("//cloudresourcemanager.googleapis.com/projects/my-project")
	if err != nil {
		t.Errorf("Parse with host failed: %v", err)
	}
	if id1Host.Project != "my-project" {
		t.Errorf("Expected project 'my-project' (with host), got '%s'", id1Host.Project)
	}

	str1, err := tmpl1.ToString(*id1)
	if err != nil {
		t.Errorf("ToString failed: %v", err)
	}
	if str1 != "projects/my-project" {
		t.Errorf("ToString mismatch: expected 'projects/my-project', got '%s'", str1)
	}

	// Test case 3: Table
	tmpl3 := Template[TableIdentity]("projects/{project}/datasets/{dataset}/tables/{table}")
	id3, err := tmpl3.Parse("projects/p1/datasets/d1/tables/t1")
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	if id3.Project != "p1" || id3.Dataset != "d1" || id3.Table != "t1" {
		t.Errorf("Parse mismatch: got %v", id3)
	}

	// Test with host for Table
	id3Host, err := tmpl3.Parse("//bigquery.googleapis.com/projects/p1/datasets/d1/tables/t1")
	if err != nil {
		t.Errorf("Parse with host failed: %v", err)
	}
	if id3Host.Project != "p1" || id3Host.Dataset != "d1" || id3Host.Table != "t1" {
		t.Errorf("Parse mismatch (with host): got %v", id3Host)
	}

	// Test case insensitive field matching
	type LowerCaseIdentity struct {
		Foo string
	}
	tmpl4 := Template[LowerCaseIdentity]("items/{foo}")
	id4, err := tmpl4.Parse("items/bar")
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	if id4.Foo != "bar" {
		t.Errorf("Expected Foo='bar', got '%s'", id4.Foo)
	}

	// Test Mismatch
	_, err = tmpl1.Parse("folders/my-folder")
	if err == nil {
		t.Errorf("Expected error for mismatching static segment")
	}
}
