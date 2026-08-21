// Copyright 2022 Google LLC
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

package k8s_test

import (
	"fmt"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/appscode/jsonpatch"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	mgr manager.Manager
)

func TestIsDeleted(t *testing.T) {
	nowTime := metav1.Now()
	testCases := []struct {
		Name           string
		Time           *metav1.Time
		ExpectedResult bool
	}{
		{"Nil time", nil, false},
		{"Now time", &nowTime, true},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			meta := metav1.ObjectMeta{
				DeletionTimestamp: tc.Time,
			}
			result := k8s.IsDeleted(&meta)
			if result != tc.ExpectedResult {
				t.Errorf("result mismatch: got '%v', want '%v'", result, tc.ExpectedResult)
			}
		})
	}
}

func TestGVKToGVR(t *testing.T) {
	tests := []struct {
		gvk         runtimeschema.GroupVersionKind
		expectedGVR runtimeschema.GroupVersionResource
	}{
		{
			gvk:         runtimeschema.GroupVersionKind{Kind: "ComputeVPNGateway"},
			expectedGVR: runtimeschema.GroupVersionResource{Resource: "computevpngateways"},
		},
		{
			gvk:         runtimeschema.GroupVersionKind{Kind: "KMSCryptoKey"},
			expectedGVR: runtimeschema.GroupVersionResource{Resource: "kmscryptokeys"},
		},
		{
			gvk:         runtimeschema.GroupVersionKind{Kind: "IAMPolicy"},
			expectedGVR: runtimeschema.GroupVersionResource{Resource: "iampolicies"},
		},
		{
			gvk:         runtimeschema.GroupVersionKind{Kind: "ComputeAddress"},
			expectedGVR: runtimeschema.GroupVersionResource{Resource: "computeaddresses"},
		},
		{
			gvk:         runtimeschema.GroupVersionKind{Kind: "FirestoreIndex"},
			expectedGVR: runtimeschema.GroupVersionResource{Resource: "firestoreindexes"},
		},
		{
			gvk:         runtimeschema.GroupVersionKind{Kind: "NetworkServicesMesh"},
			expectedGVR: runtimeschema.GroupVersionResource{Resource: "networkservicesmeshes"},
		},
		{
			gvk:         runtimeschema.GroupVersionKind{Kind: "PubSubTopic"},
			expectedGVR: runtimeschema.GroupVersionResource{Resource: "pubsubtopics"},
		},
	}
	for _, tc := range tests {
		if got, want := k8s.ToGVR(tc.gvk), tc.expectedGVR; got != want {
			t.Errorf("result mismatch: got '%v', want '%v'", got, want)
		}
	}
}

func TestHasAbandonAnnotation(t *testing.T) {
	tests := []struct {
		name                 string
		annotations          map[string]string
		hasAbandonAnnotation bool
	}{
		{
			name: "has deletion policy annotation set as abandon",
			annotations: map[string]string{
				k8s.DeletionPolicyAnnotation: k8s.DeletionPolicyAbandon,
			},
			hasAbandonAnnotation: true,
		},
		{
			name: "has deletion policy annotation set as delete",
			annotations: map[string]string{
				k8s.DeletionPolicyAnnotation: k8s.DeletionPolicyDelete,
			},
			hasAbandonAnnotation: false,
		},
		{
			name: "has deletion policy annotation set to empty string",
			annotations: map[string]string{
				k8s.DeletionPolicyAnnotation: "",
			},
			hasAbandonAnnotation: false,
		},
		{
			name:                 "has no deletion policy annotation",
			annotations:          map[string]string{},
			hasAbandonAnnotation: false,
		},
		{
			name:                 "has nil annotations map",
			hasAbandonAnnotation: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			obj := &unstructured.Unstructured{}
			obj.SetAnnotations(tc.annotations)
			actual := k8s.HasAbandonAnnotation(obj)
			if actual != tc.hasAbandonAnnotation {
				t.Errorf("incorrect value for HasAbandonAnnotation(): got %v, want %v", actual, tc.hasAbandonAnnotation)
			}
		})
	}
}

func TestSetDefaultContainerAnnotation(t *testing.T) {
	const (
		nsName    = "namespace-1"
		projectID = "project-1"
		folderID  = "1234567890"
		orgID     = "0987654321"
	)
	tests := []struct {
		name            string
		objAnnotations  map[string]string
		nsAnnotations   map[string]string
		containers      []corekccv1alpha1.Container
		expectedPatches []jsonpatch.JsonPatchOperation
		shouldErr       bool
	}{
		{
			name:          "no defaulting if containers list is empty",
			nsAnnotations: map[string]string{k8s.ProjectIDAnnotation: projectID},
			containers:    []corekccv1alpha1.Container{},
		},
		{
			name:           "prefer resource-level to namespace-level annotation for same type",
			objAnnotations: map[string]string{k8s.ProjectIDAnnotation: projectID},
			nsAnnotations:  map[string]string{k8s.ProjectIDAnnotation: "other-project-id"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
		},
		{
			name:           "prefer resource-level to namespace-level annotation for different types",
			objAnnotations: map[string]string{k8s.FolderIDAnnotation: folderID},
			nsAnnotations:  map[string]string{k8s.ProjectIDAnnotation: projectID},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
		},
		{
			name:           "prefer resource-level annotation to namespace name",
			objAnnotations: map[string]string{k8s.ProjectIDAnnotation: projectID},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
		},
		{
			name:           "add annotation from namespace-level when no resource-level annotation present",
			objAnnotations: map[string]string{"key": "value"},
			nsAnnotations:  map[string]string{k8s.ProjectIDAnnotation: projectID},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedPatches: []jsonpatch.JsonPatchOperation{{
				Operation: "add",
				Path:      fmt.Sprintf("/metadata/annotations/%v~1%v", k8s.AnnotationPrefix, "project-id"),
				Value:     projectID,
			}},
		},
		{
			name:          "defaulting creates a new annotations map when none present",
			nsAnnotations: map[string]string{k8s.ProjectIDAnnotation: projectID},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedPatches: []jsonpatch.JsonPatchOperation{{
				Operation: "add",
				Path:      "/metadata/annotations",
				Value:     map[string]interface{}{k8s.ProjectIDAnnotation: projectID},
			}},
		},
		{
			name:           "project-scoped resources use namespace name as project ID when no override present",
			objAnnotations: map[string]string{"key": "value"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedPatches: []jsonpatch.JsonPatchOperation{{
				Operation: "add",
				Path:      fmt.Sprintf("/metadata/annotations/%v~1%v", k8s.AnnotationPrefix, "project-id"),
				Value:     nsName,
			}},
		},
		{
			name:           "folder-scoped resources use folder ID annotation",
			objAnnotations: map[string]string{"key": "value"},
			nsAnnotations:  map[string]string{k8s.FolderIDAnnotation: folderID},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			expectedPatches: []jsonpatch.JsonPatchOperation{{
				Operation: "add",
				Path:      fmt.Sprintf("/metadata/annotations/%v~1%v", k8s.AnnotationPrefix, "folder-id"),
				Value:     folderID,
			}},
		},
		{
			name:           "org-scoped resources use org ID annotation",
			objAnnotations: map[string]string{"key": "value"},
			nsAnnotations:  map[string]string{k8s.OrgIDAnnotation: orgID},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			expectedPatches: []jsonpatch.JsonPatchOperation{{
				Operation: "add",
				Path:      fmt.Sprintf("/metadata/annotations/%v~1%v", k8s.AnnotationPrefix, "organization-id"),
				Value:     orgID,
			}},
		},
		{
			name: "fail if no default can be determined for non-project-scoped resources",
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name: "fail if ambiguous resource-level container annotation",
			objAnnotations: map[string]string{
				k8s.FolderIDAnnotation: folderID,
				k8s.OrgIDAnnotation:    orgID,
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name: "fail if ambiguous resource-level container annotation (with one being set to empty string)",
			objAnnotations: map[string]string{
				k8s.FolderIDAnnotation: "",
				k8s.OrgIDAnnotation:    orgID,
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name: "fail if ambiguous resource-level container annotation (with both being set to empty string)",
			objAnnotations: map[string]string{
				k8s.FolderIDAnnotation: "",
				k8s.OrgIDAnnotation:    "",
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name: "fail if ambiguous namespace-level container annotation",
			nsAnnotations: map[string]string{
				k8s.FolderIDAnnotation: folderID,
				k8s.OrgIDAnnotation:    orgID,
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name: "fail if ambiguous namespace-level container annotation (with one being set to empty string)",
			nsAnnotations: map[string]string{
				k8s.FolderIDAnnotation: "",
				k8s.OrgIDAnnotation:    orgID,
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name: "fail if ambiguous namespace-level container annotation (with both being set to empty string)",
			nsAnnotations: map[string]string{
				k8s.FolderIDAnnotation: "",
				k8s.OrgIDAnnotation:    "",
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ns := &corev1.Namespace{}
			ns.SetName(nsName)
			ns.SetAnnotations(tc.nsAnnotations)

			obj := &unstructured.Unstructured{}
			obj.SetNamespace(nsName)
			obj.SetAnnotations(tc.objAnnotations)

			newObj := obj.DeepCopy()
			err := k8s.SetDefaultContainerAnnotation(newObj, ns, tc.containers)
			if tc.shouldErr {
				if err == nil {
					t.Errorf("expected error but there was none")
				}
				return
			}

			if err != nil {
				t.Errorf("error setting default container annotation: %v", err)
				return
			}
			objRaw, err := obj.MarshalJSON()
			if err != nil {
				t.Fatalf("error marshaling old object as JSON: %v", err)
			}
			newObjRaw, err := newObj.MarshalJSON()
			if err != nil {
				t.Fatalf("error marshaling new object as JSON: %v", err)
			}
			patches := admission.PatchResponseFromRaw(objRaw, newObjRaw).Patches
			if len(patches) != len(tc.expectedPatches) {
				t.Errorf("expected %v patch(es), but got %v; expected: %+v, actual: %+v",
					len(tc.expectedPatches), len(patches), tc.expectedPatches, patches)
				return
			}
			// Should only have either 0 or 1 patches, so ordering is unimportant
			for i, p := range patches {
				if !test.Equals(t, tc.expectedPatches[i], p) {
					t.Errorf("expected patch: %+v, actual patch: %+v", tc.expectedPatches[i], p)
				}
			}
		})
	}
}

func TestIsManagedByKCC(t *testing.T) {
	tests := []struct {
		gvk            runtimeschema.GroupVersionKind
		expectedResult bool
	}{
		{
			gvk:            runtimeschema.GroupVersionKind{Group: "core.cnrm.cloud.google.com"},
			expectedResult: false,
		},
		{
			gvk:            runtimeschema.GroupVersionKind{Group: "storage.cnrm.cloud.google.com"},
			expectedResult: true,
		},
		{
			gvk:            runtimeschema.GroupVersionKind{Group: "test.cloud.google.com"},
			expectedResult: false,
		},
	}
	for _, tc := range tests {
		if got, want := k8s.IsManagedByKCC(tc.gvk), tc.expectedResult; got != want {
			t.Errorf("result mismatch: got '%v', want '%v'", got, want)
		}
	}
}

func TestMain(m *testing.M) {
	testmain.ForUnitTests(m, &mgr)
}
