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

package secretmanager

import (
	"context"
	"testing"

	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSecretVersion_Export_RedactsSecretPayload(t *testing.T) {
	ctx := context.Background()

	id, err := krm.ParseSecretVersionExternal("projects/test-project-123/secrets/my-secret-abc/versions/1")
	if err != nil {
		t.Fatalf("ParseSecretVersionExternal failed: %v", err)
	}

	actual := &pb.SecretVersion{
		Name:  "projects/test-project-123/secrets/my-secret-abc/versions/1",
		State: pb.SecretVersion_ENABLED,
	}

	adapter := &SecretVersionAdapter{
		id:     id,
		actual: actual,
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() failed: %v", err)
	}
	if u == nil {
		t.Fatal("Export() returned nil unstructured object")
	}

	// Verify GVK and Name
	if got, want := u.GetKind(), "SecretManagerSecretVersion"; got != want {
		t.Errorf("GetKind() = %q, want %q", got, want)
	}
	if got, want := u.GetName(), "1"; got != want {
		t.Errorf("GetName() = %q, want %q", got, want)
	}

	// Verify project annotation
	ann := u.GetAnnotations()
	if got, want := ann[k8s.ProjectIDAnnotation], "test-project-123"; got != want {
		t.Errorf("Project annotation = %q, want %q", got, want)
	}

	// Verify spec fields
	spec, ok := u.Object["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("spec is not a map[string]interface{}: %v", u.Object["spec"])
	}

	// CRITICAL SECURITY CHECK: secretData must NOT be present in exported spec
	if _, hasSecretData := spec["secretData"]; hasSecretData {
		t.Errorf("CRITICAL SECURITY FLAW: secretData is present in exported spec: %v", spec["secretData"])
	}

	// Verify resourceID and secretRef
	if got, want := spec["resourceID"], "1"; got != want {
		t.Errorf("resourceID = %v, want %v", got, want)
	}
	secretRef, ok := spec["secretRef"].(map[string]interface{})
	if !ok {
		t.Fatalf("secretRef is not a map: %v", spec["secretRef"])
	}
	if got, want := secretRef["external"], "projects/test-project-123/secrets/my-secret-abc"; got != want {
		t.Errorf("secretRef.external = %v, want %v", got, want)
	}
	if got, want := spec["enabled"], true; got != want {
		t.Errorf("enabled = %v, want %v", got, want)
	}
}

func TestSecretVersion_Export_DisabledVersion(t *testing.T) {
	ctx := context.Background()

	id, err := krm.ParseSecretVersionExternal("projects/test-project-123/secrets/my-secret-abc/versions/2")
	if err != nil {
		t.Fatalf("ParseSecretVersionExternal failed: %v", err)
	}

	actual := &pb.SecretVersion{
		Name:  "projects/test-project-123/secrets/my-secret-abc/versions/2",
		State: pb.SecretVersion_DISABLED,
	}

	adapter := &SecretVersionAdapter{
		id:     id,
		actual: actual,
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() failed: %v", err)
	}

	enabled, found, err := unstructured.NestedBool(u.Object, "spec", "enabled")
	if err != nil {
		t.Fatalf("error checking spec.enabled: %v", err)
	}
	if !found || enabled != false {
		t.Errorf("expected spec.enabled to be false, found=%v, val=%v", found, enabled)
	}

	// Verify secretData is not present
	if _, found, _ := unstructured.NestedFieldNoCopy(u.Object, "spec", "secretData"); found {
		t.Errorf("secretData should not be present in exported manifest")
	}
}

func TestSecret_Export(t *testing.T) {
	ctx := context.Background()

	id, err := krm.ParseSecretExternal("projects/test-project-123/secrets/my-secret-abc")
	if err != nil {
		t.Fatalf("ParseSecretExternal failed: %v", err)
	}

	actual := &pb.Secret{
		Name: "projects/test-project-123/secrets/my-secret-abc",
		Replication: &pb.Replication{
			Replication: &pb.Replication_Automatic_{
				Automatic: &pb.Replication_Automatic{},
			},
		},
		Labels: map[string]string{
			"env": "production",
		},
		VersionAliases: map[string]int64{
			"prod": 1,
		},
	}

	adapter := &Adapter{
		id:     id,
		actual: actual,
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() failed: %v", err)
	}
	if u == nil {
		t.Fatal("Export() returned nil unstructured object")
	}

	if got, want := u.GetKind(), "SecretManagerSecret"; got != want {
		t.Errorf("GetKind() = %q, want %q", got, want)
	}
	if got, want := u.GetName(), "my-secret-abc"; got != want {
		t.Errorf("GetName() = %q, want %q", got, want)
	}

	// Verify labels
	labels := u.GetLabels()
	if got, want := labels["env"], "production"; got != want {
		t.Errorf("label env = %q, want %q", got, want)
	}

	// Verify version aliases
	alias, found, err := unstructured.NestedString(u.Object, "spec", "versionAliases", "prod")
	if err != nil || !found || alias != "1" {
		t.Errorf("versionAliases.prod = %q (found: %v), want '1'", alias, found)
	}
}

func TestAdapterForURL_Parsing(t *testing.T) {
	ctx := context.Background()
	cfg := &config.ControllerConfig{}

	secModel, err := NewModel(ctx, cfg)
	if err != nil {
		t.Fatalf("NewModel failed: %v", err)
	}
	secVerModel, err := NewSecretVersionModel(ctx, cfg)
	if err != nil {
		t.Fatalf("NewSecretVersionModel failed: %v", err)
	}

	secretURL := "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret"
	versionURL := "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret/versions/1"
	invalidURL := "//secretmanager.googleapis.com/projects/my-project/other/res"

	// Test Secret Model URL handling
	ad1, err := secModel.AdapterForURL(ctx, secretURL)
	if err != nil || ad1 == nil {
		t.Errorf("secModel.AdapterForURL(%q) returned nil, expected adapter", secretURL)
	}
	ad2, err := secModel.AdapterForURL(ctx, versionURL)
	if err != nil || ad2 != nil {
		t.Errorf("secModel.AdapterForURL(%q) returned %v, expected nil", versionURL, ad2)
	}

	// Test SecretVersion Model URL handling
	ad3, err := secVerModel.AdapterForURL(ctx, versionURL)
	if err != nil || ad3 == nil {
		t.Errorf("secVerModel.AdapterForURL(%q) returned nil, expected adapter", versionURL)
	}
	ad4, err := secVerModel.AdapterForURL(ctx, secretURL)
	if err != nil || ad4 != nil {
		t.Errorf("secVerModel.AdapterForURL(%q) returned %v, expected nil", secretURL, ad4)
	}

	// Test Invalid URL
	if ad, _ := secModel.AdapterForURL(ctx, invalidURL); ad != nil {
		t.Errorf("secModel.AdapterForURL(%q) returned %v, expected nil", invalidURL, ad)
	}
	if ad, _ := secVerModel.AdapterForURL(ctx, invalidURL); ad != nil {
		t.Errorf("secVerModel.AdapterForURL(%q) returned %v, expected nil", invalidURL, ad)
	}
}
