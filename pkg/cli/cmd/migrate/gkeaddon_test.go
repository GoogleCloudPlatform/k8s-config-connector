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

package migrate

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestMigrateGKEAddonCmd(t *testing.T) {
	// Verify command and subcommands are registered correctly
	root := &cobra.Command{Use: "config-connector"}
	AddCommands(root)

	migrateCmd, _, err := root.Find([]string{"gke-addon"})
	if err != nil {
		t.Fatalf("could not find gke-addon command: %v", err)
	}

	prepareCmd, _, err := migrateCmd.Find([]string{"prepare"})
	if err != nil {
		t.Errorf("could not find prepare subcommand: %v", err)
	}
	if prepareCmd.Name() != "prepare" {
		t.Errorf("expected prepare subcommand name to be 'prepare', got %q", prepareCmd.Name())
	}

	finishCmd, _, err := migrateCmd.Find([]string{"finish"})
	if err != nil {
		t.Errorf("could not find finish subcommand: %v", err)
	}
	if finishCmd.Name() != "finish" {
		t.Errorf("expected finish subcommand name to be 'finish', got %q", finishCmd.Name())
	}
}

func TestCleanResourceForBackup(t *testing.T) {
	res := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "test.cnrm.cloud.google.com/v1beta1",
			"kind":       "TestResource",
			"metadata": map[string]interface{}{
				"name":              "test-resource",
				"namespace":         "test-ns",
				"resourceVersion":  "12345",
				"uid":              "uid-123",
				"creationTimestamp": "2026-03-03T00:00:00Z",
				"managedFields":     []interface{}{"field1"},
				"deletionTimestamp": "2026-03-03T00:00:00Z",
				"finalizers":        []interface{}{"finalizer1"},
			},
			"spec": map[string]interface{}{
				"foo": "bar",
			},
			"status": map[string]interface{}{
				"ready": true,
			},
		},
	}

	cleanResourceForBackup(res)

	if _, found := res.Object["status"]; found {
		t.Errorf("status field should have been removed")
	}

	metadata := res.Object["metadata"].(map[string]interface{})
	for _, field := range []string{"resourceVersion", "uid", "creationTimestamp", "managedFields", "deletionTimestamp"} {
		if _, found := metadata[field]; found {
			t.Errorf("metadata field %q should have been removed", field)
		}
	}

	if _, found := metadata["name"]; !found {
		t.Errorf("metadata field 'name' should have been preserved")
	}

	if _, found := res.Object["spec"]; !found {
		t.Errorf("spec field should have been preserved")
	}
}

func TestPauseReconciliation(t *testing.T) {
	ctx := context.Background()

	cc := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "core.cnrm.cloud.google.com/v1beta1",
			"kind":       "ConfigConnector",
			"metadata": map[string]interface{}{
				"name": "configconnector.core.cnrm.cloud.google.com",
			},
			"spec": map[string]interface{}{
				"mode": "cluster",
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithRuntimeObjects(cc).Build()
	kubeClient := &kubecli.Client{
		Client: fakeClient,
	}

	if err := pauseReconciliation(ctx, kubeClient); err != nil {
		t.Fatalf("pauseReconciliation failed: %v", err)
	}

	// Verify the object was updated in the fake client
	updatedCC := &unstructured.Unstructured{}
	updatedCC.SetGroupVersionKind(cc.GroupVersionKind())
	if err := fakeClient.Get(ctx, types.NamespacedName{Name: cc.GetName()}, updatedCC); err != nil {
		t.Fatalf("failed to get updated CC: %v", err)
	}

	mode, found, err := unstructured.NestedString(updatedCC.Object, "spec", "actuationMode")
	if err != nil || !found {
		t.Errorf("actuationMode not found or error: %v, found: %v", err, found)
	}
	if mode != "Paused" {
		t.Errorf("expected actuationMode to be 'Paused', got %q", mode)
	}
}

func TestResumeReconciliation(t *testing.T) {
	ctx := context.Background()

	cc := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "core.cnrm.cloud.google.com/v1beta1",
			"kind":       "ConfigConnector",
			"metadata": map[string]interface{}{
				"name": "configconnector.core.cnrm.cloud.google.com",
			},
			"spec": map[string]interface{}{
				"mode":          "cluster",
				"actuationMode": "Paused",
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithRuntimeObjects(cc).Build()
	kubeClient := &kubecli.Client{
		Client: fakeClient,
	}

	if err := resumeReconciliation(ctx, kubeClient); err != nil {
		t.Fatalf("resumeReconciliation failed: %v", err)
	}

	// Verify the object was updated in the fake client
	updatedCC := &unstructured.Unstructured{}
	updatedCC.SetGroupVersionKind(cc.GroupVersionKind())
	if err := fakeClient.Get(ctx, types.NamespacedName{Name: cc.GetName()}, updatedCC); err != nil {
		t.Fatalf("failed to get updated CC: %v", err)
	}

	mode, found, err := unstructured.NestedString(updatedCC.Object, "spec", "actuationMode")
	if err != nil || !found {
		t.Errorf("actuationMode not found or error: %v, found: %v", err, found)
	}
	if mode != "Reconciling" {
		t.Errorf("expected actuationMode to be 'Reconciling', got %q", mode)
	}
}

func TestRemoveFinalizers(t *testing.T) {
	ctx := context.Background()

	res := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "test.cnrm.cloud.google.com/v1beta1",
			"kind":       "TestResource",
			"metadata": map[string]interface{}{
				"name":       "test-resource",
				"namespace":  "test-ns",
				"finalizers": []interface{}{"finalizer1", "finalizer2"},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithRuntimeObjects(res).Build()
	kubeClient := &kubecli.Client{
		Client: fakeClient,
	}

	resources := []unstructured.Unstructured{*res}

	if err := removeFinalizers(ctx, kubeClient, resources); err != nil {
		t.Fatalf("removeFinalizers failed: %v", err)
	}

	// Verify finalizers were removed in the fake client
	updatedRes := &unstructured.Unstructured{}
	updatedRes.SetGroupVersionKind(res.GroupVersionKind())
	if err := fakeClient.Get(ctx, types.NamespacedName{Name: res.GetName(), Namespace: res.GetNamespace()}, updatedRes); err != nil {
		t.Fatalf("failed to get updated resource: %v", err)
	}

	if len(updatedRes.GetFinalizers()) != 0 {
		t.Errorf("expected 0 finalizers, got %d: %v", len(updatedRes.GetFinalizers()), updatedRes.GetFinalizers())
	}
}
