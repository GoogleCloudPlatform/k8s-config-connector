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
	"testing"

	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
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
		spec["destinationKubeConfig"] = dest
	}
	spec["suspend"] = suspend

	_ = unstructured.SetNestedMap(u.Object, spec, "spec")
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
			expectErr: true,
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
				dest, _, _ := unstructured.NestedString(s.Object, "spec", "destinationKubeConfig")
				if dest != "" {
					t.Errorf("expected empty dest, got %q", dest)
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
				dest, _, _ := unstructured.NestedString(s.Object, "spec", "destinationKubeConfig")
				if dest != "" {
					t.Errorf("expected empty dest, got %q", dest)
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
				dest, _, _ := unstructured.NestedString(s.Object, "spec", "destinationKubeConfig")
				if dest != "cluster-b" {
					t.Errorf("expected dest to be 'cluster-b', got %q", dest)
				}
				suspend, _, _ := unstructured.NestedBool(s.Object, "spec", "suspend")
				if suspend {
					t.Errorf("expected suspend to be false")
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
				client: c,
				name:   name,
			}

			err := si.EnsurePullingFromLeader(context.Background(), c, myIdentity)
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
			expectErr: true,
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
				client: c,
				name:   name,
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

func TestEnsureSyncing(t *testing.T) {
	name := types.NamespacedName{Name: "test-syncer", Namespace: "default"}

	tests := []struct {
		name           string
		syncer         *unstructured.Unstructured
		expectErr      bool
		validateSyncer func(*testing.T, client.Client)
	}{
		{
			name:      "already syncing",
			syncer:    newKRMSyncer(name, "pull", "", false),
			expectErr: false,
		},
		{
			name:      "needs syncing",
			syncer:    newKRMSyncer(name, "pull", "", true),
			expectErr: false,
			validateSyncer: func(t *testing.T, c client.Client) {
				s := &unstructured.Unstructured{}
				s.SetGroupVersionKind(KRMSyncerGVK)
				_ = c.Get(context.TODO(), name, s)
				suspend, _, _ := unstructured.NestedBool(s.Object, "spec", "suspend")
				if suspend {
					t.Errorf("expected suspend to be false")
				}
			},
		},
		{
			name:      "syncer not found",
			syncer:    nil,
			expectErr: true,
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
				client: c,
				name:   name,
			}

			err := si.EnsureSyncing(context.Background())
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}

			if tt.validateSyncer != nil {
				tt.validateSyncer(t, c)
			}
		})
	}
}
