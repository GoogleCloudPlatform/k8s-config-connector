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

package sql

import (
	"context"
	"testing"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestResolveSourceSQLInstanceRef(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := krm.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add krm to scheme: %v", err)
	}

	referencedSQLInstance := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "k8s-source-instance",
				"namespace": "test-ns",
			},
			"spec": map[string]any{
				"resourceID": "resolved-source-id",
			},
		},
	}

	kube := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(referencedSQLInstance).
		Build()

	ctx := context.Background()

	tests := []struct {
		name         string
		cloneSource  *krm.CloneSource
		wantExternal string
		wantErr      bool
	}{
		{
			name:        "nil cloneSource",
			cloneSource: nil,
			wantErr:     false,
		},
		{
			name: "bare instance name in external",
			cloneSource: &krm.CloneSource{
				SQLInstanceRef: refs.SQLInstanceRef{
					External: "my-source-instance",
				},
			},
			wantExternal: "my-source-instance",
			wantErr:      false,
		},
		{
			name: "resource path in external",
			cloneSource: &krm.CloneSource{
				SQLInstanceRef: refs.SQLInstanceRef{
					External: "projects/my-project/instances/my-source-instance",
				},
			},
			wantExternal: "my-source-instance",
			wantErr:      false,
		},
		{
			name: "selfLink in external",
			cloneSource: &krm.CloneSource{
				SQLInstanceRef: refs.SQLInstanceRef{
					External: "https://sqladmin.googleapis.com/sql/v1beta4/projects/my-project/instances/my-source-instance",
				},
			},
			wantExternal: "my-source-instance",
			wantErr:      false,
		},
		{
			name: "trailing slash in external",
			cloneSource: &krm.CloneSource{
				SQLInstanceRef: refs.SQLInstanceRef{
					External: "projects/my-project/instances/my-source-instance/",
				},
			},
			wantExternal: "my-source-instance",
			wantErr:      false,
		},
		{
			name: "valid name reference",
			cloneSource: &krm.CloneSource{
				SQLInstanceRef: refs.SQLInstanceRef{
					Name:      "k8s-source-instance",
					Namespace: "test-ns",
				},
			},
			wantExternal: "resolved-source-id",
			wantErr:      false,
		},
		{
			name: "both external and name specified",
			cloneSource: &krm.CloneSource{
				SQLInstanceRef: refs.SQLInstanceRef{
					External: "my-source-instance",
					Name:     "k8s-source-instance",
				},
			},
			wantErr: true,
		},
		{
			name: "neither external nor name specified",
			cloneSource: &krm.CloneSource{
				SQLInstanceRef: refs.SQLInstanceRef{},
			},
			wantErr: true,
		},
		{
			name: "referenced SQLInstance not found",
			cloneSource: &krm.CloneSource{
				SQLInstanceRef: refs.SQLInstanceRef{
					Name:      "nonexistent-instance",
					Namespace: "test-ns",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obj := &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					CloneSource: tc.cloneSource,
				},
			}
			obj.Namespace = "test-ns"

			err := resolveSourceSQLInstanceRef(ctx, kube, obj)
			if (err != nil) != tc.wantErr {
				t.Fatalf("resolveSourceSQLInstanceRef() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr && tc.cloneSource != nil {
				if obj.Spec.CloneSource.SQLInstanceRef.External != tc.wantExternal {
					t.Errorf("obj.Spec.CloneSource.SQLInstanceRef.External = %q, want %q", obj.Spec.CloneSource.SQLInstanceRef.External, tc.wantExternal)
				}
			}
		})
	}
}

func TestResolveFailoverDRReplicaRef(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := krm.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add krm to scheme: %v", err)
	}

	referencedSQLInstance := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "k8s-replica-instance",
				"namespace": "test-ns",
			},
			"spec": map[string]any{
				"resourceID": "resolved-replica-id",
			},
		},
	}

	kube := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(referencedSQLInstance).
		Build()

	ctx := context.Background()

	tests := []struct {
		name               string
		replicationCluster *krm.ReplicationCluster
		wantExternal       string
		wantErr            bool
	}{
		{
			name:               "nil replicationCluster",
			replicationCluster: nil,
			wantErr:            false,
		},
		{
			name: "bare instance name in external",
			replicationCluster: &krm.ReplicationCluster{
				FailoverDRReplicaRef: &refs.SQLInstanceRef{
					External: "my-replica-instance",
				},
			},
			wantExternal: "my-replica-instance",
			wantErr:      false,
		},
		{
			name: "resource path in external",
			replicationCluster: &krm.ReplicationCluster{
				FailoverDRReplicaRef: &refs.SQLInstanceRef{
					External: "projects/my-project/instances/my-replica-instance",
				},
			},
			wantExternal: "my-replica-instance",
			wantErr:      false,
		},
		{
			name: "selfLink in external",
			replicationCluster: &krm.ReplicationCluster{
				FailoverDRReplicaRef: &refs.SQLInstanceRef{
					External: "https://sqladmin.googleapis.com/sql/v1beta4/projects/my-project/instances/my-replica-instance",
				},
			},
			wantExternal: "my-replica-instance",
			wantErr:      false,
		},
		{
			name: "valid name reference",
			replicationCluster: &krm.ReplicationCluster{
				FailoverDRReplicaRef: &refs.SQLInstanceRef{
					Name:      "k8s-replica-instance",
					Namespace: "test-ns",
				},
			},
			wantExternal: "resolved-replica-id",
			wantErr:      false,
		},
		{
			name: "both external and name specified",
			replicationCluster: &krm.ReplicationCluster{
				FailoverDRReplicaRef: &refs.SQLInstanceRef{
					External: "my-replica-instance",
					Name:     "k8s-replica-instance",
				},
			},
			wantErr: true,
		},
		{
			name: "neither external nor name specified",
			replicationCluster: &krm.ReplicationCluster{
				FailoverDRReplicaRef: &refs.SQLInstanceRef{},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obj := &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					ReplicationCluster: tc.replicationCluster,
				},
			}
			obj.Namespace = "test-ns"

			err := resolveFailoverDRReplicaRef(ctx, kube, obj)
			if (err != nil) != tc.wantErr {
				t.Fatalf("resolveFailoverDRReplicaRef() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr && tc.replicationCluster != nil && tc.replicationCluster.FailoverDRReplicaRef != nil {
				if obj.Spec.ReplicationCluster.FailoverDRReplicaRef.External != tc.wantExternal {
					t.Errorf("obj.Spec.ReplicationCluster.FailoverDRReplicaRef.External = %q, want %q", obj.Spec.ReplicationCluster.FailoverDRReplicaRef.External, tc.wantExternal)
				}
			}
		})
	}
}
