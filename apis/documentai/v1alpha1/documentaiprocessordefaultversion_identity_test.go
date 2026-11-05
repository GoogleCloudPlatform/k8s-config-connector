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

package v1alpha1

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDocumentAIProcessorDefaultVersionIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DocumentAIProcessorDefaultVersionIdentity
	}{
		{
			name: "valid relative reference",
			ref:  "projects/my-project/locations/us/processors/my-processor",
			want: &DocumentAIProcessorDefaultVersionIdentity{
				Project:   "my-project",
				Location:  "us",
				Processor: "my-processor",
			},
		},
		{
			name: "valid reference with host prefix",
			ref:  "documentai.googleapis.com/projects/my-project/locations/us/processors/my-processor",
			want: &DocumentAIProcessorDefaultVersionIdentity{
				Project:   "my-project",
				Location:  "us",
				Processor: "my-processor",
			},
		},
		{
			name: "valid reference with full url",
			ref:  "//documentai.googleapis.com/projects/my-project/locations/us/processors/my-processor",
			want: &DocumentAIProcessorDefaultVersionIdentity{
				Project:   "my-project",
				Location:  "us",
				Processor: "my-processor",
			},
		},
		{
			name:    "invalid format - empty",
			ref:     "",
			wantErr: true,
		},
		{
			name:    "invalid format - missing location",
			ref:     "projects/my-project/processors/my-processor",
			wantErr: true,
		},
		{
			name:    "invalid format - wrong collection",
			ref:     "projects/my-project/locations/us/other/my-processor",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DocumentAIProcessorDefaultVersionIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if gotString := i.String(); gotString != "projects/my-project/locations/us/processors/my-processor" {
					t.Errorf("String() = %q, want %q", gotString, "projects/my-project/locations/us/processors/my-processor")
				}
				if gotHost := i.Host(); gotHost != "documentai.googleapis.com" {
					t.Errorf("Host() = %q, want %q", gotHost, "documentai.googleapis.com")
				}
				if gotParent := i.ParentString(); gotParent != "projects/my-project/locations/us" {
					t.Errorf("ParentString() = %q, want %q", gotParent, "projects/my-project/locations/us")
				}
			}
		})
	}
}

func TestGetIdentityFromDocumentAIProcessorDefaultVersionSpec(t *testing.T) {
	tests := []struct {
		name    string
		obj     *DocumentAIProcessorDefaultVersion
		want    *DocumentAIProcessorDefaultVersionIdentity
		wantErr bool
	}{
		{
			name: "resourceID is full processor reference",
			obj: &DocumentAIProcessorDefaultVersion{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-default-version",
					Namespace: "test-ns",
				},
				Spec: DocumentAIProcessorDefaultVersionSpec{
					ResourceID: common.LazyPtr("projects/my-project/locations/us/processors/my-processor"),
					Version:    "projects/my-project/locations/us/processors/my-processor/processorVersions/v1",
				},
			},
			want: &DocumentAIProcessorDefaultVersionIdentity{
				Project:   "my-project",
				Location:  "us",
				Processor: "my-processor",
			},
		},
		{
			name: "resourceID is full reference with host",
			obj: &DocumentAIProcessorDefaultVersion{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-default-version",
					Namespace: "test-ns",
				},
				Spec: DocumentAIProcessorDefaultVersionSpec{
					ResourceID: common.LazyPtr("//documentai.googleapis.com/projects/my-project/locations/us/processors/my-processor"),
					Version:    "projects/my-project/locations/us/processors/my-processor/processorVersions/v1",
				},
			},
			want: &DocumentAIProcessorDefaultVersionIdentity{
				Project:   "my-project",
				Location:  "us",
				Processor: "my-processor",
			},
		},
		{
			name: "resourceID is short name without location in spec",
			obj: &DocumentAIProcessorDefaultVersion{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-default-version",
					Namespace: "test-ns",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DocumentAIProcessorDefaultVersionSpec{
					ResourceID: common.LazyPtr("my-processor"),
					Version:    "v1",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentityFromDocumentAIProcessorDefaultVersionSpec(context.Background(), nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Fatalf("getIdentityFromDocumentAIProcessorDefaultVersionSpec() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("getIdentityFromDocumentAIProcessorDefaultVersionSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestDocumentAIProcessorDefaultVersion_GetIdentity(t *testing.T) {
	obj := &DocumentAIProcessorDefaultVersion{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-default-version",
			Namespace: "test-ns",
		},
		Spec: DocumentAIProcessorDefaultVersionSpec{
			ResourceID: common.LazyPtr("projects/my-project/locations/us/processors/my-processor"),
			Version:    "projects/my-project/locations/us/processors/my-processor/processorVersions/v1",
		},
	}

	got, err := obj.GetIdentity(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetIdentity() error = %v", err)
	}

	want := &DocumentAIProcessorDefaultVersionIdentity{
		Project:   "my-project",
		Location:  "us",
		Processor: "my-processor",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}
}

func TestDocumentAIProcessorDefaultVersionRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)

	readyObj := &DocumentAIProcessorDefaultVersion{
		TypeMeta: metav1.TypeMeta{
			APIVersion: DocumentAIProcessorDefaultVersionGVK.GroupVersion().String(),
			Kind:       DocumentAIProcessorDefaultVersionGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-version",
			Namespace: "default",
		},
		Spec: DocumentAIProcessorDefaultVersionSpec{
			ResourceID: common.LazyPtr("projects/test-project/locations/us/processors/ready-proc"),
			Version:    "projects/test-project/locations/us/processors/ready-proc/processorVersions/v1",
		},
		Status: DocumentAIProcessorDefaultVersionStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: corev1.ConditionTrue,
				},
			},
		},
	}

	notReadyObj := &DocumentAIProcessorDefaultVersion{
		TypeMeta: metav1.TypeMeta{
			APIVersion: DocumentAIProcessorDefaultVersionGVK.GroupVersion().String(),
			Kind:       DocumentAIProcessorDefaultVersionGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready-version",
			Namespace: "default",
		},
		Spec: DocumentAIProcessorDefaultVersionSpec{
			ResourceID: common.LazyPtr("projects/test-project/locations/us/processors/not-ready-proc"),
			Version:    "projects/test-project/locations/us/processors/not-ready-proc/processorVersions/v1",
		},
		Status: DocumentAIProcessorDefaultVersionStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: corev1.ConditionFalse,
				},
			},
		},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyObj, notReadyObj).Build()

	tests := []struct {
		name     string
		ref      *DocumentAIProcessorDefaultVersionRef
		wantErr  bool
		errCheck func(err error) bool
		wantExt  string
	}{
		{
			name: "external already set - valid",
			ref: &DocumentAIProcessorDefaultVersionRef{
				External: "projects/my-project/locations/us/processors/my-processor",
			},
			wantExt: "projects/my-project/locations/us/processors/my-processor",
		},
		{
			name: "external already set - invalid",
			ref: &DocumentAIProcessorDefaultVersionRef{
				External: "invalid/external/ref",
			},
			wantErr: true,
		},
		{
			name: "referenced resource is ready",
			ref: &DocumentAIProcessorDefaultVersionRef{
				Name:      "ready-version",
				Namespace: "default",
			},
			wantExt: "projects/test-project/locations/us/processors/ready-proc",
		},
		{
			name: "referenced resource is not ready",
			ref: &DocumentAIProcessorDefaultVersionRef{
				Name:      "not-ready-version",
				Namespace: "default",
			},
			wantErr: true,
			errCheck: func(err error) bool {
				return k8s.IsReferenceNotReadyError(err)
			},
		},
		{
			name: "referenced resource not found",
			ref: &DocumentAIProcessorDefaultVersionRef{
				Name:      "non-existent",
				Namespace: "default",
			},
			wantErr: true,
			errCheck: func(err error) bool {
				return k8s.IsReferenceNotFoundError(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(context.Background(), client, "default")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.errCheck != nil && !tt.errCheck(err) {
				t.Errorf("Normalize() unexpected error type: %v", err)
			}
			if !tt.wantErr && tt.ref.External != tt.wantExt {
				t.Errorf("Normalize() External = %q, want %q", tt.ref.External, tt.wantExt)
			}
		})
	}
}
