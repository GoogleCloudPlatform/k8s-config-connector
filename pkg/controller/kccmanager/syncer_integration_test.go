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

package kccmanager

import (
	"context"
	"fmt"
	"testing"

	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func buildScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()
	_ = mclv1alpha1.AddToScheme(scheme)
	return scheme
}

func newKRMSyncer(name types.NamespacedName, mode string, dest string, suspend bool) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KRMSyncerGVK)
	u.SetName(name.Name)
	u.SetNamespace(name.Namespace)

	spec := map[string]interface{}{}
	if mode != "" {
		spec["mode"] = mode
	}
	if dest != "" {
		if err := unstructured.SetNestedField(u.Object, dest, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name"); err != nil {
			panic(fmt.Sprintf("failed to set dest field: %v", err))
		}
	}
	spec["suspend"] = suspend

	// We merge the suspend and mode into the existing object if dest was set
	existingSpec, _, _ := unstructured.NestedMap(u.Object, "spec")
	if existingSpec != nil {
		for k, v := range spec {
			existingSpec[k] = v
		}
		if err := unstructured.SetNestedMap(u.Object, existingSpec, "spec"); err != nil {
			panic(fmt.Sprintf("failed to set spec field: %v", err))
		}
	} else {
		if err := unstructured.SetNestedMap(u.Object, spec, "spec"); err != nil {
			panic(fmt.Sprintf("failed to set spec field: %v", err))
		}
	}
	return u
}

func newMultiClusterLease(name types.NamespacedName, globalIdentity string) *mclv1alpha1.MultiClusterLease {
	lease := &mclv1alpha1.MultiClusterLease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name.Name,
			Namespace: name.Namespace,
		},
	}
	if globalIdentity != "" {
		lease.Status.GlobalHolderIdentity = &globalIdentity
	}
	return lease
}

func TestEnsurePullingFromLeader(t *testing.T) {
	// Note: fake.Client used in these tests does not emulate controller-runtime cache lifecycle.
	// It will not catch 'cache not started' errors that occur with the real cached client.
	name := types.NamespacedName{Name: "test-syncer", Namespace: "default"}
	myIdentity := "cluster-a"

	tests := []struct {
		name           string
		syncer         *unstructured.Unstructured
		lease          *mclv1alpha1.MultiClusterLease
		expectErr      bool
		validateSyncer func(*testing.T, client.Client)
	}{
		{
			name:      "no lease found",
			syncer:    newKRMSyncer(name, "pull", "", true),
			lease:     nil,
			expectErr: false,
		},
		{
			name:      "no global holder identity",
			syncer:    newKRMSyncer(name, "pull", "", true),
			lease:     newMultiClusterLease(name, ""),
			expectErr: false,
			validateSyncer: func(t *testing.T, c client.Client) {
				// Should be unchanged
				s := &unstructured.Unstructured{}
				s.SetGroupVersionKind(KRMSyncerGVK)
				_ = c.Get(context.TODO(), name, s)
				remote, _, _ := unstructured.NestedString(s.Object, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
				if remote != "" {
					t.Errorf("expected empty dest, got %q", remote)
				}
			},
		},
		{
			name:      "we are the leader",
			syncer:    newKRMSyncer(name, "pull", "", true),
			lease:     newMultiClusterLease(name, myIdentity),
			expectErr: false,
			validateSyncer: func(t *testing.T, c client.Client) {
				// Should be unchanged
				s := &unstructured.Unstructured{}
				s.SetGroupVersionKind(KRMSyncerGVK)
				_ = c.Get(context.TODO(), name, s)
				remote, _, _ := unstructured.NestedString(s.Object, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
				if remote != "" {
					t.Errorf("expected empty dest, got %q", remote)
				}
			},
		},
		{
			name:      "another cluster is leader - update syncer (mode was push)",
			syncer:    newKRMSyncer(name, "push", "", true),
			lease:     newMultiClusterLease(name, "cluster-b"),
			expectErr: false,
			validateSyncer: func(t *testing.T, c client.Client) {
				s := &unstructured.Unstructured{}
				s.SetGroupVersionKind(KRMSyncerGVK)
				_ = c.Get(context.TODO(), name, s)

				mode, _, _ := unstructured.NestedString(s.Object, "spec", "mode")
				if mode != "pull" {
					t.Errorf("expected mode to be 'pull', got %q", mode)
				}
				remote, _, _ := unstructured.NestedString(s.Object, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
				if remote != "cluster-b" {
					t.Errorf("expected dest to be 'cluster-b', got %q", remote)
				}
				suspend, _, _ := unstructured.NestedBool(s.Object, "spec", "suspend")
				if suspend {
					t.Errorf("expected suspend to be false")
				}

				rules, found, _ := unstructured.NestedSlice(s.Object, "spec", "rules")
				if !found {
					t.Fatalf("expected rules to be found")
				}

				// Expected length: 1 (glob) + number of exceptions
				expectedLen := 1 + len(syncerGVKsWithServiceGeneratedIDs)
				if len(rules) != expectedLen {
					t.Fatalf("expected %d rules, got %v", expectedLen, len(rules))
				}

				// Check for the glob rule
				hasGlob := false
				for _, r := range rules {
					rule := r.(map[string]interface{})
					if rule["group"] == "*.cnrm.cloud.google.com" {
						hasGlob = true
						if rule["version"] != "*" || rule["kind"] != "*" {
							t.Errorf("expected wildcard version and kind for glob rule")
						}
						syncFields := rule["syncFields"].([]interface{})
						if len(syncFields) != 1 || syncFields[0] != "status" {
							t.Errorf("expected syncFields [status] for glob rule, got %v", syncFields)
						}
						break
					}
				}
				if !hasGlob {
					t.Errorf("glob rule not found")
				}

				// Spot check one exception
				hasFolder := false
				folderGVK := schema.GroupVersionKind{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Folder"}
				for _, r := range rules {
					rule := r.(map[string]interface{})
					if rule["group"] == folderGVK.Group && rule["version"] == folderGVK.Version && rule["kind"] == folderGVK.Kind {
						hasFolder = true
						syncFields := rule["syncFields"].([]interface{})
						if len(syncFields) != 2 || syncFields[0] != "status" || syncFields[1] != "spec.resourceID" {
							t.Errorf("expected syncFields [status, spec.resourceID] for Folder rule, got %v", syncFields)
						}
						break
					}
				}
				if !hasFolder {
					t.Errorf("Folder exception rule not found")
				}
			},
		},
		{
			name:      "another cluster is leader - no changes needed",
			syncer:    newKRMSyncer(name, "pull", "cluster-b", false),
			lease:     newMultiClusterLease(name, "cluster-b"),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := fake.NewClientBuilder().WithScheme(buildScheme())
			if tt.syncer != nil {
				builder.WithObjects(tt.syncer)
			}
			if tt.lease != nil {
				builder.WithObjects(tt.lease)
			}
			c := builder.Build()

			si := &SyncerIntegration{
				client:          c,
				apiReader:       c,
				name:            name,
				replicationMode: "Status",
			}

			err := si.EnsurePullingFromLeader(context.Background(), myIdentity)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}

			if tt.validateSyncer != nil {
				tt.validateSyncer(t, c)
			}
		})
	}
}

func TestEnsureSuspended(t *testing.T) {
	name := types.NamespacedName{Name: "test-syncer", Namespace: "default"}

	tests := []struct {
		name           string
		syncer         *unstructured.Unstructured
		expectErr      bool
		validateSyncer func(*testing.T, client.Client)
	}{
		{
			name:      "already suspended",
			syncer:    newKRMSyncer(name, "pull", "", true),
			expectErr: false,
		},
		{
			name:      "needs suspending",
			syncer:    newKRMSyncer(name, "pull", "", false),
			expectErr: false,
			validateSyncer: func(t *testing.T, c client.Client) {
				s := &unstructured.Unstructured{}
				s.SetGroupVersionKind(KRMSyncerGVK)
				_ = c.Get(context.TODO(), name, s)
				suspend, _, _ := unstructured.NestedBool(s.Object, "spec", "suspend")
				if !suspend {
					t.Errorf("expected suspend to be true")
				}
			},
		},
		{
			name:      "syncer not found",
			syncer:    nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := fake.NewClientBuilder().WithScheme(buildScheme())
			if tt.syncer != nil {
				builder.WithObjects(tt.syncer)
			}
			c := builder.Build()

			si := &SyncerIntegration{
				client:    c,
				apiReader: c,
				name:      name,
			}

			err := si.EnsureSuspended(context.Background())
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}

			if tt.validateSyncer != nil {
				tt.validateSyncer(t, c)
			}
		})
	}
}
