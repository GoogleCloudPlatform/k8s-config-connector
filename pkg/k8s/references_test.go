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
	"reflect"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSetDefaultHierarchicalReference(t *testing.T) {
	tests := []struct {
		name             string
		resource         *k8s.Resource
		ns               *corev1.Namespace
		hierarchicalRefs []corekccv1alpha1.HierarchicalReference
		containers       []corekccv1alpha1.Container

		expectedResource *k8s.Resource
		shouldErr        bool
	}{
		{
			name: "no defaulting if resource doesn't support hierarchical references",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
		},
		{
			name: "no defaulting if resource specifies a reference already",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"name": "project-id-from-spec",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"name": "project-id-from-spec",
					},
				},
			},
		},
		{
			name: "default from resource annotation if specified and supported over namespace annotation of same type",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"external": "project-id-from-resource-annotation",
					},
				},
			},
		},
		{
			name: "default from resource annotation if specified and supported over namespace annotation of different type",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
					"folderRef": map[string]interface{}{
						"external": "folder-id-from-resource-annotation",
					},
				},
			},
		},
		{
			name: "default from resource annotation if specified and supported over namespace name",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"external": "project-id-from-resource-annotation",
					},
				},
			},
		},
		{
			name: "default from namespace annotation if resource annotations are not supported",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"external": "project-id-from-namespace-annotation",
					},
				},
			},
		},
		{
			name: "default from namespace annotation if no resource annotation specified",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedResource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"external": "project-id-from-namespace-annotation",
					},
				},
			},
		},
		{
			name: "default from namespace name if no namespace annotation specified and resource annotations are not supported",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"external": "namespace-name",
					},
				},
			},
		},
		{
			name: "default from namespace name if no namespace or resource annotations specified",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expectedResource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"external": "namespace-name",
					},
				},
			},
		},
		{
			name: "error if no namespace or resource annotations specified and resource does not support project references",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			shouldErr: true,
		},
		{
			name: "only default supported reference from resource annotation",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-resource-annotation",
						k8s.OrgIDAnnotation:     "org-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-namespace-annotation",
						k8s.OrgIDAnnotation:     "org-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-resource-annotation",
						k8s.OrgIDAnnotation:     "org-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
					"folderRef": map[string]interface{}{
						"external": "folder-id-from-resource-annotation",
					},
				},
			},
		},
		{
			name: "only default supported reference from namespace annotation (no resource annotations specified)",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-namespace-annotation",
						k8s.OrgIDAnnotation:     "org-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			expectedResource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
					"folderRef": map[string]interface{}{
						"external": "folder-id-from-namespace-annotation",
					},
				},
			},
		},
		{
			name: "only default supported reference from namespace annotation (no resource annotations supported)",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-namespace-annotation",
						k8s.OrgIDAnnotation:     "org-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			expectedResource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
					"folderRef": map[string]interface{}{
						"external": "folder-id-from-namespace-annotation",
					},
				},
			},
		},
		{
			name: "error if multiple references supported and multiple resource annotations specified",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			shouldErr: true,
		},
		{
			name: "error if multiple references supported and multiple resource annotations specified with one being set to empty string",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						k8s.FolderIDAnnotation:  "",
					},
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			shouldErr: true,
		},
		{
			name: "error if multiple references supported and multiple namespace annotations specified (no resource annotations specified)",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			shouldErr: true,
		},
		{
			name: "error if multiple references supported and multiple namespace annotations specified with one being set to empty string (no resource annotations specified)",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			shouldErr: true,
		},
		{
			name: "error if multiple references supported and multiple namespace annotations specified (no resource annotations supported)",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			shouldErr: true,
		},
		{
			name: "error if multiple references supported and multiple namespace annotations specified with one being set to empty string (no resource annotations supported)",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						k8s.FolderIDAnnotation:  "",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			shouldErr: true,
		},
		{
			name: "defaulting creates a new spec map if none present",
			resource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			expectedResource: &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
					},
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "project-id-from-resource-annotation",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := k8s.SetDefaultHierarchicalReference(tc.resource, tc.ns, tc.hierarchicalRefs, tc.containers)
			if tc.shouldErr && err == nil {
				t.Fatalf("got no error, but wanted one")
			} else if !tc.shouldErr && err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}

			if err == nil {
				if !test.Equals(t, tc.expectedResource, tc.resource) {
					diff := cmp.Diff(tc.expectedResource, tc.resource)
					t.Fatalf("unexpected diff in resource (-want +got):\n%v", diff)
				}
			}
		})
	}
}

func TestGetHierarchicalReference(t *testing.T) {
	tests := []struct {
		name             string
		resource         *k8s.Resource
		hierarchicalRefs []corekccv1alpha1.HierarchicalReference

		resourceRef     *corekccv1alpha1.ResourceReference
		hierarchicalRef corekccv1alpha1.HierarchicalReference
		shouldErr       bool
	}{
		{
			name:     "return nil resource reference if resource doesn't have a spec",
			resource: &k8s.Resource{},
		},
		{
			name: "return nil resource reference if resource doesn't support hierarchical references",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-dep",
					},
				},
			},
		},
		{
			name: "return nil resource reference if resource supports a hierarchical reference, but spec does not specify that reference",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
		},
		{
			name: "return resource reference if resource supports a hierarchical reference, and spec specifies that reference",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-dep",
					},
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
			resourceRef: &corekccv1alpha1.ResourceReference{
				Name: "folder-dep",
			},
			hierarchicalRef: corekccv1alpha1.HierarchicalReference{
				Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
				Key:  "folderRef",
			},
		},
		{
			name: "return nil resource reference if resource supports multiple hierarchical references, and spec specifies none of those references",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
					Key:  "organizationRef",
				},
			},
		},
		{
			name: "return resource reference if resource supports multiple hierarchical references, and spec specifies one of those references",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-dep",
					},
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
					Key:  "organizationRef",
				},
			},
			resourceRef: &corekccv1alpha1.ResourceReference{
				Name: "folder-dep",
			},
			hierarchicalRef: corekccv1alpha1.HierarchicalReference{
				Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
				Key:  "folderRef",
			},
		},
		{
			name: "error if resource supports multiple hierarchical references, but spec specifies more than one of those references",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-dep",
					},
					"organizationRef": map[string]interface{}{
						"external": "123456789",
					},
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
					Key:  "organizationRef",
				},
			},
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			resourceRef, hierarchicalRef, err := k8s.GetHierarchicalReference(tc.resource, tc.hierarchicalRefs)
			if tc.shouldErr {
				if err == nil {
					t.Fatalf("got no error, but wanted one")
				}
				if resourceRef != nil {
					t.Fatalf("got a non-nil resource reference, but wanted nil since function returned an error")
				}
				if !isEmpty(hierarchicalRef) {
					t.Fatalf("got a non-empty hierarchical reference, but wanted empty since function returned an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}

			if !test.Equals(t, tc.resourceRef, resourceRef) {
				diff := cmp.Diff(tc.resourceRef, resourceRef)
				t.Fatalf("unexpected diff in resulting resource reference (-want +got):\n%v", diff)
			}
			if !test.Equals(t, tc.hierarchicalRef, hierarchicalRef) {
				diff := cmp.Diff(tc.hierarchicalRef, hierarchicalRef)
				t.Fatalf("unexpected diff in resulting hierarchical reference (-want +got):\n%v", diff)
			}
		})
	}
}

func TestSetHierarchicalReference(t *testing.T) {
	tests := []struct {
		name            string
		resource        *k8s.Resource
		hierarchicalRef *corekccv1alpha1.HierarchicalReference
		val             string

		expectedResource *k8s.Resource
	}{
		{
			name: "add hierarchical reference if not present",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			hierarchicalRef: &corekccv1alpha1.HierarchicalReference{
				Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
				Key:  "projectRef",
			},
			val: "project-id",
			expectedResource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "val",
					"projectRef": map[string]interface{}{
						"external": "project-id",
					},
				},
			},
		},
		{
			name: "overwrite hierarchical reference if present",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "old-project-id",
					},
				},
			},
			hierarchicalRef: &corekccv1alpha1.HierarchicalReference{
				Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
				Key:  "projectRef",
			},
			val: "project-id",
			expectedResource: &k8s.Resource{
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "project-id",
					},
				},
			},
		},
		{
			name:     "add hierarchical reference even if spec not present",
			resource: &k8s.Resource{},
			hierarchicalRef: &corekccv1alpha1.HierarchicalReference{
				Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
				Key:  "projectRef",
			},
			val: "project-id",
			expectedResource: &k8s.Resource{
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "project-id",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if err := k8s.SetHierarchicalReference(tc.resource, tc.hierarchicalRef, tc.val); err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !test.Equals(t, tc.expectedResource, tc.resource) {
				diff := cmp.Diff(tc.expectedResource, tc.resource)
				t.Fatalf("unexpected diff in resource (-want +got):\n%v", diff)
			}
		})
	}
}

func isEmpty(h corekccv1alpha1.HierarchicalReference) bool {
	empty := corekccv1alpha1.HierarchicalReference{}
	return reflect.DeepEqual(h, empty)
}
